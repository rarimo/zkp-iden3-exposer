package wallet

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/decred/dcrd/bech32"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ripemd160"
	"strings"
)

type Wallet struct {
	PrivateKeyHex string
	PrivateKey    secp256k1.PrivateKey
	PubKey        secp256k1.PublicKey
	Address       string
}

func NewWallet(privateKeyHex string, addressPrefix string) (*Wallet, error) {
	sanitizedPrivateKey := strings.TrimPrefix(privateKeyHex, "0x")

	privateKeyBytes, err := hex.DecodeString(sanitizedPrivateKey)

	if err != nil {
		return nil, errors.Wrap(err, "Error decoding private key")
	}

	privateKey := secp256k1.PrivKeyFromBytes(privateKeyBytes)

	pubKey := privateKey.PubKey()

	pubKeyHasher := sha256.New()
	pubKeyHasher.Write(pubKey.SerializeCompressed())
	pubKeyHashed := pubKeyHasher.Sum(nil)

	pubKeyRipemd160Hasher := ripemd160.New()
	pubKeyRipemd160Hasher.Write(pubKeyHashed)
	pubKeyHashedRipemd160 := pubKeyRipemd160Hasher.Sum(nil)

	covertedPubKeyHashedRipemd160, err := bech32.ConvertBits(pubKeyHashedRipemd160, 8, 5, true)

	if err != nil {
		return nil, errors.Wrap(err, "Error converting public key hashed ripemd160")
	}

	address, err := bech32.Encode(addressPrefix, covertedPubKeyHashedRipemd160)

	return &Wallet{
		PrivateKeyHex: privateKeyHex,
		PrivateKey:    *privateKey,
		PubKey:        *pubKey,
		Address:       address,
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

	privateKey := w.PrivateKey.ToECDSA()

	signature, err := privateKey.Sign(rand.Reader, bytesToSign, nil)
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
