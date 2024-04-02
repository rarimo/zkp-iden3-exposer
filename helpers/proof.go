package helpers

import (
	"encoding/hex"
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/constants"
	"github.com/rarimo/zkp-iden3-exposer/contracts"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"math/big"
	"net/http"
)

func NewProofFromJson(proofJsonBytes []byte) (*merkletree.Proof, error) {
	proof := &merkletree.Proof{}

	err := proof.UnmarshalJSON(proofJsonBytes)

	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal proofJson")
	}

	return proof, nil
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
		return nil, errors.Wrap(err, "failed to get revocation status")
	}

	type TreeStateJson struct {
		State              string `json:"state"`
		ClaimsTreeRoot     string `json:"claimsTreeRoot"`
		RevocationTreeRoot string `json:"revocationTreeRoot"`
		RootOfRoots        string `json:"rootOfRoots"`
	}

	type RevStatusJson struct {
		Mtp    json.RawMessage `json:"mtp"`
		Issuer TreeStateJson   `json:"issuer"`
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

func CheckVCAndGetCoreClaim(vc types.W3CCredential) (*core.Claim, *core.Claim, error) {
	revStatus, err := GetRevocationStatus(vc.CredentialStatus.Id, nil)

	if err != nil {
		return nil, nil, err
	}

	if revStatus.Mtp.Existence {
		return nil, nil, errors.New("Credential is revoked")
	}

	if len(vc.Proof) == 0 {
		return nil, nil, errors.New("Proof is empty")
	}

	var sigProof = types.BJJSignatureProofRaw{}

	for _, proof := range vc.Proof {
		if proof != nil {
			if err := json.Unmarshal(proof, &sigProof); err != nil {
				continue
			}

			if sigProof.Type == string(constants.BJJSignature) {
				break
			}
		}
	}

	if &sigProof == nil {
		return nil, nil, errors.New("Signature proof is empty")
	}

	sigProofCoreClaim := core.Claim{}

	err = sigProofCoreClaim.FromHex(sigProof.CoreClaim)

	if err != nil {
		return nil, nil, err
	}

	var mtProof = types.Iden3SparseMerkleTreeProofRaw{}

	for _, proof := range vc.Proof {
		if proof != nil {
			if err := json.Unmarshal(proof, &mtProof); err != nil {
				continue
			}

			if mtProof.Type == string(constants.Iden3SparseMerkleTreeProof) {
				break
			}
		}
	}

	if &mtProof == nil {
		return nil, nil, errors.New("Signature proof is empty")
	}

	mtProofCoreClaim := core.Claim{}

	err = mtProofCoreClaim.FromHex(mtProof.CoreClaim)

	if err != nil {
		return nil, nil, err
	}

	return &sigProofCoreClaim, &mtProofCoreClaim, nil
}

//func PrepareNonMerklizedQuery(proofQuery types.ProofQuery, vc types.W3CCredential) (*circuits.Query, error) {
//}
//
//func PrepareMerklizedQuery(proofQuery types.ProofQuery, vc types.W3CCredential) (*circuits.Query, error) {
//}
//
//func ToCircuitsQuery(proofQuery types.ProofQuery, vc types.W3CCredential, coreClaim core.Claim) (*circuits.Query, error) {
//	// PrepareNonMerklizedQuery
//
//	// PrepareMerklizedQuery
//
//	mtPosition, err := coreClaim.GetMerklizedPosition()
//
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to get merklized position")
//	}
//
//	if mtPosition == core.MerklizedRootPositionNone {
//		return PrepareNonMerklizedQuery(proofQuery, vc)
//	}
//
//	return PrepareMerklizedQuery(proofQuery, vc)
//}
