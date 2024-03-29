package types

type ChainZkpInfo struct {
	TargetChainId              int
	TargetRpcUrl               string
	TargetStateContractAddress string
	CoreApiUrl                 string
	CoreEvmRpcApiUrl           string
	CoreStateContractAddress   string
}

type CircuitPair struct {
	Wasm       []byte
	ProvingKey []byte
}
