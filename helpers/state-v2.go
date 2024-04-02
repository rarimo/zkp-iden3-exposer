package helpers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/zkp-iden3-exposer/contracts"
	"math/big"
)

func GetGISTProof(coreEvmRpcUrl string, coreStateContractAddress string, userId *big.Int, rootHash *big.Int) (*contracts.IStateGistProof, error) {
	ethClient, err := ethclient.Dial(coreEvmRpcUrl)

	stateV2Caller, err := contracts.NewStateV2Caller(common.HexToAddress(coreStateContractAddress), ethClient)

	if err != nil {
		return nil, err
	}

	if rootHash != nil {
		gistProof, err := stateV2Caller.GetGISTProofByRoot(nil, userId, rootHash)

		if err != nil {
			return nil, err
		}

		return &gistProof, nil
	}

	gistProof, err := stateV2Caller.GetGISTProof(&bind.CallOpts{}, userId)

	if err != nil {
		return nil, err
	}

	return &gistProof, nil
}
