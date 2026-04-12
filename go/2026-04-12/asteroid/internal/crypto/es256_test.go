package crypto

import (
	"testing"
)

func TestES256Generator_Generate(t *testing.T) {
	gen := NewES256Generator()

	kp, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if kp.Algorithm != "ES256" {
		t.Errorf("Algorithm = %v, want ES256", kp.Algorithm)
	}

	if kp.KeyID == "" {
		t.Error("KeyID should not be empty")
	}

	if kp.PrivateKey == nil {
		t.Error("PrivateKey should not be nil")
	}

	if kp.PublicKey == nil {
		t.Error("PublicKey should not be nil")
	}
}

func TestES256Generator_Alg(t *testing.T) {
	gen := NewES256Generator()
	if got := gen.Alg(); got != "ES256" {
		t.Errorf("Alg() = %v, want ES256", got)
	}
}

func TestES256Signer_Sign(t *testing.T) {
	gen := NewES256Generator()
	signer := NewES256Signer()

	kp, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	payload := []byte("test payload")
	signature, err := signer.Sign(payload, kp)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	if len(signature) == 0 {
		t.Error("signature should not be empty")
	}
}

func TestES256Signer_SignAndVerify(t *testing.T) {
	gen := NewES256Generator()
	signer := NewES256Signer()

	kp, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	payload := []byte("test payload for signing and verification")
	signature, err := signer.Sign(payload, kp)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	if !signer.Verify(payload, signature, kp) {
		t.Error("Verify() returned false, want true")
	}
}

func TestES256Signer_VerifyInvalidSignature(t *testing.T) {
	gen := NewES256Generator()
	signer := NewES256Signer()

	kp, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	payload := []byte("test payload")
	invalidSignature := []byte("invalid signature")

	if signer.Verify(payload, invalidSignature, kp) {
		t.Error("Verify() returned true for invalid signature")
	}
}

func TestES256Signer_SignJWT(t *testing.T) {
	gen := NewES256Generator()
	signer := NewES256Signer()

	kp, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	payload := []byte("eyJhbGciOiJFUzI1NiJ9.eyJzdWIiOiJ0ZXN0In0")
	signature, err := signer.SignJWT(payload, kp)
	if err != nil {
		t.Fatalf("SignJWT() error = %v", err)
	}

	// ES256 signature should be 64 bytes (32 bytes for R, 32 bytes for S)
	if len(signature) != 64 {
		t.Errorf("SignJWT() signature length = %d, want 64", len(signature))
	}
}

func TestGenerateKIDFromECDSAPublicKey(t *testing.T) {
	gen := NewES256Generator()

	kp, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	kid := kp.KeyID
	if kid == "" {
		t.Error("KID should not be empty")
	}

	// KID should be base64url encoded
	if len(kid) == 0 {
		t.Error("KID should have content")
	}
}
