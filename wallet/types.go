package wallet

import (
	"github.com/rarimo/zkp-iden3-exposer/wallet/overrides"
)

type Account struct {
	Algo    string           `json:"algo"`
	Address string           `json:"address"`
	PubKey  overrides.PubKey `json:"pubKey"`
}
