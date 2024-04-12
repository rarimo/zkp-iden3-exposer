package zkp_iden3_exposer

import (
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	"github.com/iden3/go-jwz/v2"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/internal/zkp/instances"
	"github.com/rarimo/zkp-iden3-exposer/internal/zkp/types"
	"io"
	"os"
	"testing"
)

var identityConfig = types.IdentityConfig{
	PkHex: "9a5305fa4c55cbf517c99693a7ec6766203c88feab50c944c00feec051d5dab7",

	IdType: [2]byte{
		0x01,
		0x00,
	},
	SchemaHashHex: "cca3371a6cb1b715004407e325bd993c",

	TargetChainId:              11155111,
	TargetRpcUrl:               "https://endpoints.omniatech.io/v1/eth/sepolia/public",
	TargetStateContractAddress: "0x8a9F505bD8a22BF09b0c19F65C17426cd33f3912",

	CoreApiUrl:               "https://rpc-api.node1.mainnet-beta.rarimo.com",
	CoreEvmRpcApiUrl:         "https://rpc.evm.node1.mainnet-beta.rarimo.com",
	CoreStateContractAddress: "0x753a8678c85d5fb70A97CFaE37c84CE2fD67EDE8",
}

func getFile(path string) ([]byte, error) {
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

func getGroth16AuthV2ZKProof(identity instances.Identity, offer types.ClaimOffer) ([]byte, error) {
	wasm, err := getFile("./internal/zkp/assets/circuits/auth/circuit.wasm")

	if err != nil {
		return nil, errors.Wrap(err, "Error getting wasm file")
	}

	provingKey, err := getFile("./internal/zkp/assets/circuits/auth/circuit_final.zkey")

	if err != nil {
		return nil, errors.Wrap(err, "Error getting proving key")
	}

	circuitsPair := types.CircuitPair{
		Wasm:       wasm,
		ProvingKey: provingKey,
	}

	claimDetailsJson, err := instances.GetClaimDetailsJson(offer)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting claim details")
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
		return nil, errors.Wrap(err, "Error creating token")
	}

	_, err = token.Prove(circuitsPair.ProvingKey, circuitsPair.Wasm)

	if err != nil {
		return nil, errors.Wrap(err, "Error proving token")
	}

	zkProofRaw, err := json.Marshal(token.ZkProof)

	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling ZK proof")
	}

	return zkProofRaw, nil
}

func TestConnector(t *testing.T) {
	issuerApi := "https://issuer.polygon.robotornot.mainnet-beta.rarimo.com"
	claimType := "urn:uuid:6dff4518-5177-4f39-af58-9c156d9b6309"

	identityConfigJson, err := json.Marshal(identityConfig)

	if err != nil {
		t.Errorf("Error marshalling identity config: %v", err)
	}

	identityInstance, err := getIdentity(identityConfigJson)

	if err != nil {
		t.Errorf("Error getting identity instance: %v", err)
	}

	offer := types.ClaimOffer{}

	t.Run("Should get offer json", func(t *testing.T) {
		offerJson, err := GetOfferJson(issuerApi, identityInstance.DID.String(), claimType)

		if err != nil {
			t.Errorf("Error getting offer json: %v", err)
		}

		err = json.Unmarshal(offerJson, &offer)

		if err != nil {
			t.Errorf("Error unmarshalling offer json: %v", err)
		}
	})
	t.Run("Should get auth v2 inputs", func(t *testing.T) {
		offerJson, err := json.Marshal(offer)

		if err != nil {
			t.Errorf("Error marshalling offer: %v", err)
		}

		authV2InputsJson, err := GetAuthV2Inputs(identityConfigJson, offerJson)

		if err != nil {
			t.Errorf("Error getting auth v2 inputs: %v", err)
		}

		println(string(authV2InputsJson))
	})
	t.Run("Should get VC", func(t *testing.T) {
		/* IMITATE GROTH16 AUTHV2 PROVE, should be done in mobile device */
		zkProofRaw, err := getGroth16AuthV2ZKProof(*identityInstance, offer)

		if err != nil {
			t.Errorf("Error getting Groth16 AuthV2 ZK Proof: %v", err)
		}

		offerJson, err := json.Marshal(offer)

		if err != nil {
			t.Errorf("Error marshalling offer: %v", err)
		}

		vc, err := GetVC(identityConfigJson, offerJson, zkProofRaw)

		if err != nil {
			t.Errorf("Error getting VC: %v", err)
		}

		println(string(vc))

	})
	t.Run("Should get atomic query mtp v2 on-chain inputs", func(t *testing.T) {
		vcB, err := getFile("./internal/zkp/mocks/vc.json")

		if err != nil {
			t.Errorf("Error getting file: %v", err)
		}

		// Test the GetIdentity function
		atomicQueryMTVV2OnChainInputs, err := GetAtomicQueryMTVV2OnChainInputs(
			identityConfigJson,
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
