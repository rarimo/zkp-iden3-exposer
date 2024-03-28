package types

import "github.com/rarimo/go-merkletree"

type TreeState struct {
	State              merkletree.Hash
	RootOfRoots        merkletree.Hash
	ClaimsTreeRoot     merkletree.Hash
	RevocationTreeRoot merkletree.Hash
}

type RevocationStatus struct {
	Mtp    merkletree.Proof
	Issuer TreeState
}
