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

type IdentityConfig struct {
	PkHex string `json:"pkHex"`

	IdType        [2]byte `json:"idType"`
	SchemaHashHex string  `json:"schemaHashHex"`

	TargetChainId              int    `json:"targetChainId"`
	TargetRpcUrl               string `json:"targetRpcUrl"`
	TargetStateContractAddress string `json:"targetStateContractAddress"`

	CoreApiUrl               string `json:"coreApiUrl"`
	CoreEvmRpcApiUrl         string `json:"coreEvmRpcApiUrl"`
	CoreStateContractAddress string `json:"coreStateContractAddress"`
}
