package memory

import (
	"context"
	"sync"

	"asteroid/internal/store/entity"
)

// NonceStore is an in-memory implementation of store.NonceStore
type NonceStore struct {
	mu     sync.RWMutex
	nonces map[string]bool // key: "clientID:nonce"
}

// NewNonceStore creates a new in-memory nonce store
func NewNonceStore() *NonceStore {
	return &NonceStore{
		nonces: make(map[string]bool),
	}
}

// MarkNonceSeen marks a nonce as seen for a given client
// Returns ErrNonceAlreadySeen if the nonce was already used
func (s *NonceStore) MarkNonceSeen(ctx context.Context, nonce string, clientID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := clientID + ":" + nonce
	if s.nonces[key] {
		return entity.ErrNonceAlreadySeen
	}
	s.nonces[key] = true
	return nil
}
