package wallet

import "github.com/decred/dcrd/dcrec/secp256k1/v4"

type Account struct {
	Algo    string              `json:"algo"`
	Address string              `json:"address"`
	PubKey  secp256k1.PublicKey `json:"pubKey"`
}
