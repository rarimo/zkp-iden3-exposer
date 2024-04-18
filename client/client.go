package client

import (
	"context"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	client "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pkg/errors"
	ethermint "github.com/rarimo/rarimo-core/ethermint/types"
	"github.com/rarimo/zkp-iden3-exposer/wallet"
	"google.golang.org/grpc"
)

type ChainConfig struct {
	ChainId     string `json:"chainId"`
	Denom       string `json:"denom"`
	MinGasPrice uint64 `json:"minGasPrice"`
	GasLimit    uint64 `json:"gasLimit"`
}

type Client struct {
	Cli         *grpc.ClientConn
	Signer      wallet.Wallet
	ChainConfig ChainConfig
}

func NewClient(cli *grpc.ClientConn, chainConfig ChainConfig, signer wallet.Wallet) (*Client, error) {
	return &Client{
		Cli:         cli,
		Signer:      signer,
		ChainConfig: chainConfig,
	}, nil
}

func (c *Client) submitTx(msgs ...sdk.Msg) ([]byte, error) {
	txConfig := authTx.NewTxConfig(
		codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		[]signing.SignMode{signing.SignMode_SIGN_MODE_DIRECT},
	)

	builder := txConfig.NewTxBuilder()
	err := builder.SetMsgs(msgs...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to set messages")
	}

	builder.SetGasLimit(c.ChainConfig.GasLimit)
	builder.SetFeeAmount(
		sdk.Coins{
			sdk.Coin{
				Denom:  c.ChainConfig.Denom,
				Amount: sdk.NewInt(int64(c.ChainConfig.GasLimit * c.ChainConfig.MinGasPrice)),
			},
		},
	)

	accountResp, err := authtypes.NewQueryClient(c.Cli).Account(
		context.TODO(),
		&authtypes.QueryAccountRequest{Address: c.Signer.Address},
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get account")
	}

	account := ethermint.EthAccount{}
	err = account.Unmarshal(accountResp.Account.Value)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal account")
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
		return nil, errors.Wrap(err, "Failed to set signature")
	}

	signerData := xauthsigning.SignerData{
		ChainID:       c.ChainConfig.ChainId,
		AccountNumber: account.AccountNumber,
		Sequence:      accountSequence,
	}

	// Generate the bytes to be signed.
	signBytes, err := txConfig.SignModeHandler().GetSignBytes(txConfig.SignModeHandler().DefaultMode(), signerData, builder.GetTx())
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get sign bytes")
	}

	// Sign those bytes
	signature, err := c.Signer.PrivateKey.Sign(signBytes)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to sign bytes")
	}

	// Construct the SignatureV2 struct
	sigData := signing.SingleSignatureData{
		SignMode:  txConfig.SignModeHandler().DefaultMode(),
		Signature: signature,
	}

	sigV2 := signing.SignatureV2{
		PubKey:   c.Signer.PrivateKey.PubKey(),
		Data:     &sigData,
		Sequence: accountSequence,
	}

	err = builder.SetSignatures(sigV2)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to set signature")
	}

	tx, err := txConfig.TxEncoder()(builder.GetTx())
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode tx")
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
