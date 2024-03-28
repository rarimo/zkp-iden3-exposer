package helpers

import (
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rarimo/go-merkletree"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"net/http"
)

type NodeAuxJSON struct {
	Key   string
	Value string
}

type ProofJson struct {
	Existence bool         `json:"existence"`
	Siblings  []string     `json:"siblings"`
	NodeAux   *NodeAuxJSON `json:"node_aux"` // TODO: mb json could be NodeAux
}

type TreeStateJson struct {
	State              string `json:"state"`
	ClaimsTreeRoot     string `json:"claims_tree_root"`
	RevocationTreeRoot string `json:"revocation_tree_root"`
	RootOfRoots        string `json:"root_of_roots"`
}

type RevStatusJson struct {
	Mtp    ProofJson     `json:"mtp"`
	Issuer TreeStateJson `json:"issuer"`
}

func NewProofFromJson(proofJson ProofJson) (*merkletree.Proof, error) {
	Siblings := make([](*merkletree.Hash), len(proofJson.Siblings))

	for i, sibling := range proofJson.Siblings {
		siblingHash, err := merkletree.NewHashFromHex(sibling)

		if err != nil {
			return nil, err
		}

		Siblings[i] = siblingHash
	}

	NodeAux := merkletree.NodeAux{}

	key, err := merkletree.NewHashFromHex(proofJson.NodeAux.Key)

	if err != nil {
		return nil, err
	}

	value, err := merkletree.NewHashFromHex(proofJson.NodeAux.Value)

	if err != nil {
		return nil, err
	}

	if proofJson.NodeAux != nil {
		NodeAux = merkletree.NodeAux{
			Key:   key,
			Value: value,
		}
	}

	return &merkletree.Proof{
		Existence: proofJson.Existence,
		Siblings:  Siblings,
		NodeAux:   &NodeAux,
	}, nil
}

func BuildTreeState(
	state string,
	claimsTreeRoot string,
	revocationTreeRoot string,
	rootOfRoots string,
) (*types.TreeState, error) {
	State, err := merkletree.NewHashFromHex(state)

	if err != nil {
		return nil, err
	}

	ClaimsTreeRoot, err := merkletree.NewHashFromHex(claimsTreeRoot)

	if err != nil {
		return nil, err
	}

	RevocationTreeRoot, err := merkletree.NewHashFromHex(revocationTreeRoot)

	if err != nil {
		return nil, err
	}

	RootOfRoots, err := merkletree.NewHashFromHex(rootOfRoots)

	if err != nil {
		return nil, err
	}

	return &types.TreeState{
		State:              *State,
		ClaimsTreeRoot:     *ClaimsTreeRoot,
		RevocationTreeRoot: *RevocationTreeRoot,
		RootOfRoots:        *RootOfRoots,
	}, nil
}

func GetRevocationStatus(
	url string,
	endianSwappedCoreStateHash *string,
) (*types.RevocationStatus, error) {
	revStatusUrl := url

	if endianSwappedCoreStateHash != nil {
		revStatusUrl += "?state=" + *endianSwappedCoreStateHash
	}

	response, err := http.Get(revStatusUrl)

	if err != nil {
		return nil, err
	}

	var revStatusJson RevStatusJson

	if err := json.NewDecoder(response.Body).Decode(&revStatusJson); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	mtp, err := NewProofFromJson(revStatusJson.Mtp)

	if err != nil {
		return nil, err
	}

	issuer, err := BuildTreeState(
		revStatusJson.Issuer.State,
		revStatusJson.Issuer.ClaimsTreeRoot,
		revStatusJson.Issuer.RevocationTreeRoot,
		revStatusJson.Issuer.RootOfRoots,
	)

	if err != nil {
		return nil, err
	}

	return &types.RevocationStatus{
		Mtp:    *mtp,
		Issuer: *issuer,
	}, nil

}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func ConvertEndianSwappedCoreStateHashHex(hash string) (*string, error) {
	//const convertedStateHash = fromLittleEndian(
	//	Hex.decodeString(hash.slice(2)),
	//).toString(16);
	//
	//return convertedStateHash?.length < 64
	//? `0x0${convertedStateHash}`
	//: `0x${convertedStateHash}`;

	hexDecodedString, err := hex.DecodeString(hash[2:])

	if err != nil {
		return nil, err
	}

	// Reverse the byte order (little-endian to big-endian)
	reverseBytes(hexDecodedString)

	// Convert the reversed byte array back to a hex string
	convertedStateHash := hex.EncodeToString(hexDecodedString)

	// Add "0x" prefix if necessary
	if len(convertedStateHash) < 64 {
		result := "0x0" + convertedStateHash

		return &result, nil
	}

	result := "0x" + convertedStateHash

	return &result, nil
}
