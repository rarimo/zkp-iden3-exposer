package zkp_iden3

import (
	"encoding/hex"
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	"github.com/iden3/go-jwz/v2"
	"github.com/iden3/go-merkletree-sql/v2"
	types2 "github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/helpers"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"math/big"
	"time"
)

type ZkpGenConfig struct {
	ChainInfo types.ChainZkpInfo
}

type ZkpGenCommonInputs struct {
	//CircuitClaim      types.CircuitClaim
	//Query             circuits.Query
	//NodeAuxNonRev     types.NodeAuxValue
	//ClaimNonRevStatus types.RevocationStatus
	//Value             []string
	//timestamp         *int
}

type ZkpGen struct {
	Config   ZkpGenConfig
	Identity *Identity
}

func NewZkpGen(config ZkpGenConfig, identity *Identity) *ZkpGen {
	return &ZkpGen{
		Config:   config,
		Identity: identity,
	}
}

func (z *ZkpGen) GenerateProof(
	coreStateHash string,
	operationGistHash string,
	vc verifiable.W3CCredential,
	proofRequest types.CreateProofRequest,
	Circuits types.CircuitPair,
) (*types2.ZKProof, error) {
	userId, err := z.Identity.ID()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get ID")
	}

	jsonString, err := json.Marshal(vc.CredentialStatus)

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal credential status")
	}

	var credStatus verifiable.CredentialStatus

	if err = json.Unmarshal(jsonString, &credStatus); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal credential status")
	}

	resolver := helpers.CredentialStatusResolver{
		Url:                        credStatus.ID,
		EndianSwappedCoreStateHash: nil,
	}

	verifiable.DefaultCredentialStatusResolverRegistry.Register(verifiable.SparseMerkleTreeProof, &resolver)
	//verifiable.DefaultCredentialStatusResolverRegistry.Register(verifiable.BJJSignatureProofType, &resolver) // FIXME

	revStatus, err := verifiable.ValidateCredentialStatus(nil, credStatus)

	if err != nil {
		return nil, errors.Wrap(err, "failed to validate credential status")
	}

	circuitIdProofTypeMap := map[circuits.CircuitID]verifiable.ProofType{
		circuits.AtomicQueryMTPV2CircuitID:        verifiable.Iden3SparseMerkleTreeProofType,
		circuits.AtomicQueryMTPV2OnChainCircuitID: verifiable.Iden3SparseMerkleTreeProofType,
		circuits.AtomicQuerySigV2OnChainCircuitID: verifiable.BJJSignatureProofType,
		circuits.AtomicQuerySigV2CircuitID:        verifiable.BJJSignatureProofType,
	}

	coreClaim, err := vc.GetCoreClaimFromProof(circuitIdProofTypeMap[proofRequest.CircuitId])

	if err != nil {
		return nil, errors.Wrap(err, "failed to get core claim from vc")
	}

	// TODO: implement
	query := circuits.Query{}

	// TODO: implement + revStatus
	claimWithMTPProof := circuits.ClaimWithMTPProof{}

	// TODO: check if this is correct
	_operationGistHash, err := merkletree.NewHashFromHex(operationGistHash)

	gistProofRaw, err := helpers.GetGISTProof(
		z.Identity.Config.ChainInfo.CoreEvmRpcApiUrl,
		z.Identity.Config.ChainInfo.CoreStateContractAddress,
		userId.BigInt(),
		_operationGistHash.BigInt(),
	)

	gistProof, err := helpers.ToGISTProof(*gistProofRaw)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get GIST proof")
	}

	targetChallenge := proofRequest.Challenge

	if targetChallenge == nil {
		accAddr := *proofRequest.AccountAddress

		*targetChallenge = accAddr[2:]
	}

	hexDecodedTargetChallenge, err := hex.DecodeString(*targetChallenge)

	if err != nil {
		return nil, errors.Wrap(err, "failed to decode account address")
	}

	challenge := helpers.FromLittleEndian(hexDecodedTargetChallenge)

	signature := z.Identity.PrivateKey.SignPoseidon(challenge)

	mtpv2OnchainInputs := circuits.AtomicQueryMTPV2OnChainInputs{
		ID:                       userId,
		ProfileNonce:             big.NewInt(0),
		ClaimSubjectProfileNonce: big.NewInt(0),

		Claim:                    claimWithMTPProof,
		SkipClaimRevocationCheck: false,

		RequestID: big.NewInt(0), // TODO: get from proofRequest

		CurrentTimeStamp: time.Now().Unix(),

		AuthClaim:          z.Identity.CoreAuthClaim,
		AuthClaimIncMtp:    z.Identity.AuthClaimIncProof,
		AuthClaimNonRevMtp: z.Identity.AuthClaimNonRevProof,
		TreeState:          *z.Identity.TreeState,

		GISTProof: *gistProof,

		Signature: signature,
		Challenge: challenge,

		Query: query,
	}

	encodedInputs, err := mtpv2OnchainInputs.InputsMarshal()

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal inputs")
	}

	println(string(encodedInputs))

	zkProof, err := jwz.ProvingMethodGroth16AuthV2Instance.Prove(encodedInputs, Circuits.ProvingKey, Circuits.Wasm)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create proof")
	}

	return zkProof, nil
}

//func (z *ZkpGen) PrepareInputs(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {
//	atomicQueryMTPV2OnChainInputs := circuits.AtomicQueryMTPV2OnChainInputs{}
//
//	encodedInputs, err := atomicQueryMTPV2OnChainInputs.InputsMarshal()
//
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to marshal inputs")
//	}
//
//	return encodedInputs, nil
//}
