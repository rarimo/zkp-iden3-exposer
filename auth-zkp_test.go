package zkp_iden3

import (
	"encoding/json"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func GetFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := file.Close()

		if err != nil {
			panic(err)
		}
	}()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetOffer(issuerApi string, identity *Identity, claimType string) (types.ClaimOffer, error) {
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

func GetVC(identity Identity, offer types.ClaimOffer) (*verifiable.W3CCredential, error) {
	wasmFileBytes, err := GetFile("./assets/circuits/auth/circuit.wasm")

	if err != nil {
		return nil, errors.Wrap(err, "Error getting wasm file")
	}

	provingKeyFileBytes, err := GetFile("./assets/circuits/auth/circuit_final.zkey")

	if err != nil {
		return nil, errors.Wrap(err, "Error getting proving key file")
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
		return nil, errors.Wrap(err, "Error getting vc")
	}

	return vc, nil
}

func TestGetVerifiableCredentials(t *testing.T) {
	issuerApi := "https://issuer.polygon.robotornot.mainnet-beta.rarimo.com"
	claimType := "urn:uuid:6dff4518-5177-4f39-af58-9c156d9b6309"
	identity := getIdentity(nil)
	offer := types.ClaimOffer{}

	t.Run("should get offer", func(t *testing.T) {
		claimOffer, err := GetOffer(issuerApi, &identity, claimType)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		offer = claimOffer
	})

	t.Run("should get vc", func(t *testing.T) {
		vc, err := GetVC(identity, offer)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if strings.Contains(vc.ID, offer.Body.Credentials[0].Id) == false {
			t.Errorf("Error: %v", err)
		}
	})
}
