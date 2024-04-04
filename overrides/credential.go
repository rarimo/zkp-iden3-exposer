package overrides

import "github.com/iden3/go-schema-processor/v2/verifiable"

type W3CCredential struct {
	//ID string `json:"id,omitempty"`
	//
	//Context           []string                    `json:"@context"`
	//Type              []string                    `json:"type"`
	//Expiration        *time.Time                  `json:"expirationDate,omitempty"`
	//IssuanceDate      *time.Time                  `json:"issuanceDate,omitempty"`
	//CredentialSubject map[string]interface{}      `json:"credentialSubject"`
	//CredentialStatus  interface{}                 `json:"credentialStatus,omitempty"`
	//Issuer            string                      `json:"issuer"`
	//CredentialSchema  verifiable.CredentialSchema `json:"credentialSchema"`
	//RefreshService    *verifiable.RefreshService  `json:"refreshService,omitempty"`
	//DisplayMethod     *verifiable.DisplayMethod   `json:"displayMethod,omitempty"`

	verifiable.W3CCredential

	Proof CredentialProofs `json:"proof"`
}
