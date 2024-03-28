package types

import (
	"encoding/json"
	"github.com/rarimo/go-merkletree"
)

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

type ClaimOffer struct {
	Body struct {
		Credentials []struct {
			Description string `json:"description"`
			Id          string `json:"id"`
		} `json:"Credentials"`
		Url string `json:"url"`
	} `json:"body"`
	From     string  `json:"from"`
	Id       string  `json:"id"`
	ThreadID *string `json:"threadID"`
	To       string  `json:"to"`
	Typ      *string `json:"typ"`
	Type     string  `json:"type"`
}

type CredentialStatus struct {
	Id              string            `json:"id"`
	Type            string            `json:"type"`
	RevocationNonce int               `json:"revocationNonce"`
	StatusIssuer    *CredentialStatus `json:"statusIssuer"`
}

type BJJSignatureProofRaw struct {
	// FIXME: constants.BJJSignature
	Type       string `json:"type"`
	IssuerData struct {
		Id        string `json:"id"`
		UpdateUrl string `json:"updateUrl"`
		State     struct {
			ClaimsTreeRoot string `json:"claimsTreeRoot"`
			Value          string `json:"value"`
		} `json:"state"`
		AuthCoreClaim string `json:"authCoreClaim"`
		Mtp           struct {
			Existence bool     `json:"existence"`
			Siblings  []string `json:"siblings"`
		} `json:"mtp"`
		CredentialStatus struct {
			Id              string `json:"id"`
			RevocationNonce int    `json:"revocationNonce"`
			Type            string `json:"type"`
		} `json:"credentialStatus"`
	} `json:"issuerData"`
	CoreClaim string `json:"coreClaim"`
	Signature string `json:"signature"`
}

type Iden3SparseMerkleTreeProofRaw struct {
	Id string `json:"id"`
	//FIXME: ProofType.Iden3SparseMerkleTreeProof
	Type       string `json:"type"`
	IssuerData struct {
		Id        string `json:"id"`
		UpdateUrl string `json:"updateUrl"`
		State     struct {
			TxId               string `json:"txId"`
			BlockTimestamp     int    `json:"blockTimestamp"`
			BlockNumber        int    `json:"blockNumber"`
			RootOfRoots        string `json:"rootOfRoots"`
			ClaimsTreeRoot     string `json:"claimsTreeRoot"`
			RevocationTreeRoot string `json:"revocationTreeRoot"`
			Value              string `json:"value"`
		} `json:"state"`
	} `json:"issuerData"`
	CoreClaim string `json:"coreClaim"`
	Mtp       struct {
		Existence bool     `json:"existence"`
		Siblings  []string `json:"siblings"`
	} `json:"mtp"`
}

type W3CCredential struct {
	Id             string   `json:"id"`
	AtContext      []string `json:"@context"`
	Type           []string `json:"type"`
	ExpirationDate *string
	IssuanceDate   *string `json:"issuanceDate"`
	// FIXME: empty struct
	CredentialSubject struct{}         `json:"credentialSubject"`
	CredentialStatus  CredentialStatus `json:"credentialStatus"`
	Issuer            string
	CredentialSchema  struct {
		Id   string
		Type string
	}
	/* tuple: [ BJJSignatureProofRaw, Iden3SparseMerkleTreeProofRaw ] */
	Proof []json.RawMessage `json:"proof"`
}
