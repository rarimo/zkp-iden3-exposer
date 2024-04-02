package constants

const (
	DefaultMTLevels                   = 40
	DefaultValueArraySize             = 64
	DefaultMTLevelsOnChain            = 64
	DefaultMTLevelsClaimsMerklization = 32
)

type ProofType string

const (
	BJJSignature               ProofType = "BJJSignature2021"
	Iden3SparseMerkleTreeProof ProofType = "Iden3SparseMerkleTreeProof"
)
