package userinfo

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidSub   = errors.New("invalid sub")
)

// Provider abstracts user information retrieval
// Used for token validation and userinfo endpoint
type Provider interface {
	Fetch(ctx context.Context, sub string) (map[string]any, error)
}
