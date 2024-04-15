package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/rarimo/zkp-iden3-exposer/internal/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"testing"
	"time"
)

func TestClient_SubmitTx(t *testing.T) {
	pk := "1cbd5d2d1801e964736881fc0584473f23ba82669599ac65957fb4f2caf43e17"
	addressPrefix := "rarimo"

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

	t.Run("Should sign", func(t *testing.T) {
		client := Client{
			cli:      grpcClient,
			ChainId:  "rarimo_42-1",
			Prefix:   addressPrefix,
			GasLimit: 1000000,
			GasPrice: 0,
		}

		wallet, err := wallet.NewWallet(pk, addressPrefix)
		if err != nil {
			t.Errorf("Error creating wallet: %v", err)
		}

		// FIXME: panic: invalid Bech32 prefix; expected cosmos, got rarimo
		msgSend := types.NewMsgSend(
			sdk.MustAccAddressFromBech32(wallet.Address),
			sdk.MustAccAddressFromBech32("rarimo1apm2p4k97euu8k8lxg9974kxvfnah8zj7lnydf"),
			[]sdk.Coin{
				sdk.Coin{
					Denom:  "stake",
					Amount: sdk.NewInt(1000), // 1000000 = 1 stake, 1000 = 0.001 stake
				},
			},
		)

		txResp, err := client.SubmitTx(*wallet, msgSend)
		if err != nil {
			t.Errorf("Error submitting tx: %v", err)
		}

		println(string(txResp))
	})
}
