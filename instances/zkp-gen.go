package instances

import (
	"github.com/iden3/go-circuits/v2"
	"github.com/rarimo/zkp-iden3-exposer/types"
)

type ZkpGenConfig struct {
	ChainInfo types.ChainZkpInfo
}

type ZkpGenCommonInputs struct {
	CircuitClaim      types.CircuitClaim
	Query             circuits.Query
	NodeAuxNonRev     types.NodeAuxValue
	ClaimNonRevStatus types.RevocationStatus
	Value             []string
	timestamp         *int
}

type ZkpGen struct {
	Config   ZkpGenConfig
	Identity *Identity
}

//func (z *ZkpGen) GenerateProof(
//	coreStateHash string,
//	operationGistHash string,
//) (jwz.Token, error) {
//}
