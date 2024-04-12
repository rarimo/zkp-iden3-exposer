package instances

import (
	"github.com/iden3/go-circuits/v2"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql/v2"
	merkletree_db_memory "github.com/iden3/go-merkletree-sql/v2/db/memory"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/internal/zkp/constants"
	"github.com/rarimo/zkp-iden3-exposer/internal/zkp/helpers"
	"github.com/rarimo/zkp-iden3-exposer/internal/zkp/types"
	"math/big"
)

type IdentityConfig struct {
	IdType        [2]byte
	SchemaHashHex string
	ChainInfo     types.ChainZkpInfo
}

type Identity struct {
	Config                    IdentityConfig
	PrivateKey                babyjub.PrivateKey
	DID                       w3c.DID
	AuthClaimIncProof         *merkletree.Proof
	AuthClaimIncProofSiblings []*merkletree.Hash
	AuthClaimNonRevProof      *merkletree.Proof
	TreeState                 *circuits.TreeState
	CoreAuthClaim             *core.Claim
}

func NewIdentity(config IdentityConfig, privateKeyHex *string) (*Identity, error) {
	identity := Identity{}

	identity.Config = config

	privateKey, err := helpers.InitSK(privateKeyHex)

	if err != nil {
		return nil, err
	}

	identity.PrivateKey = *privateKey

	coreAuthClaim, err := identity.createCoreAuthClaim()

	identity.CoreAuthClaim = coreAuthClaim

	hi, hv, err := identity.CoreAuthClaim.HiHv()

	if err != nil {
		return nil, err
	}

	claimsDB := merkletree_db_memory.NewMemoryStorage()      //.WithPrefix([]byte("claims"))
	revocationsDB := merkletree_db_memory.NewMemoryStorage() //.WithPrefix([]byte("revocations"))
	rootsDB := merkletree_db_memory.NewMemoryStorage()       //.WithPrefix([]byte("roots"))

	claimsTree, err := merkletree.NewMerkleTree(nil, claimsDB, 32)

	if err != nil {
		return nil, err
	}

	revocationsTree, err := merkletree.NewMerkleTree(nil, revocationsDB, 32)

	if err != nil {
		return nil, err
	}

	rootsTree, err := merkletree.NewMerkleTree(nil, rootsDB, 32)

	if err != nil {
		return nil, err
	}

	err = claimsTree.Add(nil, hi, hv)

	if err != nil {
		return nil, errors.Wrap(err, "failed to add hi, hv to claims tree")
	}

	claimsTreeRoot := claimsTree.Root()
	revocationsTreeRoot := revocationsTree.Root()
	rootOfRoots := rootsTree.Root()

	idenState, err := core.IdenState(claimsTreeRoot.BigInt(), revocationsTreeRoot.BigInt(), rootOfRoots.BigInt())

	if err != nil {
		return nil, err
	}

	did, err := core.NewDIDFromIdenState(identity.Config.IdType, idenState)

	if err != nil {
		return nil, err
	}

	identity.DID = *did

	coreAuthClaimHIndex, err := identity.CoreAuthClaim.HIndex()

	if err != nil {
		return nil, err
	}

	authClaimIncProof, _, err := claimsTree.GenerateProof(nil, coreAuthClaimHIndex, claimsTreeRoot)

	if err != nil {
		return nil, err
	}

	identity.AuthClaimIncProof = authClaimIncProof

	authClaimIncProofSiblings := helpers.PrepareSiblingsStr(*authClaimIncProof, constants.DefaultMTLevels)

	authClaimNonRevProof, _, err := revocationsTree.GenerateProof(

		nil,
		new(big.Int).SetUint64(identity.CoreAuthClaim.GetRevocationNonce()),
		revocationsTreeRoot,
	)

	if err != nil {
		return nil, err
	}

	identity.AuthClaimIncProofSiblings = authClaimIncProofSiblings
	identity.AuthClaimNonRevProof = authClaimNonRevProof

	stateHash, err := merkletree.HashElems(
		claimsTreeRoot.BigInt(),
		revocationsTreeRoot.BigInt(),
		rootOfRoots.BigInt(),
	)

	if err != nil {
		return nil, err
	}

	identity.TreeState = &circuits.TreeState{
		State:          stateHash,
		ClaimsRoot:     claimsTreeRoot,
		RevocationRoot: revocationsTreeRoot,
		RootOfRoots:    rootOfRoots,
	}

	return &identity, nil
}

func (i *Identity) ID() (*core.ID, error) {
	parsedDID, err := w3c.ParseDID(i.DID.String())

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse DID")
	}

	id, err := core.IDFromDID(*parsedDID)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get ID from DID")
	}

	return &id, nil
}

func (i *Identity) createCoreAuthClaim() (*core.Claim, error) {
	hash, err := core.NewSchemaHashFromHex(i.Config.SchemaHashHex)

	if err != nil {
		return nil, err
	}

	revNonce := make([]byte, 64)
	key := i.PrivateKey.Public()

	claim, err := core.NewClaim(
		hash,
		core.WithIndexDataInts(key.X, key.Y),
		core.WithRevocationNonce(helpers.FromBigEndian(revNonce).Uint64()),
	)

	if err != nil {
		return nil, err
	}

	return claim, nil
}

func (i *Identity) PrepareAuthV2Inputs(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {
	hashBigInt := helpers.FromBigEndian(hash)

	signature := i.PrivateKey.SignPoseidon(hashBigInt)

	userId, err := i.ID()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get ID")
	}

	gistProofRaw, err := helpers.GetGISTProof(
		i.Config.ChainInfo.CoreEvmRpcApiUrl,
		i.Config.ChainInfo.CoreStateContractAddress,
		userId.BigInt(),
		nil,
	)

	gistProof, err := helpers.ToGISTProof(*gistProofRaw)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get GIST proof")
	}

	preparedInputs := circuits.AuthV2Inputs{
		GenesisID:    userId,
		ProfileNonce: big.NewInt(0),

		AuthClaim: i.CoreAuthClaim,

		AuthClaimIncMtp:    i.AuthClaimIncProof,
		AuthClaimNonRevMtp: i.AuthClaimNonRevProof,
		TreeState:          *i.TreeState,

		GISTProof: *gistProof,

		Signature: signature,
		Challenge: hashBigInt,
	}

	encodedInputs, err := preparedInputs.InputsMarshal()

	if err != nil {
		return nil, err
	}

	return encodedInputs, nil
}