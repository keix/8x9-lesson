package crypto

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

// GenerateKIDFromECDSAPublicKey generates a Key ID from an ECDSA public key.
// The KID is derived from the SHA-256 hash of the public key coordinates.
func GenerateKIDFromECDSAPublicKey(pub *ecdsa.PublicKey) string {
	// Concatenate X and Y coordinates
	data := append(pub.X.Bytes(), pub.Y.Bytes()...)

	// Hash with SHA-256
	hash := sha256.Sum256(data)

	// Return base64url-encoded first 16 bytes
	return base64.RawURLEncoding.EncodeToString(hash[:16])
}

// Base64URLEncode encodes bytes to base64url format (used for JWK).
func Base64URLEncode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// BigIntToBase64URL converts a big.Int to base64url format.
func BigIntToBase64URL(n *big.Int) string {
	return Base64URLEncode(n.Bytes())
}

// ECDSAPublicKeyToJWK converts an ECDSA public key to JWK format.
func ECDSAPublicKeyToJWK(pub *ecdsa.PublicKey, kid, alg string) map[string]string {
	// P-256 curve has 32-byte coordinates
	const coordSize = 32

	// Pad X and Y to 32 bytes
	xBytes := padToSize(pub.X.Bytes(), coordSize)
	yBytes := padToSize(pub.Y.Bytes(), coordSize)

	return map[string]string{
		"kty": "EC",
		"crv": "P-256",
		"use": "sig",
		"alg": alg,
		"kid": kid,
		"x":   Base64URLEncode(xBytes),
		"y":   Base64URLEncode(yBytes),
	}
}

// padToSize pads data to the specified size with leading zeros.
func padToSize(data []byte, size int) []byte {
	if len(data) >= size {
		return data
	}
	padded := make([]byte, size)
	copy(padded[size-len(data):], data)
	return padded
}

// FormatKeyID formats a key ID with algorithm prefix.
func FormatKeyID(alg, kid string) string {
	return fmt.Sprintf("%s-%s", alg, kid)
}
