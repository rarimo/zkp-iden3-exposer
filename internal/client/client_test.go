package client

import (
	"github.com/rarimo/zkp-iden3-exposer/internal/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"testing"
	"time"
)

func TestClient_SubmitTx(t *testing.T) {
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

	// FIXME: {"txhash":"17CF651BB2C1FD80297ED454D51114A1F36F9DA73203D0C2850A4D094FBB4704","codespace":"sdk","code":2,"raw_log":"unable to resolve type URL /: tx parse error","logs":null,"gas_used":6666,"events":null}
	t.Run("Should sign and send TX", func(t *testing.T) {
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
