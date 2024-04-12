package instances

import (
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	"github.com/iden3/go-jwz/v2"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/internal/zkp/overrides"
	"github.com/rarimo/zkp-iden3-exposer/internal/zkp/types"
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

func GetVC(identity Identity, offer types.ClaimOffer) (*overrides.W3CCredential, error) {
	wasm, err := GetFile("../assets/circuits/auth/circuit.wasm")

	if err != nil {
		return nil, errors.Wrap(err, "Error getting wasm file")
	}

	provingKey, err := GetFile("../assets/circuits/auth/circuit_final.zkey")

	if err != nil {
		return nil, errors.Wrap(err, "Error getting proving key file")
	}

	circuitsPair := types.CircuitPair{
		Wasm:       wasm,
		ProvingKey: provingKey,
	}

	claimDetailsJson, err := GetClaimDetailsJson(offer)

	if err != nil {
		return nil, err
	}

	preparer := jwz.ProofInputsPreparerHandlerFunc(func(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {
		return identity.PrepareAuthV2Inputs(hash, circuitID)
	})

	token, err := jwz.NewWithPayload(
		jwz.ProvingMethodGroth16AuthV2Instance,
		claimDetailsJson,
		preparer,
	)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting AuthV2Inputs")
	}

	jwzTokenRaw, err := token.Prove(circuitsPair.ProvingKey, circuitsPair.Wasm)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting JWZ token")
	}

	vc, err := LoadVC(offer.Body.Url, jwzTokenRaw)

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

	t.Run("should get AuthV2 inputs", func(t *testing.T) {
		_, err := GetAuthV2Inputs(identity, offer)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
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
