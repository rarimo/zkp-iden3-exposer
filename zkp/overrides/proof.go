package overrides

import (
	"encoding/json"
	"errors"
	core "github.com/iden3/go-iden3-core/v2"
	mt "github.com/iden3/go-merkletree-sql/v2"
	"github.com/iden3/go-schema-processor/v2/verifiable"
)

func validateHexCoreClaim(in string) error {
	var claim core.Claim
	err := claim.FromHex(in)
	return err
}

// Iden3SparseMerkleTreeProof JSON-LD structure
type Iden3SparseMerkleTreeProof struct {
	verifiable.Iden3SparseMerkleTreeProof

	ID string `json:"id"`
}

func (p *Iden3SparseMerkleTreeProof) UnmarshalJSON(in []byte) error {
	var obj struct {
		Type       verifiable.ProofType `json:"type"`
		IssuerData json.RawMessage      `json:"issuerData"`
		CoreClaim  string               `json:"coreClaim"`
		MTP        *mt.Proof            `json:"mtp"`
		ID         string               `json:"id"`
	}
	err := json.Unmarshal(in, &obj)
	if err != nil {
		return err
	}
	if obj.Type != verifiable.Iden3SparseMerkleTreeProofType {
		return errors.New("invalid proof type")
	}
	p.Type = obj.Type
	err = json.Unmarshal(obj.IssuerData, &p.IssuerData)
	if err != nil {
		return err
	}
	if err := validateHexCoreClaim(obj.CoreClaim); err != nil {
		return err
	}
	p.CoreClaim = obj.CoreClaim
	p.MTP = obj.MTP
	p.ID = obj.ID
	return nil
}

func (p *Iden3SparseMerkleTreeProof) ProofType() verifiable.ProofType {
	return p.Type
}

func (p *Iden3SparseMerkleTreeProof) GetCoreClaim() (*core.Claim, error) {
	var coreClaim core.Claim
	err := coreClaim.FromHex(p.CoreClaim)
	return &coreClaim, err
}

type CredentialProofs []verifiable.CredentialProof

type jsonObj = map[string]any

func reUnmarshalFromObj(obj jsonObj, v interface{}) error {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return json.Unmarshal(objBytes, v)
}

func extractProof(proof any) (verifiable.CredentialProof, error) {
	proofJ, ok := proof.(jsonObj)
	if !ok {
		return nil, errors.New("proof is not an object")
	}
	proofType, ok := proofJ["type"].(string)
	if !ok {
		return nil, errors.New("proof type is not specified")
	}

	switch verifiable.ProofType(proofType) {
	case verifiable.BJJSignatureProofType:
		var proof verifiable.BJJSignatureProof2021
		err := reUnmarshalFromObj(proofJ, &proof)
		return &proof, err
	case verifiable.Iden3SparseMerkleTreeProofType:
		var proof Iden3SparseMerkleTreeProof
		err := reUnmarshalFromObj(proofJ, &proof)
		return &proof, err
	default:
		var proof verifiable.CommonProof
		err := reUnmarshalFromObj(proofJ, &proof)
		return &proof, err
	}
}

func (cps *CredentialProofs) UnmarshalJSON(bs []byte) error {
	var i interface{}
	err := json.Unmarshal(bs, &i)
	if err != nil {
		return err
	}
	switch p := i.(type) {
	case []interface{}:
		for _, proofI := range p {
			proof, err := extractProof(proofI)
			if err != nil {
				return err
			}
			*cps = append(*cps, proof)
		}
	case interface{}:
		proof, err := extractProof(p)
		if err != nil {
			return err
		}
		*cps = append(*cps, proof)
	default:
		return errors.New("proof is not an array or an object")
	}
	return nil
}
