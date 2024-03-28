package instances

import (
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/rarimo/go-circuits"
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
}

type Identity struct {
	Config                    IdentityConfig
	PrivateKey                babyjub.PrivateKey
	DID                       w3c.DID
	AuthClaimIncProofSiblings []*merkletree.Hash
	AuthClaimNonRevProof      merkletree.Proof
	TreeState                 types.TreeState
	CoreAuthClaim             core.Claim
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

	identity.CoreAuthClaim = *coreAuthClaim

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
	identity.AuthClaimNonRevProof = *authClaimNonRevProof

	stateHash, err := merkletree.HashElems(
		claimsTreeRoot.BigInt(),
		revocationsTreeRoot.BigInt(),
		rootOfRoots.BigInt(),
	)

	if err != nil {
		return nil, err
	}

	identity.TreeState = types.TreeState{
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

func (i *Identity) IdentityBigIntString() (*string, error) {
	parsedDID, err := w3c.ParseDID(i.DidString())

	if err != nil {
		return nil, err
	}

	id, err := core.IDFromDID(*parsedDID)

	if err != nil {
		return nil, err
	}

	idBigIntString := id.BigInt().String()

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

func (i *Identity) PrepareAuthV2Inputs(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {}
