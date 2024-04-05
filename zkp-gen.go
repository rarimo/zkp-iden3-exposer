package zkp_iden3

import (
	"encoding/hex"
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-jwz/v2"
	types2 "github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/helpers"
	"github.com/rarimo/zkp-iden3-exposer/overrides"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"math/big"
	"time"
)

func prepareCommonInputs(
	coreStateHash string,
	vc overrides.W3CCredential,
	proofRequest types.CreateProofRequest,
) (*circuits.ClaimWithMTPProof, *circuits.Query, error) {
	jsonString, err := json.Marshal(vc.CredentialStatus)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to marshal credential status")
	}

	var credStatus verifiable.CredentialStatus

	if err = json.Unmarshal(jsonString, &credStatus); err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal credential status")
	}

	resolver := helpers.CredentialStatusResolver{
		Url:                        credStatus.ID,
		EndianSwappedCoreStateHash: nil,
	}

	verifiable.DefaultCredentialStatusResolverRegistry.Register(verifiable.SparseMerkleTreeProof, &resolver)
	//verifiable.DefaultCredentialStatusResolverRegistry.Register(verifiable.BJJSignatureProofType, &resolver) // FIXME

	revStatus, err := verifiable.ValidateCredentialStatus(nil, credStatus)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to validate credential status")
	}

	circuitIdProofTypeMap := map[circuits.CircuitID]verifiable.ProofType{
		circuits.AtomicQueryMTPV2CircuitID:        verifiable.Iden3SparseMerkleTreeProofType,
		circuits.AtomicQueryMTPV2OnChainCircuitID: verifiable.Iden3SparseMerkleTreeProofType,
		circuits.AtomicQuerySigV2OnChainCircuitID: verifiable.BJJSignatureProofType,
		circuits.AtomicQuerySigV2CircuitID:        verifiable.BJJSignatureProofType,
	}

	coreClaim, err := vc.GetCoreClaimFromProof(circuitIdProofTypeMap[proofRequest.CircuitId])

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get core claim from vc")
	}

	query, err := helpers.ConvertProofRequestToCircuitQuery(&vc, &proofRequest)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to convert proof request to circuit query")
	}

	claimWithMTPProof := circuits.ClaimWithMTPProof{}

	claimWithMTPProof.Claim = coreClaim

	issuerDID, err := w3c.ParseDID(vc.Issuer)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse issuer DID")
	}

	issuerID, err := core.IDFromDID(*issuerDID)

	claimWithMTPProof.IssuerID = &issuerID

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get ID from DID")
	}

	smtProof := overrides.Iden3SparseMerkleTreeProof{}

	for _, proof := range vc.Proof {
		if proof.ProofType() == verifiable.Iden3SparseMerkleTreeProofType {
			encodedProof, err := json.Marshal(proof)

			if err != nil {
				return nil, nil, errors.Wrap(err, "failed to marshal proof")
			}

			if err := json.Unmarshal(encodedProof, &smtProof); err != nil {
				return nil, nil, errors.Wrap(err, "failed to unmarshal proof")
			}
		}
	}

	stateHashEndian, err := helpers.ConvertEndianSwappedCoreStateHashHex(coreStateHash)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to convert endian swapped core state hash")
	}

	smtRevStatus, err := helpers.GetRevocationStatus(smtProof.ID, stateHashEndian)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get revocation status")
	}

	smtRevStatusTreeState, err := helpers.BuildTreeState(
		*smtRevStatus.Issuer.State,
		*smtRevStatus.Issuer.ClaimsTreeRoot,
		*smtRevStatus.Issuer.RevocationTreeRoot,
		*smtRevStatus.Issuer.RootOfRoots,
	)

	incProof := circuits.MTProof{
		Proof:     &smtRevStatus.MTP,
		TreeState: *smtRevStatusTreeState,
	}

	claimWithMTPProof.IncProof = incProof

	revStatusIssuerTreeState, err := helpers.BuildTreeState(
		*revStatus.Issuer.State,
		*revStatus.Issuer.ClaimsTreeRoot,
		*revStatus.Issuer.RevocationTreeRoot,
		*revStatus.Issuer.RootOfRoots,
	)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to build rev status issuer tree state")
	}

	claimWithMTPProof.NonRevProof = circuits.MTProof{
		Proof:     &revStatus.MTP,
		TreeState: *revStatusIssuerTreeState,
	}

	return &claimWithMTPProof, query, nil
}

type AtomicQueryMTPV2OnChainProof struct {
	Identity Identity

	CoreStateHash     string
	OperationGistHash string
	VC                overrides.W3CCredential
	ProofRequest      types.CreateProofRequest
	Circuits          types.CircuitPair
}

func (a *AtomicQueryMTPV2OnChainProof) GenerateProof() (*types2.ZKProof, error) {
	claimWithMTPProof, query, err := prepareCommonInputs(a.CoreStateHash, a.VC, a.ProofRequest)

	userId, err := a.Identity.ID()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get ID")
	}

	operationGistHashBigInt := new(big.Int)

	operationGistHashBigInt.SetString(a.OperationGistHash, 16)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get hash from operationGistHash hex")
	}

	gistProofRaw, err := helpers.GetGISTProof(
		a.Identity.Config.ChainInfo.CoreEvmRpcApiUrl,
		a.Identity.Config.ChainInfo.CoreStateContractAddress,
		userId.BigInt(),
		operationGistHashBigInt,
	)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get GIST proof raw")
	}

	gistProof, err := helpers.ToGISTProof(*gistProofRaw)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get GIST proof")
	}

	hexDecodedChallenge, err := hex.DecodeString(a.ProofRequest.Challenge)

	if err != nil {
		return nil, errors.Wrap(err, "failed to decode challenge hex")
	}

	challenge := helpers.FromLittleEndian(hexDecodedChallenge)

	signature := a.Identity.PrivateKey.SignPoseidon(challenge)

	requestId := big.NewInt(0)

	if &a.ProofRequest.Id != nil {
		requestId.SetString(a.ProofRequest.Id, 10)
	}

	mtpv2OnchainInputs := circuits.AtomicQueryMTPV2OnChainInputs{
		ID:                       userId,
		ProfileNonce:             big.NewInt(0),
		ClaimSubjectProfileNonce: big.NewInt(0),

		Claim:                    *claimWithMTPProof,
		SkipClaimRevocationCheck: false,

		RequestID: requestId,

		CurrentTimeStamp: time.Now().Unix(),

		AuthClaim:          a.Identity.CoreAuthClaim,
		AuthClaimIncMtp:    a.Identity.AuthClaimIncProof,
		AuthClaimNonRevMtp: a.Identity.AuthClaimNonRevProof,
		TreeState:          *a.Identity.TreeState,

		GISTProof: *gistProof,

		Signature: signature,
		Challenge: challenge,

		Query: *query,
	}

	encodedInputs, err := mtpv2OnchainInputs.InputsMarshal()

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal inputs")
	}

	zkProof, err := jwz.ProvingMethodGroth16AuthV2Instance.Prove(encodedInputs, a.Circuits.ProvingKey, a.Circuits.Wasm)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create proof")
	}

	return zkProof, nil
}
