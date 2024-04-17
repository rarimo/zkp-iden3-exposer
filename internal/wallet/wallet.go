package wallet

import (
	"encoding/hex"
	"github.com/decred/dcrd/bech32"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/internal/wallet/overrides"
	"strings"
)

type Wallet struct {
	PrivateKeyHex string
	PrivateKey    overrides.PrivKey
	PubKey        overrides.PubKey
	Address       string
}

func NewWallet(privateKeyHex string, addressPrefix string) (*Wallet, error) {
	sanitizedPrivateKey := strings.TrimPrefix(privateKeyHex, "0x")

	privateKeyBytes, err := hex.DecodeString(sanitizedPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "Error decoding private key")
	}

	privateKey := overrides.GenPrivKeyFromSecret(privateKeyBytes)

	pubKey := privateKey.PubKey()

	covertedPubKeyHashedRipemd160, err := bech32.ConvertBits(pubKey.Address(), 8, 5, true)

	if err != nil {
		return nil, errors.Wrap(err, "Error converting public key hashed ripemd160")
	}

	address, err := bech32.Encode(addressPrefix, covertedPubKeyHashedRipemd160)

	return &Wallet{
		PrivateKeyHex: privateKeyHex,
		PrivateKey:    *privateKey,
		PubKey: overrides.PubKey{
			Key: pubKey.Bytes(),
		},
		Address: address,
	}, nil
}

func (w *Wallet) GetAccounts() []Account {
	return []Account{
		{
			Algo:    "secp256k1",
			Address: w.Address,
			PubKey:  w.PubKey,
		},
	}
}

func (w *Wallet) SignDirect(signerAddress string, bytesToSign []byte) ([]byte, error) {
	accounts := w.GetAccounts()

	account := Account{}

	for _, a := range accounts {
		if a.Address == signerAddress {
			account = a
		}
	}

	if &account.Address == nil {
		return nil, errors.New("Signer address not found")
	}

	signature, err := w.PrivateKey.Sign(bytesToSign)
	if err != nil {
		return nil, errors.Wrap(err, "Error signing")
	}

	return signature, nil
}

//func (w *Wallet) SignAmino(
//	signerAddress string,
//	signDoc legacytx.StdSignDoc,
//) ([]byte, error) {
//	accounts := w.GetAccounts()
//
//	account := Account{}
//
//	for _, a := range accounts {
//		if a.Address == signerAddress {
//			account = a
//		}
//	}
//
//	if &account.Address == nil {
//		return nil, errors.New("Signer address not found")
//	}
//
//	if &account.PubKey == nil {
//		return nil, errors.New("Signer public key not found")
//	}
//}
