package client

import (
	"github.com/rarimo/zkp-iden3-exposer/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	pk := "1cbd5d2d1801e964736881fc0584473f23ba82669599ac65957fb4f2caf43e17"
	addressPrefix := "rarimo"

	w, err := wallet.NewWallet(pk, addressPrefix)
	if err != nil {
		t.Errorf("Error creating wallet: %v", err)
	}

	grpcClient, err := grpc.Dial(
		"104.196.227.66:9090",
		//"rpc-api.node1.mainnet-beta.rarimo.com:443",
		//"rpc.node1.mainnet-beta.rarimo.com:443",
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second, // wait time before ping if no activity
			Timeout: 20 * time.Second, // ping timeout
		}),
	)
	if err != nil {
		t.Errorf("Error dialing grpc: %v", err)
	}

	t.Run("Should Send tokens", func(t *testing.T) {
		client := Client{
			Cli:      grpcClient,
			Signer:   *w,
			ChainId:  "rarimo_42-1",
			Prefix:   addressPrefix,
			GasLimit: 1000000,
			GasPrice: 0,
		}

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
