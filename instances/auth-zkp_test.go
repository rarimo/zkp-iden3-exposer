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
	identity := getIdentity(nil)
	offer := types.ClaimOffer{}

	t.Run("should get offer", func(t *testing.T) {
		response, err := http.Get(issuerApi + "/v1/credentials/" + identity.DID.String() + "/" + claimType)

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
