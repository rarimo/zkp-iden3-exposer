package zkp_iden3

import (
	"github.com/rarimo/zkp-iden3-exposer/types"
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
		zkpGen := NewZkpGen(ZkpGenConfig{
			ChainInfo: identity.Config.ChainInfo,
		}, &identity)

		proofRequest := types.CreateProofRequest{}

		circuitPair := types.CircuitPair{}

		if circuitPair.Wasm, err = GetFile("./assets/circuits/credentialAtomicQueryMTPV2OnChain/circuit.wasm"); err != nil {
			t.Errorf("Error getting file: %v", err)
		}

		if circuitPair.ProvingKey, err = GetFile("./assets/circuits/credentialAtomicQueryMTPV2OnChain/circuit_final.zkey"); err != nil {
			t.Errorf("Error getting file: %v", err)
		}

		_, err := zkpGen.GenerateProof("", "", *vc, proofRequest, circuitPair)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
