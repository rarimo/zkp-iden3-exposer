package helpers

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/iden3/go-circuits/v2"
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/iden3/go-schema-processor/v2/loaders"
	"github.com/iden3/go-schema-processor/v2/merklize"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/contracts"
	"github.com/rarimo/zkp-iden3-exposer/overrides"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"math/big"
	"net/http"
)

func GetRevocationStatus(url string, EndianSwappedCoreStateHash *string) (verifiable.RevocationStatus, error) {
	revStatusUrl := url

	if EndianSwappedCoreStateHash != nil {
		revStatusUrl += "?state=" + *EndianSwappedCoreStateHash
	}

	response, err := http.Get(revStatusUrl)

	if err != nil {
		return verifiable.RevocationStatus{}, errors.Wrap(err, "failed to get revocation status")
	}

	var revStatus verifiable.RevocationStatus

	if err := json.NewDecoder(response.Body).Decode(&revStatus); err != nil {
		return verifiable.RevocationStatus{}, errors.Wrap(err, "failed to unmarshal")
	}

	return revStatus, nil
}

type CredentialStatusResolver struct {
	Url                        string
	EndianSwappedCoreStateHash *string
}

func (c *CredentialStatusResolver) Resolve(ctx context.Context, credentialStatus verifiable.CredentialStatus) (verifiable.RevocationStatus, error) {
	return GetRevocationStatus(c.Url, c.EndianSwappedCoreStateHash)
}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func ConvertEndianSwappedCoreStateHashHex(hash string) (*string, error) {
	hexDecodedString, err := hex.DecodeString(hash[2:])

	if err != nil {
		return nil, err
	}

	fromLE := FromLittleEndian(hexDecodedString)

	// Convert the reversed byte array back to a hex string
	convertedStateHash := hex.EncodeToString(fromLE.Bytes())

	// Add "0x" prefix if necessary
	if len(convertedStateHash) < 64 {
		result := "0x0" + convertedStateHash

		return &result, nil
	}

	result := "0x" + convertedStateHash

	return &result, nil
}

func FromLittleEndian(bytes []byte) *big.Int {
	n256 := big.NewInt(256)
	result := big.NewInt(0)
	base := big.NewInt(1)

	for _, b := range bytes {
		byteBigInt := big.NewInt(int64(b))
		result.Add(result, new(big.Int).Mul(base, byteBigInt))
		base.Mul(base, n256)
	}

	return result
}

func FromBigEndian(bytes []byte) *big.Int {
	reverseBytes(bytes)

	return FromLittleEndian(bytes)
}

func BuildTreeState(
	state string,
	claimsTreeRoot string,
	revocationTreeRoot string,
	rootOfRoots string,
) (*circuits.TreeState, error) {
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

	return &circuits.TreeState{
		State:          State,
		ClaimsRoot:     ClaimsTreeRoot,
		RevocationRoot: RevocationTreeRoot,
		RootOfRoots:    RootOfRoots,
	}, nil
}

func PrepareSiblingsStr(proof merkletree.Proof, levels int) []*merkletree.Hash {
	//siblings := proof.Siblings
	siblings := proof.AllSiblings()

	// Add the rest of empty levels to the siblings
	for i := len(siblings); i < levels; i++ {
		siblings = append(siblings, &merkletree.HashZero)
	}

	return siblings
}

func ToGISTProof(gistProofRaw contracts.IStateGistProof) (*circuits.GISTProof, error) {
	var existence = gistProofRaw.Existence
	nodeAux := merkletree.NodeAux{}

	if existence == false && gistProofRaw.AuxExistence == true {
		nodeAuxKey, err := merkletree.NewHashFromBigInt(gistProofRaw.AuxIndex)

		if err != nil {
			return nil, errors.Wrap(err, "failed to create nodeAuxKey")
		}

		nodeAux.Key = nodeAuxKey

		nodeAuxValue, err := merkletree.NewHashFromBigInt(gistProofRaw.AuxValue)

		if err != nil {
			return nil, errors.Wrap(err, "failed to create nodeAuxValue")
		}

		nodeAux.Value = nodeAuxValue
	}

	allSiblings := make([]*merkletree.Hash, len(gistProofRaw.Siblings))

	for i, sibling := range gistProofRaw.Siblings {
		siblingHash, err := merkletree.NewHashFromBigInt(sibling)

		if err != nil {
			return nil, errors.Wrap(err, "failed to create sibling")
		}

		allSiblings[i] = siblingHash
	}

	gistProof := &circuits.GISTProof{}

	gistProofProof, err := merkletree.NewProofFromData(
		existence,
		allSiblings,
		&nodeAux,
	)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create proof")
	}

	gistProof.Proof = gistProofProof

	gistProofRoot, err := merkletree.NewHashFromBigInt(gistProofRaw.Root)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create root")
	}

	gistProof.Root = gistProofRoot

	return gistProof, nil
}

func ConvertProofRequestToCircuitQuery(vc *overrides.W3CCredential, request *types.CreateProofRequest) (*circuits.Query, error) {
	value, ok := new(big.Int).SetString(request.Query.SubjectFieldValue, 10)

	if !ok {
		return nil, errors.New("failed to parse value")
	}

	query := circuits.Query{
		Operator:  request.Query.Operator,
		Values:    []*big.Int{value},
		SlotIndex: 0,
	}

	vcCopy := *vc

	vcCopy.Proof = nil

	credentialJson, err := json.Marshal(vcCopy)

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal vcCopy")
	}

	merklizer, err := merklize.MerklizeJSONLD(nil, bytes.NewReader(credentialJson))

	if err != nil {
		return nil, errors.Wrap(err, "failed to merklize")
	}

	docLoader := loaders.NewDocumentLoader(nil, "", loaders.WithHTTPClient(&http.Client{}))

	remoteDocument, err := docLoader.LoadDocument(vc.Context[2])

	if err != nil {
		return nil, errors.Wrap(err, "failed to load remote document")
	}

	marshaledDocument, err := json.Marshal(remoteDocument.Document)

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal document")
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to load remote document")
	}

	path, err := merklize.NewFieldPathFromContext(marshaledDocument, vc.Type[1], request.Query.SubjectFieldName)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create field path")
	}

	err = path.Prepend("https://www.w3.org/2018/credentials#credentialSubject")

	if err != nil {
		return nil, errors.Wrap(err, "failed to prepend path")
	}

	proof, proofValue, err := merklizer.Proof(nil, path)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create proof")
	}

	pathKey, err := path.MtEntry()

	if err != nil {
		return nil, fmt.Errorf("error getting path key: %v", err)
	}

	mtEntry, err := proofValue.MtEntry()

	if err != nil {
		return nil, fmt.Errorf("error getting mt entry: %v", err)
	}

	var siblings []*merkletree.Hash

	for _, sibling := range proof.AllSiblings() {
		siblingText, err := sibling.MarshalText()
		if err != nil {
			return nil, fmt.Errorf("error marshaling sibling: %v", err)
		}

		newSibling := merkletree.Hash{}
		if err := newSibling.UnmarshalText(siblingText); err != nil {
			return nil, fmt.Errorf("error unmarshaling sibling: %v", err)
		}

		siblings = append(siblings, &newSibling)
	}

	keyHash, err := merkletree.NewHashFromBigInt(pathKey)

	if err != nil {
		return nil, fmt.Errorf("error creating key hash: %v", err)
	}

	valueHash, err := merkletree.NewHashFromBigInt(mtEntry)

	if err != nil {
		return nil, fmt.Errorf("error creating value hash: %v", err)
	}

	valueProofMTP, err := merkletree.NewProofFromData(
		proof.Existence,
		siblings,
		&merkletree.NodeAux{
			Key:   keyHash,
			Value: valueHash,
		},
	)

	query.ValueProof = &circuits.ValueProof{
		Path:  pathKey,
		MTP:   valueProofMTP,
		Value: mtEntry,
	}

	return &query, nil
}
