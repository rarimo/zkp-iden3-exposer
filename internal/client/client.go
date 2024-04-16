package client

import (
	"context"
	"encoding/json"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	client "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pkg/errors"
	ethermint "github.com/rarimo/rarimo-core/ethermint/types"
	"github.com/rarimo/zkp-iden3-exposer/internal/wallet"
	"google.golang.org/grpc"
)

type Client struct {
	Cli      *grpc.ClientConn
	Signer   wallet.Wallet
	ChainId  string
	Prefix   string
	GasLimit int
	GasPrice int
}

func (c *Client) submitTx(msgs ...sdk.Msg) ([]byte, error) {
	txConfig := tx.NewTxConfig(
		codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		[]signing.SignMode{signing.SignMode_SIGN_MODE_DIRECT},
	)

	builder := txConfig.NewTxBuilder()
	err := builder.SetMsgs(msgs...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to set messages")
	}

	builder.SetGasLimit(uint64(c.GasLimit))
	builder.SetFeeAmount(
		sdk.Coins{
			sdk.Coin{
				Denom:  "stake",
				Amount: sdk.NewInt(int64(c.GasLimit * c.GasPrice)),
			},
		},
	)

	accountResp, err := authtypes.NewQueryClient(c.Cli).Account(
		context.TODO(),
		&authtypes.QueryAccountRequest{Address: c.Signer.Address},
	)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	account := ethermint.EthAccount{}
	err = account.Unmarshal(accountResp.Account.Value)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	accountSequence := account.GetSequence()

	err = builder.SetSignatures(signing.SignatureV2{
		PubKey: &c.Signer.PubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  txConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: accountSequence,
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	signerData := xauthsigning.SignerData{
		ChainID:       c.ChainId,
		AccountNumber: account.AccountNumber,
		Sequence:      accountSequence,
	}

	sigV2, err := clienttx.SignWithPrivKey(
		txConfig.SignModeHandler().DefaultMode(),
		signerData,
		builder,
		&c.Signer.PrivateKey,
		txConfig,
		accountSequence,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to sign tx")
	}

	err = builder.SetSignatures(sigV2)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to set signature")
	}

	tx, err := txConfig.TxEncoder()(builder.GetTx())
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	grpcRes, err := client.NewServiceClient(c.Cli).BroadcastTx(
		context.TODO(),
		&client.BroadcastTxRequest{
			Mode:    client.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: tx,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to broadcast tx")
	}

	data, err := json.Marshal(grpcRes.TxResponse)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal tx response")
	}

	return data, nil
}

func (c *Client) Send(addrFrom, addrTo string, amount int64, denom string) ([]byte, error) {
	msgSend := &bank.MsgSend{
		FromAddress: addrFrom,
		ToAddress:   addrTo,
		Amount: sdk.Coins{
			sdk.Coin{
				Denom:  denom,
				Amount: sdk.NewInt(amount),
			},
		},
	}

	txResp, err := c.submitTx(msgSend)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to submit tx")
	}

	return txResp, nil
}