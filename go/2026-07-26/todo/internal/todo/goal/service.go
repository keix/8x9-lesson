package goal

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"todo/internal/store"
	"todo/internal/store/entity"
)

// Service handles goal business logic
type Service struct {
	GoalStore store.GoalStore
	TaskStore store.TaskStore
}

// NewService creates a new goal service
func NewService(goalStore store.GoalStore, taskStore store.TaskStore) *Service {
	return &Service{
		GoalStore: goalStore,
		TaskStore: taskStore,
	}
}

// CreateRequest represents the data needed to create a goal
type CreateRequest struct {
	Title       string
	Description string
}

// UpdateRequest represents the data needed to update a goal
type UpdateRequest struct {
	ID          string
	Title       string
	Description string
}

// Create registers a new goal
func (s *Service) Create(ctx context.Context, req *CreateRequest) (*Result, ErrorType, error) {
	// TODO: Step 1 - Validate request
	// - Title must not be empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Build entity.Goal
	// - Use uuid.NewString() for ID generation
	// - Set CreatedAt / UpdatedAt to time.Now()

	// TODO: Step 3 - Save with s.GoalStore.SaveGoal()
	// - Return ErrorServerError with the error if save fails

	// TODO: Step 4 - Return &Result{Goal: goal}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// Get retrieves a goal by ID
func (s *Service) Get(ctx context.Context, id string) (*Result, ErrorType, error) {
	// TODO: Step 1 - Validate that id is not empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Fetch with s.GoalStore.GetGoal()
	// - Return ErrorNotFound if errors.Is(err, entity.ErrGoalNotFound)
	// - Return ErrorServerError for any other error

	// TODO: Step 3 - Return &Result{Goal: goal}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// List returns all goals
func (s *Service) List(ctx context.Context) (*ListResult, ErrorType, error) {
	// TODO: Step 1 - Fetch with s.GoalStore.ListGoals()
	// - Return ErrorServerError if it fails

	// TODO: Step 2 - Return &ListResult{Goals: goals}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// Update modifies an existing goal
func (s *Service) Update(ctx context.Context, req *UpdateRequest) (*Result, ErrorType, error) {
	// TODO: Step 1 - Validate request
	// - ID and Title must not be empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Fetch the existing goal with s.GoalStore.GetGoal()
	// - Return ErrorNotFound if errors.Is(err, entity.ErrGoalNotFound)

	// TODO: Step 3 - Apply changes
	// - Overwrite Title / Description
	// - Set UpdatedAt to time.Now() (keep CreatedAt as is)

	// TODO: Step 4 - Save with s.GoalStore.SaveGoal()

	// TODO: Step 5 - Return &Result{Goal: goal}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// Delete removes a goal and its tasks
func (s *Service) Delete(ctx context.Context, id string) (ErrorType, error) {
	// TODO: Step 1 - Validate that id is not empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Delete with s.GoalStore.DeleteGoal()
	// - Return ErrorNotFound if errors.Is(err, entity.ErrGoalNotFound)

	// TODO: Step 3 - Cascade delete tasks with s.TaskStore.DeleteTasksByGoal()
	// - A goal's tasks must not survive their goal

	// TODO: Step 4 - Return ErrorNone, nil

	// Stub implementation - returns error for now
	return ErrorServerError, nil
}

// Suppress unused import warnings for stub
var (
	_ = uuid.NewString
	_ = time.Now
	_ = errors.Is
	_ = entity.ErrGoalNotFound
)
