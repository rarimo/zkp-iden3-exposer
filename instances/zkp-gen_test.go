package instances

import (
	"encoding/hex"
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-jwz/v2"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"net/http"
	"testing"
)

func TestGenerateProof(t *testing.T) {
	issuerApi := "https://issuer.polygon.robotornot.mainnet-beta.rarimo.com"
	claimType := "urn:uuid:6dff4518-5177-4f39-af58-9c156d9b6309"
	identity := getIdentity(nil)
	offer, err := GetOffer(issuerApi, &identity, claimType)

	if err != nil {
		t.Errorf("Error getting offer: %v", err)
	}

	vc, err := GetVC(identity, offer)

	if err != nil {
		t.Errorf("Error getting VC: %v", err)
	}

	t.Run("should generate proof", func(t *testing.T) {
		accountAddress := "0xEA931A38726546cB7B5992483867387fC9FAdF7b"

		proofRequest := types.CreateProofRequest{
			CircuitId: circuits.AtomicQueryMTPV2OnChainCircuitID,
			Challenge: accountAddress[2:],
			Query: types.ProofQuery{
				SubjectFieldName:  "isNatural",
				SubjectFieldValue: "1",
				Operator:          circuits.EQ,
			},
		}

		circuitPair := types.CircuitPair{}

		if circuitPair.Wasm, err = GetFile("../assets/circuits/credentialAtomicQueryMTPV2OnChain/circuit.wasm"); err != nil {
			t.Errorf("Error getting file: %v", err)
		}

		if circuitPair.ProvingKey, err = GetFile("../assets/circuits/credentialAtomicQueryMTPV2OnChain/circuit_final.zkey"); err != nil {
			t.Errorf("Error getting file: %v", err)
		}

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
			t.Errorf("Error decoding response: %v", err)
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
			t.Errorf("Error decoding response: %v", err)
		}

		atomicQueryMTPV2OnChainProof := NewAtomicQueryMTPV2OnChainProof(
			identity,

			stateInfoResponse.State.Hash,
			operationResponse.Operation.Details.GISTHash,
			*vc,
			proofRequest,
		)

		inputs, err := atomicQueryMTPV2OnChainProof.GetInputs()

		if err != nil {
			t.Errorf("Error getting inputs: %v", err)
		}

		_, err = jwz.ProvingMethodGroth16AuthV2Instance.Prove(inputs, circuitPair.ProvingKey, circuitPair.Wasm)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
