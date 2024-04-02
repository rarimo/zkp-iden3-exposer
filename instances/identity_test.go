package instances

import (
	"github.com/rarimo/zkp-iden3-exposer/types"
	"testing"
)

func getIdentity(pkHex *string) Identity {
	PK := pkHex

	if PK == nil {
		_pkHex := "9a5305fa4c55cbf517c99693a7ec6766203c88feab50c944c00feec051d5dab7"
		PK = &_pkHex
	}

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
	}, PK)

	return *identity
}

func TestNewIdentity(t *testing.T) {
	t.Run("Should create identity", func(t *testing.T) {
		identity := getIdentity(nil)

		if identity.DID.String() == "" {
			t.Errorf("Error: %v", "DID is empty")
		}
	})

	t.Run("Should create identity from privateKeyHex", func(t *testing.T) {
		pkHex := "9a5305fa4c55cbf517c99693a7ec6766203c88feab50c944c00feec051d5dab7"
		didString := "did:iden3:readonly:tSpQ56dBXo3Druez8wAbTTqd9yV1K2q4TwFu2taQj"

		identity := getIdentity(&pkHex)

		if identity.DID.String() != didString {
			t.Errorf("Expected: %v, got: %v", didString, identity.DID.String())
		}
	})

	// TODO: add throw error tests
}
