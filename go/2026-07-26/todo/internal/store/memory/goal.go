package memory

import (
	"context"
	"sort"
	"sync"

	"todo/internal/store/entity"
)

// GoalStore is an in-memory implementation of store.GoalStore
type GoalStore struct {
	mu    sync.RWMutex
	goals map[string]*entity.Goal
}

// NewGoalStore creates a new in-memory goal store
func NewGoalStore() *GoalStore {
	return &GoalStore{
		goals: make(map[string]*entity.Goal),
	}
}

// SaveGoal inserts or updates a goal
func (s *GoalStore) SaveGoal(ctx context.Context, goal *entity.Goal) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.goals[goal.ID] = goal
	return nil
}

// GetGoal retrieves a goal by ID
func (s *GoalStore) GetGoal(ctx context.Context, id string) (*entity.Goal, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	goal, ok := s.goals[id]
	if !ok {
		return nil, entity.ErrGoalNotFound
	}
	return goal, nil
}

// ListGoals returns all goals ordered by creation time
func (s *GoalStore) ListGoals(ctx context.Context) ([]*entity.Goal, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	goals := make([]*entity.Goal, 0, len(s.goals))
	for _, g := range s.goals {
		goals = append(goals, g)
	}
	sort.Slice(goals, func(i, j int) bool {
		return goals[i].CreatedAt.Before(goals[j].CreatedAt)
	})
	return goals, nil
}

// DeleteGoal removes a goal by ID
func (s *GoalStore) DeleteGoal(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.goals[id]; !ok {
		return entity.ErrGoalNotFound
	}
	delete(s.goals, id)
	return nil
}
