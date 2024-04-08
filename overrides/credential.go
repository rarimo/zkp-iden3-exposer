package overrides

import "github.com/iden3/go-schema-processor/v2/verifiable"

type W3CCredential struct {
	verifiable.W3CCredential

	Proof CredentialProofs `json:"proof"`
}
