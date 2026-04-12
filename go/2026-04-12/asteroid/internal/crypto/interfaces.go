package crypto

import (
	"crypto"
	"time"
)

// KeyPair represents a cryptographic key pair with metadata.
type KeyPair struct {
	PrivateKey crypto.PrivateKey
	PublicKey  crypto.PublicKey
	Algorithm  string
	KeyID      string
	CreatedAt  time.Time
}

// Generator creates cryptographic key pairs for specific algorithms.
type Generator interface {
	// Generate creates a new key pair.
	Generate() (*KeyPair, error)
	// Alg returns the algorithm identifier (e.g., "ES256").
	Alg() string
}

// Signer signs payloads using key pairs.
type Signer interface {
	// Sign creates a signature for the given payload using the key pair.
	Sign(payload []byte, kp *KeyPair) ([]byte, error)
}
