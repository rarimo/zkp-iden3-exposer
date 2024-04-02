package instances

import (
	"github.com/rarimo/zkp-iden3-exposer/helpers"
	"testing"
)

func TestGenerateZkProof(t *testing.T) {
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

	t.Run("should check vc and get core claim", func(t *testing.T) {
		_, _, err := helpers.CheckVCAndGetCoreClaim(*vc)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
