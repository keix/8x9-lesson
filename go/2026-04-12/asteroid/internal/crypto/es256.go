package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"
)

// ES256Generator generates ECDSA key pairs for ES256 algorithm.
type ES256Generator struct{}

// ES256Signer signs payloads using ECDSA with SHA-256.
type ES256Signer struct{}

// NewES256Generator creates a new ES256 key generator.
func NewES256Generator() *ES256Generator {
	return &ES256Generator{}
}

// NewES256Signer creates a new ES256 signer.
func NewES256Signer() *ES256Signer {
	return &ES256Signer{}
}

// Generate creates a new ECDSA key pair using P-256 curve.
func (g *ES256Generator) Generate() (*KeyPair, error) {
	// Generate ECDSA private key with P-256 curve
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKeyGenerationFailed, err)
	}

	// Generate Key ID from public key
	kid := GenerateKIDFromECDSAPublicKey(&privateKey.PublicKey)

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
		Algorithm:  g.Alg(),
		KeyID:      kid,
		CreatedAt:  time.Now(),
	}, nil
}

// Alg returns the algorithm identifier.
func (g *ES256Generator) Alg() string {
	return "ES256"
}

// Sign creates an ECDSA signature for the payload.
func (s *ES256Signer) Sign(payload []byte, kp *KeyPair) ([]byte, error) {
	// Type assert to ECDSA private key
	privateKey, ok := kp.PrivateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, ErrInvalidKeyType
	}

	// Hash the payload with SHA-256
	hash := sha256.Sum256(payload)

	// Sign using ECDSA (ASN.1 DER format)
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSigningFailed, err)
	}

	return signature, nil
}

// SignJWT signs a JWT payload and returns the signature in JWS format.
// JWS requires the signature in R||S format (concatenated, not ASN.1).
func (s *ES256Signer) SignJWT(payload []byte, kp *KeyPair) ([]byte, error) {
	// Type assert to ECDSA private key
	privateKey, ok := kp.PrivateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, ErrInvalidKeyType
	}

	// Hash the payload with SHA-256
	hash := sha256.Sum256(payload)

	// Sign using ECDSA
	r, ss, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSigningFailed, err)
	}

	// Convert to JWS format (R||S, each 32 bytes for P-256)
	const sigSize = 32
	signature := make([]byte, sigSize*2)
	rBytes := r.Bytes()
	sBytes := ss.Bytes()

	// Pad R and S to 32 bytes each
	copy(signature[sigSize-len(rBytes):sigSize], rBytes)
	copy(signature[sigSize*2-len(sBytes):], sBytes)

	return signature, nil
}

// Verify verifies an ECDSA signature (ASN.1 DER format).
func (s *ES256Signer) Verify(payload, signature []byte, kp *KeyPair) bool {
	publicKey, ok := kp.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return false
	}

	hash := sha256.Sum256(payload)
	return ecdsa.VerifyASN1(publicKey, hash[:], signature)
}
