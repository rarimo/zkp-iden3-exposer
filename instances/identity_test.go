package instances

import "testing"

func TestNewIdentity(t *testing.T) {
	t.Run("Should create identity", func(t *testing.T) {
		_, err := NewIdentity(IdentityConfig{
			IdType: [2]byte{
				0x01,
				0x00,
			},
			SchemaHashHex: "cca3371a6cb1b715004407e325bd993c",
		}, nil)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("Should create identity from privateKeyHex", func(t *testing.T) {
		pkHex := "9a5305fa4c55cbf517c99693a7ec6766203c88feab50c944c00feec051d5dab7"
		didString := "did:iden3:readonly:tSpQ56dBXo3Druez8wAbTTqd9yV1K2q4TwFu2taQj"

		identity, err := NewIdentity(IdentityConfig{
			IdType: [2]byte{
				0x01,
				0x00,
			},
			SchemaHashHex: "cca3371a6cb1b715004407e325bd993c",
		}, &pkHex)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if identity.DidString() != didString {
			t.Errorf("Expected: %v, got: %v", didString, identity.DidString())
		}
	})

	// TODO: add throw error tests
}
