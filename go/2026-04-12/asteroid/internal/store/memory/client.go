package memory

import (
	"context"
	"sync"

	"asteroid/internal/store/entity"
)

// ClientStore is an in-memory implementation of store.ClientStore
type ClientStore struct {
	mu      sync.RWMutex
	clients map[string]*entity.Client
}

// NewClientStore creates a new in-memory client store
func NewClientStore() *ClientStore {
	return &ClientStore{
		clients: make(map[string]*entity.Client),
	}
}

// GetClient retrieves a client by ID
func (s *ClientStore) GetClient(ctx context.Context, clientID string) (*entity.Client, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	client, ok := s.clients[clientID]
	if !ok {
		return nil, entity.ErrClientNotFound
	}
	return client, nil
}

// RegisterClient adds a client to the store (for testing/setup)
func (s *ClientStore) RegisterClient(client *entity.Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[client.ID] = client
}
