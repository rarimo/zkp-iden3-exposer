package types

import (
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql/v2"
)

// custom, TODO: mb can find a existed type
type RevocationStatus struct {
	Mtp    merkletree.Proof
	Issuer circuits.TreeState
}

// custom
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
	Id                string           `json:"id"`
	AtContext         []string         `json:"@context"`
	Type              []string         `json:"type"`
	ExpirationDate    *string          `json:"expirationDate"`
	IssuanceDate      *string          `json:"issuanceDate"`
	CredentialSubject json.RawMessage  `json:"credentialSubject"`
	CredentialStatus  CredentialStatus `json:"credentialStatus"`
	Issuer            string           `json:"issuer"`
	CredentialSchema  struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"credentialSchema"`
	/* tuple: [ BJJSignatureProofRaw, Iden3SparseMerkleTreeProofRaw ] */
	Proof []json.RawMessage `json:"proof"`
}

type AgentResponse struct {
	Body struct {
		Credential W3CCredential `json:"credential"`
	} `json:"body"`
}

// custom, TODO: mb can find a existed type
type NodeAuxValue struct {
	Key   merkletree.Hash
	Value merkletree.Hash
	NoAux string
}

// custom
type WrappedProof struct {
	Proof     merkletree.Proof
	TreeState circuits.TreeState
}

// custom
type SignatureProof struct {
	Signature             *babyjub.Signature
	IssuerAuthClaim       *core.Claim
	IssuerAuthIncProof    WrappedProof
	IssuerAuthNonRevProof WrappedProof
}

// custom
type CircuitClaim struct {
	IssuerId       string
	Claim          core.Claim
	SignatureProof *SignatureProof
	IncProof       *WrappedProof
}

// custom
type ProofQuery struct {
	AllowedIssuers      []string        `json:"allowedIssuers"`
	CredentialSubject   json.RawMessage `json:"credentialSubject"`
	Schema              string          `json:"schema"`
	ClaimId             string          `json:"claimId"`
	CredentialSubjectId string          `json:"credentialSubjectId"`
	Context             string          `json:"context"`
	Type                []string        `json:"type"`
}

// custom
type CreateProofRequest struct {
	id             *int
	accountAddress *string
	circuitId      circuits.CircuitID
	challenge      *string // bigint string
	query          ProofQuery
}
