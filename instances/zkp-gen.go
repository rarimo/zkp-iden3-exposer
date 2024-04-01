package instances

import (
	"github.com/rarimo/zkp-iden3-exposer/types"
)

type ZkpGenConfig struct {
	ChainInfo types.ChainZkpInfo
}

type ZkpGenCommonInputs struct {
	CircuitClaim      types.CircuitClaim
	Query             types.Query
	NodeAuxNonRev     types.NodeAuxValue
	ClaimNonRevStatus types.RevocationStatus
	Value             []string
	timestamp         *int
}

type ZkpGen struct {
	Config   ZkpGenConfig
	Identity *Identity
}

//func (z *ZkpGen) GenerateProof() (jwz.Token, error) {}
