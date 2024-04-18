package zkp_iden3_exposer

import (
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	"github.com/iden3/go-jwz/v2"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/zkp/instances"
	"github.com/rarimo/zkp-iden3-exposer/zkp/types"
	"io"
	"os"
	"testing"
)

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
	wasm, err := getFile("./zkp/assets/circuits/auth/circuit.wasm")

	if err != nil {
		return nil, errors.Wrap(err, "Error getting wasm file")
	}

	provingKey, err := getFile("./zkp/assets/circuits/auth/circuit_final.zkey")

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

	connector := NewConnector(
		"1cbd5d2d1801e964736881fc0584473f23ba82669599ac65957fb4f2caf43e17",

		[]byte{1, 0},
		"cca3371a6cb1b715004407e325bd993c",

		11155111,
		"https://endpoints.omniatech.io/v1/eth/sepolia/public",
		"0x8a9F505bD8a22BF09b0c19F65C17426cd33f3912",

		"https://rpc-api.node1.mainnet-beta.rarimo.com",
		"https://rpc.evm.node1.mainnet-beta.rarimo.com",
		"0x753a8678c85d5fb70A97CFaE37c84CE2fD67EDE8",

		"rarimo_42-1",
		"rarimo",
		"stake",
		"104.196.227.66:9090",
		0,
		1000000,
		true,
	)

	identity, err := instances.NewIdentity(instances.IdentityConfig{
		IdType:        [2]byte(connector.IdType),
		SchemaHashHex: connector.SchemaHashHex,
		ChainInfo: types.ChainZkpInfo{
			TargetChainId:              connector.TargetChainId,
			TargetRpcUrl:               connector.TargetRpcUrl,
			TargetStateContractAddress: connector.TargetStateContractAddress,
			CoreApiUrl:                 connector.CoreApiUrl,
			CoreEvmRpcApiUrl:           connector.CoreEvmRpcApiUrl,
			CoreStateContractAddress:   connector.CoreStateContractAddress,
		},
	}, &connector.PkHex)
	if err != nil {
		t.Errorf("Error creating identity: %v", err)
	}

	offer := types.ClaimOffer{}

	walletAddress := ""

	t.Run("Should get offer json", func(t *testing.T) {
		offerJson, err := connector.GetOfferJson(issuerApi, identity.DID.String(), claimType)

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

		authV2InputsJson, err := connector.GetAuthV2Inputs(offerJson)

		if err != nil {
			t.Errorf("Error getting auth v2 inputs: %v", err)
		}

		println(string(authV2InputsJson))
	})
	t.Run("Should get VC", func(t *testing.T) {
		/* IMITATE GROTH16 AUTHV2 PROVE, should be done in mobile device */
		zkProofRaw, err := getGroth16AuthV2ZKProof(*identity, offer)
		if err != nil {
			t.Errorf("Error getting Groth16 AuthV2 ZK Proof: %v", err)
		}

		offerJson, err := json.Marshal(offer)
		if err != nil {
			t.Errorf("Error marshalling offer: %v", err)
		}

		vc, err := connector.GetVC(offerJson, zkProofRaw)

		if err != nil {
			t.Errorf("Error getting VC: %v", err)
		}

		println(string(vc))

	})
	t.Run("Should get atomic query mtp v2 on-chain inputs", func(t *testing.T) {
		vcB, err := getFile("./zkp/mocks/vc.json")
		if err != nil {
			t.Errorf("Error getting file: %v", err)
		}

		// Test the GetIdentity function
		atomicQueryMTVV2OnChainInputs, err := connector.GetAtomicQueryMTVV2OnChainInputs(
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
	t.Run("Should get wallet address", func(t *testing.T) {
		addr, err := connector.WalletGetAddress()
		if err != nil {
			t.Errorf("Error getting wallet address: %v", err)
		}

		walletAddress = addr

		println(walletAddress)
	})
	t.Run("Should send `MsgSend` tx", func(t *testing.T) {
		txResp, err := connector.WalletSend(
			walletAddress,
			"rarimo1apm2p4k97euu8k8lxg9974kxvfnah8zj7lnydf",
			1000, // 1000000 = 1 Stake, 1000 = 0.001 Stake
		)
		if err != nil {
			t.Errorf("Error sending tx: %v", err)
		}

		println(string(txResp))
	})
}
