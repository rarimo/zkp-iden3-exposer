package zkp_iden3_exposer

import (
	"github.com/iden3/go-circuits/v2"
	"io"
	"os"
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

func TestConnector(t *testing.T) {
	PK := "9a5305fa4c55cbf517c99693a7ec6766203c88feab50c944c00feec051d5dab7"
	issuerApi := "https://issuer.polygon.robotornot.mainnet-beta.rarimo.com"
	claimType := "urn:uuid:6dff4518-5177-4f39-af58-9c156d9b6309"

	t.Run("should get did string and did id bigint string", func(t *testing.T) {
		response, err := GetIdentity(PK)

		if err != nil {
			t.Errorf("Error getting identity: %v", err)
		}

		println(string(response))
	})
	t.Run("Should get auth v2 inputs", func(t *testing.T) {
		response, err := GetAuthV2Inputs(PK, issuerApi, claimType)

		if err != nil {
			t.Errorf("Error getting auth v2 inputs: %v", err)
		}

		println(string(response))
	})
	t.Run("Should get atomic query mtp v2 on-chain inputs", func(t *testing.T) {
		vcB, err := GetFile("./mocks/vc.json")

		if err != nil {
			t.Errorf("Error getting file: %v", err)
		}

		// Test the GetIdentity function
		atomicQueryMTVV2OnChainInputs, err := GetAtomicQueryMTVV2OnChainInputs(
			"9a5305fa4c55cbf517c99693a7ec6766203c88feab50c944c00feec051d5dab7",
			vcB,
			"credentialAtomicQueryMTPV2OnChain",
			"EA931A38726546cB7B5992483867387fC9FAdF7b",
			"isNatural",
			"1",
			circuits.EQ,
		)

		if err != nil {
			t.Errorf("Error getting atomic query MTV V2 on-chain inputs: %v", err)
		}

		println(string(atomicQueryMTVV2OnChainInputs))
	})
}
