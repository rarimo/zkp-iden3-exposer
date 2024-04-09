package zkp_iden3_exposer

import (
	"encoding/hex"
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/instances"
	"github.com/rarimo/zkp-iden3-exposer/overrides"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"net/http"
)

func getIdentity(pkHex *string) instances.Identity {
	identity, _ := instances.NewIdentity(instances.IdentityConfig{
		IdType: [2]byte{
			0x01,
			0x00,
		},
		SchemaHashHex: "cca3371a6cb1b715004407e325bd993c",
		ChainInfo: types.ChainZkpInfo{
			TargetChainId:              11155111,
			TargetRpcUrl:               "https://endpoints.omniatech.io/v1/eth/sepolia/public",
			TargetStateContractAddress: "0x8a9F505bD8a22BF09b0c19F65C17426cd33f3912",

			CoreApiUrl:               "https://rpc-api.node1.mainnet-beta.rarimo.com",
			CoreEvmRpcApiUrl:         "https://rpc.evm.node1.mainnet-beta.rarimo.com",
			CoreStateContractAddress: "0x753a8678c85d5fb70A97CFaE37c84CE2fD67EDE8",
		},
	}, pkHex)

	return *identity
}

func getOffer(issuerApi string, identity *instances.Identity, claimType string) (types.ClaimOffer, error) {
	offer := types.ClaimOffer{}

	response, err := http.Get(issuerApi + "/v1/credentials/" + identity.DID.String() + "/" + claimType)

	if err != nil {
		return offer, errors.Wrap(err, "Error getting offer")
	}

	if err := json.NewDecoder(response.Body).Decode(&offer); err != nil {
		return offer, errors.Wrap(err, "Error decoding offer")
	}

	return offer, nil
}

func GetIdentity(pkHex string) ([]byte, error) {
	if pkHex == "" || &pkHex == nil {
		return nil, errors.New("Private key is required")
	}

	identity := getIdentity(&pkHex)

	did := identity.DID

	id, err := identity.ID()

	if err != nil {
		return nil, err
	}

	identityJson, err := json.Marshal(map[string]string{
		"did":       did.String(),
		"didBigInt": id.BigInt().String(),
	})

	if err != nil {
		return nil, err
	}

	return identityJson, nil
}

func GetAuthV2Inputs(
	privateKeyHex string,
	issuerApi string,
	claimType string,
) ([]byte, error) {
	if privateKeyHex == "" || &privateKeyHex == nil {
		return nil, errors.New("Private key is required")
	}

	identity := getIdentity(&privateKeyHex)

	offer, err := getOffer(issuerApi, &identity, claimType)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting offer")
	}

	return instances.GetAuthV2Inputs(identity, offer)
}

func GetAtomicQueryMTVV2OnChainInputs(
	privateKeyHex string,
	jsonVC []byte,

	circuitId string,
	challenge string,

	subjectFieldName string,
	subjectFieldValue string,
	operator int,
) ([]byte, error) {
	if privateKeyHex == "" || &privateKeyHex == nil {
		return nil, errors.New("Private key is required")
	}

	identity := getIdentity(&privateKeyHex)

	proofRequest := types.CreateProofRequest{
		CircuitId: circuits.CircuitID(circuitId),
		Challenge: challenge,
		Query: types.ProofQuery{
			SubjectFieldName:  subjectFieldName,
			SubjectFieldValue: subjectFieldValue,
			Operator:          operator,
		},
	}

	vc := overrides.W3CCredential{}

	if err := json.Unmarshal(jsonVC, &vc); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling vc")
	}

	vc.W3CCredential.Proof = verifiable.CredentialProofs(vc.Proof)

	type StateInfo struct {
		Index                    string `json:"index"`
		Hash                     string `json:"hash"`
		CreatedAtTimestamp       string `json:"createdAtTimestamp"`
		CreatedAtBlock           string `json:"createdAtBlock"`
		LastUpdateOperationIndex string `json:"lastUpdateOperationIndex"`
	}

	type StateInfoResponse struct {
		State StateInfo `json:"state"`
	}

	issuerDID, _ := w3c.ParseDID(vc.Issuer)

	issuerID, _ := core.IDFromDID(*issuerDID)

	issuerHexId := "0x" + hex.EncodeToString(issuerID.BigInt().Bytes())

	response, _ := http.Get(identity.Config.ChainInfo.CoreApiUrl + "/rarimo/rarimo-core/identity/state/" + issuerHexId)

	stateInfoResponse := StateInfoResponse{}

	if err := json.NewDecoder(response.Body).Decode(&stateInfoResponse); err != nil {
		return nil, errors.Wrap(err, "Error decoding response")
	}

	type OperationStatus string

	const (
		Signed      OperationStatus = "SIGNED"
		Initialized OperationStatus = "INITIALIZED"
		Approved    OperationStatus = "APPROVED"
		NotApproved OperationStatus = "NOT_APPROVED"
	)

	type Operation struct {
		Index         string `json:"index"`
		OperationType string `json:"operationType"`
		Details       struct {
			AtType        string `json:"@type"`
			Contract      string `json:"contract"`
			Chain         string `json:"chain"`
			GISTHash      string `json:"GISTHash"`
			StateRootHash string `json:"stateRootHash"`
			Timestamp     string `json:"timestamp"`
		} `json:"details"`
		Status    OperationStatus `json:"status"`
		Creator   string          `json:"creator"`
		Timestamp string          `json:"timestamp"`
	}

	type OperationResponse struct {
		Operation Operation `json:"operation"`
	}

	response, _ = http.Get(identity.Config.ChainInfo.CoreApiUrl + "/rarimo/rarimo-core/rarimocore/operation/" + stateInfoResponse.State.LastUpdateOperationIndex)

	operationResponse := OperationResponse{}

	if err := json.NewDecoder(response.Body).Decode(&operationResponse); err != nil {
		return nil, errors.Wrap(err, "Error decoding operation response")
	}

	atomicQueryMTPV2OnChainProof := instances.NewAtomicQueryMTPV2OnChainProof(
		identity,

		stateInfoResponse.State.Hash,
		operationResponse.Operation.Details.GISTHash,
		vc,
		proofRequest,
	)

	inputs, err := atomicQueryMTPV2OnChainProof.GetInputs()

	if err != nil {
		return nil, errors.Wrap(err, "Error getting inputs")
	}

	return inputs, nil
}

//func RemoveCredentials() {}

//func GetCredentials() {}

//func CheckStateContractSync() {}

//func CheckCredentialExistence() {}
