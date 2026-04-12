package crypto

import "errors"

var (
	// ErrInvalidKeyType is returned when the key type does not match the expected type.
	ErrInvalidKeyType = errors.New("invalid key type")

	// ErrKeyGenerationFailed is returned when key generation fails.
	ErrKeyGenerationFailed = errors.New("key generation failed")

	// ErrSigningFailed is returned when signing fails.
	ErrSigningFailed = errors.New("signing failed")
)
