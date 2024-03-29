package instances

import (
	"encoding/json"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func getIdentity() Identity {
	PK := "9a5305fa4c55cbf517c99693a7ec6766203c88feab50c944c00feec051d5dab7"

	identity, _ := NewIdentity(IdentityConfig{
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
	}, &PK)

	return *identity
}

func getFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func TestGetVerifiableCredentials(t *testing.T) {
	issuerApi := "https://issuer.polygon.robotornot.mainnet-beta.rarimo.com"
	claimType := "urn:uuid:6dff4518-5177-4f39-af58-9c156d9b6309"
	identity := getIdentity()
	offer := types.ClaimOffer{}

	t.Run("should get offer", func(t *testing.T) {
		response, err := http.Get(issuerApi + "/v1/credentials/" + identity.DidString() + "/" + claimType)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if err := json.NewDecoder(response.Body).Decode(&offer); err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("should get vc", func(t *testing.T) {
		wasmFileBytes, err := getFile("../assets/circuits/auth/circuit.wasm")

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		provingKeyFileBytes, err := getFile("../assets/circuits/auth/circuit_final.zkey")

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		circuits := types.CircuitPair{
			Wasm:       wasmFileBytes,
			ProvingKey: provingKeyFileBytes,
		}

		authZkp := NewAuthZkp(AuthZkpConfig{
			ChainInfo: identity.Config.ChainInfo,
			Circuits:  circuits,
		}, identity)

		vc, err := authZkp.GetVerifiableCredentials(offer)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if strings.Contains(vc.Id, offer.Body.Credentials[0].Id) == false {
			t.Errorf("Error: %v", err)
		}
	})
}
