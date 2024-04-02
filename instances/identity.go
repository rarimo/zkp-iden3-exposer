package instances

import (
	"encoding/json"
	"github.com/iden3/go-circuits"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/rarimo/go-merkletree"
	merkletree_db_memory "github.com/rarimo/go-merkletree/db/memory"
	"github.com/rarimo/zkp-iden3-exposer/constants"
	"github.com/rarimo/zkp-iden3-exposer/helpers"
	"github.com/rarimo/zkp-iden3-exposer/types"
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
	AuthClaimIncProofSiblings []*merkletree.Hash
	AuthClaimNonRevProof      *merkletree.Proof
	TreeState                 *types.TreeState
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

	claimsDB := merkletree_db_memory.NewMemoryStorage().WithPrefix([]byte("claims"))
	revocationsDB := merkletree_db_memory.NewMemoryStorage().WithPrefix([]byte("revocations"))
	rootsDB := merkletree_db_memory.NewMemoryStorage().WithPrefix([]byte("roots"))

	claimsTree, err := merkletree.NewMerkleTree(claimsDB, 32)

	if err != nil {
		return nil, err
	}

	revocationsTree, err := merkletree.NewMerkleTree(revocationsDB, 32)

	if err != nil {
		return nil, err
	}

	rootsTree, err := merkletree.NewMerkleTree(rootsDB, 32)

	if err != nil {
		return nil, err
	}

	claimsTree.Add(hi, hv)

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

	authClaimIncProof, _, err := claimsTree.GenerateProof(coreAuthClaimHIndex, claimsTreeRoot)

	if err != nil {
		return nil, err
	}

	authClaimIncProofSiblings := helpers.PrepareSiblingsStr(*authClaimIncProof, constants.DefaultMTLevels)

	authClaimNonRevProof, _, err := revocationsTree.GenerateProof(
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

	identity.TreeState = &types.TreeState{
		State:              *stateHash,
		ClaimsTreeRoot:     *claimsTreeRoot,
		RevocationTreeRoot: *revocationsTreeRoot,
		RootOfRoots:        *rootOfRoots,
	}

	return &identity, nil
}

func (i *Identity) DidString() string {
	return i.DID.String()
}

func (i *Identity) IdentityBigInt() (*big.Int, error) {
	parsedDID, err := w3c.ParseDID(i.DidString())

	if err != nil {
		return nil, err
	}

	id, err := core.IDFromDID(*parsedDID)

	if err != nil {
		return nil, err
	}

	return id.BigInt(), nil
}

func (i *Identity) IdentityBigIntString() (*string, error) {
	id, err := i.IdentityBigInt()

	if err != nil {
		return nil, err
	}

	idBigIntString := id.String()

	return &idBigIntString, nil
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

	userId, err := i.IdentityBigInt()

	if err != nil {
		return nil, err
	}

	gistProofRaw, err := helpers.GetGISTProof(
		i.Config.ChainInfo.CoreEvmRpcApiUrl,
		i.Config.ChainInfo.CoreStateContractAddress,
		userId,
		nil,
	)

	gistProof, err := helpers.ToGISTProof(*gistProofRaw)
	globalNodeAux := helpers.GetNodeAuxValue(gistProof.Proof)
	nodeAuxAuth := helpers.GetNodeAuxValue(*i.AuthClaimNonRevProof)

	preparedInputs := types.AuthV2CircuitInputs{
		GenesisID:    userId.String(),
		ProfileNonce: "0",

		AuthClaim:    i.CoreAuthClaim,
		AuthClaimMtp: i.AuthClaimIncProofSiblings,

		AuthClaimNonRevMtp:      helpers.PrepareSiblingsStr(*i.AuthClaimNonRevProof, constants.DefaultMTLevels),
		AuthClaimNonRevMtpAuxHi: &nodeAuxAuth.Key,
		AuthClaimNonRevMtpAuxHv: &nodeAuxAuth.Value,
		AuthClaimNonRevMtpNoAux: nodeAuxAuth.NoAux,

		Challenge:             hashBigInt.String(),
		ChallengeSignatureR8X: signature.R8.X.String(),
		ChallengeSignatureR8Y: signature.R8.Y.String(),
		ChallengeSignatureS:   signature.S.String(),

		ClaimsTreeRoot: &i.TreeState.ClaimsTreeRoot,
		RevTreeRoot:    &i.TreeState.RevocationTreeRoot,
		RootsTreeRoot:  &i.TreeState.RootOfRoots,
		State:          &i.TreeState.State,

		GISTRoot:     &gistProof.Root,
		GISTMtp:      helpers.PrepareSiblingsStr(gistProof.Proof, constants.DefaultMTLevelsOnChain),
		GISTMtpAuxHi: &globalNodeAux.Key,
		GISTMtpAuxHv: &globalNodeAux.Value,
		GISTMtpNoAux: globalNodeAux.NoAux,
	}

	encodedInputs, err := json.Marshal(preparedInputs)

	if err != nil {
		return nil, err
	}

	return encodedInputs, nil
}
