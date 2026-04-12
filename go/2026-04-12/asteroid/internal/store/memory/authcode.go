package memory

import (
	"context"
	"sync"

	"asteroid/internal/store/entity"
)

// AuthCodeStore is an in-memory implementation of store.AuthCodeStore
type AuthCodeStore struct {
	mu    sync.RWMutex
	codes map[string]*entity.AuthCode
}

// NewAuthCodeStore creates a new in-memory auth code store
func NewAuthCodeStore() *AuthCodeStore {
	return &AuthCodeStore{
		codes: make(map[string]*entity.AuthCode),
	}
}

// SaveAuthCode stores an authorization code
func (s *AuthCodeStore) SaveAuthCode(ctx context.Context, code *entity.AuthCode) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.codes[code.Code] = code
	return nil
}

// GetAuthCode retrieves an authorization code
func (s *AuthCodeStore) GetAuthCode(ctx context.Context, code string) (*entity.AuthCode, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	authCode, ok := s.codes[code]
	if !ok {
		return nil, entity.ErrAuthCodeNotFound
	}
	return authCode, nil
}

// DeleteAuthCode removes an authorization code (single-use)
func (s *AuthCodeStore) DeleteAuthCode(ctx context.Context, code string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.codes, code)
	return nil
}
