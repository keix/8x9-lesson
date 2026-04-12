package store

import (
	"context"

	"asteroid/internal/store/entity"
)

// ClientStore manages OAuth 2.0 client data
type ClientStore interface {
	GetClient(ctx context.Context, clientID string) (*entity.Client, error)
}

// AuthCodeStore manages authorization codes
type AuthCodeStore interface {
	SaveAuthCode(ctx context.Context, code *entity.AuthCode) error
	GetAuthCode(ctx context.Context, code string) (*entity.AuthCode, error)
	DeleteAuthCode(ctx context.Context, code string) error
}

// NonceStore tracks nonces to prevent replay attacks
type NonceStore interface {
	MarkNonceSeen(ctx context.Context, nonce string, clientID string) error
}
