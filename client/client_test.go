package client

import (
	"github.com/rarimo/zkp-iden3-exposer/wallet"
	"testing"
)

func TestClient(t *testing.T) {
	pk := "1cbd5d2d1801e964736881fc0584473f23ba82669599ac65957fb4f2caf43e17"
	addressPrefix := "rarimo"

	w, err := wallet.NewWallet(pk, addressPrefix)
	if err != nil {
		t.Errorf("Error creating wallet: %v", err)
	}

	t.Run("Should Send tokens", func(t *testing.T) {
		client, err := NewClient(
			ChainConfig{
				ChainId:     "rarimo_42-1",
				Denom:       "stake",
				Addr:        "104.196.227.66:9090",
				MinGasPrice: 0,
				GasLimit:    1000000,
				TLS:         true,
			},
			*w,
		)

		txResp, err := client.Send(
			w.Address,
			"rarimo1apm2p4k97euu8k8lxg9974kxvfnah8zj7lnydf",
			1000, // 1000000 = 1 Stake, 1000 = 0.001 Stake
			"stake",
		)
		if err != nil {
			t.Errorf("Error sending tx: %v", err)
		}

		println(string(txResp))
	})
}
