package instances

import (
	"encoding/json"
	"github.com/rarimo/zkp-iden3-exposer/helpers"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"net/http"
	"strings"
	"testing"
)

func TestGenerateZkProof(t *testing.T) {
	issuerApi := "https://issuer.polygon.robotornot.mainnet-beta.rarimo.com"
	claimType := "urn:uuid:6dff4518-5177-4f39-af58-9c156d9b6309"
	identity := getIdentity(nil)
	offer := types.ClaimOffer{}
	vc := &types.W3CCredential{}

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

		vc, err = authZkp.GetVerifiableCredentials(offer)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if strings.Contains(vc.Id, offer.Body.Credentials[0].Id) == false {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("should check vc and get core claim", func(t *testing.T) {
		_, _, err := helpers.CheckVCAndGetCoreClaim(*vc)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
