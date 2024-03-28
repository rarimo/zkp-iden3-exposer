package helpers

import (
	"encoding/hex"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

func InitSK(skHex *string) (*babyjub.PrivateKey, error) {
	if skHex == nil {

		var pk = babyjub.NewRandPrivKey()

		return &pk, nil
	}

	var sk babyjub.PrivateKey

	_, err := hex.Decode(sk[:], []byte(*skHex))

	if err != nil {
		return nil, err
	}

	return &sk, nil
}

// TODO: ParseDid will come from the identity package
