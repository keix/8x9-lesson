package source

import (
	"context"
	"sync"

	"asteroid/internal/userinfo"
)

// MemoryProvider is an in-memory implementation of userinfo.Provider
type MemoryProvider struct {
	mu    sync.RWMutex
	users map[string]map[string]any
}

// NewMemoryProvider creates a new in-memory userinfo provider
func NewMemoryProvider() *MemoryProvider {
	return &MemoryProvider{
		users: make(map[string]map[string]any),
	}
}

// Fetch retrieves user information by subject identifier
func (p *MemoryProvider) Fetch(ctx context.Context, sub string) (map[string]any, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	user, ok := p.users[sub]
	if !ok {
		return nil, userinfo.ErrUserNotFound
	}
	return user, nil
}

// RegisterUser adds a user to the provider (for testing/setup)
func (p *MemoryProvider) RegisterUser(sub string, claims map[string]any) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.users[sub] = claims
}
