package overrides

import (
	"bytes"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/decred/dcrd/crypto/ripemd160"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"
	"hash"
	"io"
	"math/big"
)

var (
	_              cryptotypes.PrivKey  = &PrivKey{}
	_              codec.AminoMarshaler = &PrivKey{}
	oneInitializer                      = []byte{0x01}
)

const (
	PrivKeySize = 32
	keyType     = "secp256k1"
	PrivKeyName = "tendermint/PrivKeySecp256k1"
	PubKeyName  = "tendermint/PubKeySecp256k1"
)

func (*PrivKey) XXX_MessageName() string {
	return "cosmos.crypto.secp256k1.PrivKey"
	//return "cosmos.crypto.secp256r1.PrivKey"
}

// mac returns an HMAC of the given key and message.
func mac(alg func() hash.Hash, k, m []byte) []byte {
	h := hmac.New(alg, k)
	h.Write(m)
	return h.Sum(nil)
}

// https://tools.ietf.org/html/rfc6979#section-2.3.3
func int2octets(v *big.Int, rolen int) []byte {
	out := v.Bytes()

	// left pad with zeros if it's too short
	if len(out) < rolen {
		out2 := make([]byte, rolen)
		copy(out2[rolen-len(out):], out)
		return out2
	}

	// drop most significant bytes if it's too long
	if len(out) > rolen {
		out2 := make([]byte, rolen)
		copy(out2, out[len(out)-rolen:])
		return out2
	}

	return out
}

// hashToInt converts a hash value to an integer. There is some disagreement
// about how this is done. [NSA] suggests that this is done in the obvious
// manner, but [SECG] truncates the hash to the bit-length of the curve order
// first. We follow [SECG] because that's what OpenSSL does. Additionally,
// OpenSSL right shifts excess bits from the number if the hash is too large
// and we mirror that too.
// This is borrowed from crypto/ecdsa.
func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

// https://tools.ietf.org/html/rfc6979#section-2.3.4
func bits2octets(in []byte, curve elliptic.Curve, rolen int) []byte {
	z1 := hashToInt(in, curve)
	z2 := new(big.Int).Sub(z1, curve.Params().N)
	if z2.Sign() < 0 {
		return int2octets(z1, rolen)
	}
	return int2octets(z2, rolen)
}

// nonceRFC6979 generates an ECDSA nonce (`k`) deterministically according to RFC 6979.
// It takes a 32-byte hash as an input and returns 32-byte nonce to be used in ECDSA algorithm.
func nonceRFC6979(privkey *big.Int, hash []byte) *big.Int {

	curve := secp256k1.S256()
	q := curve.Params().N
	x := privkey
	alg := sha256.New

	qlen := q.BitLen()
	holen := alg().Size()
	rolen := (qlen + 7) >> 3
	bx := append(int2octets(x, rolen), bits2octets(hash, curve, rolen)...)

	// Step B
	v := bytes.Repeat(oneInitializer, holen)

	// Step C (Go zeroes the all allocated memory)
	k := make([]byte, holen)

	// Step D
	k = mac(alg, k, append(append(v, 0x00), bx...))

	// Step E
	v = mac(alg, k, v)

	// Step F
	k = mac(alg, k, append(append(v, 0x01), bx...))

	// Step G
	v = mac(alg, k, v)

	// Step H
	for {
		// Step H1
		var t []byte

		// Step H2
		for len(t)*8 < qlen {
			v = mac(alg, k, v)
			t = append(t, v...)
		}

		// Step H3
		secret := hashToInt(t, curve)
		if secret.Cmp(one) >= 0 && secret.Cmp(q) < 0 {
			return secret
		}
		k = mac(alg, k, append(v, 0x00))
		v = mac(alg, k, v)
	}
}

// signRFC6979 generates a deterministic ECDSA signature according to RFC 6979 and BIP 62.
func signRFC6979(privateKey *secp256k1.PrivateKey, hash []byte) (r, s *big.Int, err error) {
	privkey := privateKey.ToECDSA()
	N := secp256k1.S256().N
	//halfOrder := secp256k1.S256().halfOrder
	halfOrder := new(big.Int).Rsh(N, 1)
	k := nonceRFC6979(privkey.D, hash)
	inv := new(big.Int).ModInverse(k, N)
	//r, _ := privkey.Curve.ScalarBaseMult(k.Bytes())
	r, _ = privkey.Curve.ScalarBaseMult(k.Bytes())
	r.Mod(r, N)

	if r.Sign() == 0 {
		return nil, nil, errors.New("calculated R is zero")
	}

	e := hashToInt(hash, privkey.Curve)
	//s := new(big.Int).Mul(privkey.D, r)
	s = new(big.Int).Mul(privkey.D, r)
	s.Add(s, e)
	s.Mul(s, inv)
	s.Mod(s, N)

	if s.Cmp(halfOrder) == 1 {
		s.Sub(N, s)
	}
	if s.Sign() == 0 {
		return nil, nil, errors.New("calculated S is zero")
	}
	return r, s, nil
}

func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	pk := secp256k1.PrivKeyFromBytes(privKey.Key)

	r, s, err := signRFC6979(pk, crypto.Sha256(msg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	return serializeSig(r, s), nil
}

// Serialize signature to R || S.
// R, S are padded to 32 bytes respectively.
func serializeSig(r, s *big.Int) []byte {
	rBytes := r.Bytes()
	sBytes := s.Bytes()
	sigBytes := make([]byte, 64)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[32-len(rBytes):32], rBytes)
	copy(sigBytes[64-len(sBytes):64], sBytes)
	return sigBytes
}

// Bytes returns the byte representation of the Private Key.
func (privKey *PrivKey) Bytes() []byte {
	return privKey.Key
}

// PubKey performs the point-scalar multiplication from the privKey on the
// generator point to get the pubkey.
func (privKey *PrivKey) PubKey() cryptotypes.PubKey {
	pk := secp256k1.PrivKeyFromBytes(privKey.Key)
	pubKey := pk.PubKey()

	return &PubKey{Key: pubKey.SerializeCompressed()}
}

// Equals - you probably don't need to use this.
// Runs in constant time based on length of the
func (privKey *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	return privKey.Type() == other.Type() && subtle.ConstantTimeCompare(privKey.Bytes(), other.Bytes()) == 1
}

func (privKey *PrivKey) Type() string {
	return keyType
}

// MarshalAmino overrides Amino binary marshalling.
func (privKey PrivKey) MarshalAmino() ([]byte, error) {
	return privKey.Key, nil
}

// UnmarshalAmino overrides Amino binary marshalling.
func (privKey *PrivKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PrivKeySize {
		return fmt.Errorf("invalid privkey size")
	}
	privKey.Key = bz

	return nil
}

// MarshalAminoJSON overrides Amino JSON marshalling.
func (privKey PrivKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return privKey.MarshalAmino()
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (privKey *PrivKey) UnmarshalAminoJSON(bz []byte) error {
	return privKey.UnmarshalAmino(bz)
}

// GenPrivKey generates a new ECDSA private key on curve secp256k1 private key.
// It uses OS randomness to generate the private key.
func GenPrivKey() *PrivKey {
	return &PrivKey{Key: genPrivKey(crypto.CReader())}
}

// genPrivKey generates a new secp256k1 private key using the provided reader.
func genPrivKey(rand io.Reader) []byte {
	var privKeyBytes [PrivKeySize]byte
	d := new(big.Int)
	for {
		privKeyBytes = [PrivKeySize]byte{}
		_, err := io.ReadFull(rand, privKeyBytes[:])
		if err != nil {
			panic(err)
		}

		d.SetBytes(privKeyBytes[:])
		// break if we found a valid point (i.e. > 0 and < N == curverOrder)
		isValidFieldElement := 0 < d.Sign() && d.Cmp(secp256k1.S256().N) < 0
		if isValidFieldElement {
			break
		}
	}

	return privKeyBytes[:]
}

var one = new(big.Int).SetInt64(1)

// GenPrivKeyFromSecret hashes the secret with SHA2, and uses
// that 32 byte output to create the private key.
//
// It makes sure the private key is a valid field element by setting:
//
// c = sha256(secret)
// k = (c mod (n − 1)) + 1, where n = curve order.
//
// NOTE: secret should be the output of a KDF like bcrypt,
// if it's derived from user input.
func GenPrivKeyFromSecret(secret []byte) *PrivKey {
	//pk := secp256k1.PrivKeyFromBytes(secret)

	return &PrivKey{Key: secret}

	//secHash := sha256.Sum256(secret)
	//// to guarantee that we have a valid field element, we use the approach of:
	//// "Suite B Implementer’s Guide to FIPS 186-3", A.2.1
	//// https://apps.nsa.gov/iaarchive/library/ia-guidance/ia-solutions-for-classified/algorithm-guidance/suite-b-implementers-guide-to-fips-186-3-ecdsa.cfm
	//// see also https://github.com/golang/go/blob/0380c9ad38843d523d9c9804fe300cb7edd7cd3c/src/crypto/ecdsa/ecdsa.go#L89-L101
	//fe := new(big.Int).SetBytes(secHash[:])
	//n := new(big.Int).Sub(secp256k1.S256().N, one)
	//fe.Mod(fe, n)
	//fe.Add(fe, one)
	//
	//feB := fe.Bytes()
	//privKey32 := make([]byte, PrivKeySize)
	//// copy feB over to fixed 32 byte privKey32 and pad (if necessary)
	//copy(privKey32[32-len(feB):32], feB)
	//
	//return &PrivKey{Key: privKey32}
}

//-------------------------------------

var (
	_ cryptotypes.PubKey   = &PubKey{}
	_ codec.AminoMarshaler = &PubKey{}
)

// PubKeySize is comprised of 32 bytes for one field element
// (the x-coordinate), plus one byte for the parity of the y-coordinate.
const PubKeySize = 33

func (*PubKey) XXX_MessageName() string {
	//return "cosmos.crypto.secp256r1.PubKey"
	return "cosmos.crypto.secp256k1.PubKey"
}

func (m *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	//TODO implement me
	panic("implement me")
}

// Address returns a Bitcoin style addresses: RIPEMD160(SHA256(pubkey))
func (pubKey *PubKey) Address() crypto.Address {
	if len(pubKey.Key) != PubKeySize {
		panic("length of pubkey is incorrect")
	}

	sha := sha256.Sum256(pubKey.Key)
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha[:]) // does not error
	return hasherRIPEMD160.Sum(nil)
}

// Bytes returns the pubkey byte format.
func (pubKey *PubKey) Bytes() []byte {
	return pubKey.Key
}

func (pubKey *PubKey) String() string {
	return fmt.Sprintf("PubKeySecp256k1{%X}", pubKey.Key)
}

func (pubKey *PubKey) Type() string {
	return keyType
}

func (pubKey *PubKey) Equals(other cryptotypes.PubKey) bool {
	return pubKey.Type() == other.Type() && bytes.Equal(pubKey.Bytes(), other.Bytes())
}

// MarshalAmino overrides Amino binary marshalling.
func (pubKey PubKey) MarshalAmino() ([]byte, error) {
	return pubKey.Key, nil
}

// UnmarshalAmino overrides Amino binary marshalling.
func (pubKey *PubKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PubKeySize {
		return errors.Wrap(sdkErrors.ErrInvalidPubKey, "invalid pubkey size")
	}
	pubKey.Key = bz

	return nil
}

// MarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey PubKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return pubKey.MarshalAmino()
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey *PubKey) UnmarshalAminoJSON(bz []byte) error {
	return pubKey.UnmarshalAmino(bz)
}
