package types

import (
	"encoding/json"
	core "github.com/iden3/go-iden3-core/v2"
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

type GISTProof struct {
	Root  merkletree.Hash
	Proof merkletree.Proof
}

type NodeAuxValue struct {
	Key   merkletree.Hash
	Value merkletree.Hash
	NoAux string
}

type AuthV2CircuitInputs struct {
	GenesisID    string `json:"genesisID"`
	ProfileNonce string `json:"profileNonce"`

	AuthClaim    *core.Claim        `json:"authClaim"`
	AuthClaimMtp []*merkletree.Hash `json:"authClaimIncMtp"`

	AuthClaimNonRevMtp      []*merkletree.Hash `json:"authClaimNonRevMtp"`
	AuthClaimNonRevMtpAuxHi *merkletree.Hash   `json:"authClaimNonRevMtpAuxHi"`
	AuthClaimNonRevMtpAuxHv *merkletree.Hash   `json:"authClaimNonRevMtpAuxHv"`
	AuthClaimNonRevMtpNoAux string             `json:"authClaimNonRevMtpNoAux"`

	Challenge             string `json:"challenge"`
	ChallengeSignatureR8X string `json:"challengeSignatureR8x"`
	ChallengeSignatureR8Y string `json:"challengeSignatureR8y"`
	ChallengeSignatureS   string `json:"challengeSignatureS"`

	ClaimsTreeRoot *merkletree.Hash `json:"claimsTreeRoot"`
	RevTreeRoot    *merkletree.Hash `json:"revTreeRoot"`
	RootsTreeRoot  *merkletree.Hash `json:"rootsTreeRoot"`
	State          *merkletree.Hash `json:"state"`

	GISTRoot     *merkletree.Hash   `json:"gistRoot"`
	GISTMtp      []*merkletree.Hash `json:"gistMtp"`
	GISTMtpAuxHi *merkletree.Hash   `json:"gistMtpAuxHi"`
	GISTMtpAuxHv *merkletree.Hash   `json:"gistMtpAuxHv"`
	GISTMtpNoAux string             `json:"gistMtpNoAux"`
}
