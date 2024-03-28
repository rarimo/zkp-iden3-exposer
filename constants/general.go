package constants

const (
	DefaultMTLevels                   = 40
	DefaultValueArraySize             = 64
	DefaultMTLevelsOnChain            = 64
	DefaultMTLevelsClaimsMerklization = 32
)

type Operators int

const (
	NOOP Operators = iota
	EQ
	LT
	GT
	IN
	NIN
	NE
)

var QueryOperators = map[string]Operators{
	"$noop": NOOP,
	"$eq":   EQ,
	"$lt":   LT,
	"$gt":   GT,
	"$in":   IN,
	"$nin":  NIN,
	"$ne":   NE,
}

type ProofType string

const (
	BJJSignature               ProofType = "BJJSignature2021"
	Iden3SparseMerkleTreeProof ProofType = "Iden3SparseMerkleTreeProof"
)
