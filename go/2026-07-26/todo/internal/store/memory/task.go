package memory

import (
	"context"
	"sort"
	"sync"

	"todo/internal/store/entity"
)

// TaskStore is an in-memory implementation of store.TaskStore
type TaskStore struct {
	mu    sync.RWMutex
	tasks map[string]*entity.Task
}

// NewTaskStore creates a new in-memory task store
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[string]*entity.Task),
	}
}

// SaveTask inserts or updates a task
func (s *TaskStore) SaveTask(ctx context.Context, task *entity.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.ID] = task
	return nil
}

// GetTask retrieves a task by ID
func (s *TaskStore) GetTask(ctx context.Context, id string) (*entity.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.tasks[id]
	if !ok {
		return nil, entity.ErrTaskNotFound
	}
	return task, nil
}

// ListTasksByGoal returns all tasks of a goal ordered by creation time
func (s *TaskStore) ListTasksByGoal(ctx context.Context, goalID string) ([]*entity.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*entity.Task, 0)
	for _, t := range s.tasks {
		if t.GoalID == goalID {
			tasks = append(tasks, t)
		}
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
	})
	return tasks, nil
}

// DeleteTask removes a task by ID
func (s *TaskStore) DeleteTask(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return entity.ErrTaskNotFound
	}
	delete(s.tasks, id)
	return nil
}

// DeleteTasksByGoal removes all tasks that belong to a goal (cascade delete)
func (s *TaskStore) DeleteTasksByGoal(ctx context.Context, goalID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, t := range s.tasks {
		if t.GoalID == goalID {
			delete(s.tasks, id)
		}
	}
	return nil
}
