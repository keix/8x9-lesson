package task

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"todo/internal/store"
	"todo/internal/store/entity"
)

// Service handles task business logic
type Service struct {
	TaskStore store.TaskStore
	GoalStore store.GoalStore // needed to verify the parent goal exists
}

// NewService creates a new task service
func NewService(taskStore store.TaskStore, goalStore store.GoalStore) *Service {
	return &Service{
		TaskStore: taskStore,
		GoalStore: goalStore,
	}
}

// CreateRequest represents the data needed to create a task
type CreateRequest struct {
	GoalID string
	Title  string
}

// UpdateRequest represents the data needed to update a task
type UpdateRequest struct {
	ID    string
	Title string
	Done  bool
}

// Create registers a new task under a goal
func (s *Service) Create(ctx context.Context, req *CreateRequest) (*Result, ErrorType, error) {
	// TODO: Step 1 - Validate request
	// - GoalID and Title must not be empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Verify the parent goal exists
	// - Use s.GoalStore.GetGoal()
	// - Return ErrorGoalNotFound if errors.Is(err, entity.ErrGoalNotFound)
	// - A task must never point at a goal that does not exist

	// TODO: Step 3 - Build entity.Task
	// - Use uuid.NewString() for ID generation
	// - Done starts as false
	// - Set CreatedAt / UpdatedAt to time.Now()

	// TODO: Step 4 - Save with s.TaskStore.SaveTask()

	// TODO: Step 5 - Return &Result{Task: task}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// Get retrieves a task by ID
func (s *Service) Get(ctx context.Context, id string) (*Result, ErrorType, error) {
	// TODO: Step 1 - Validate that id is not empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Fetch with s.TaskStore.GetTask()
	// - Return ErrorNotFound if errors.Is(err, entity.ErrTaskNotFound)

	// TODO: Step 3 - Return &Result{Task: task}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// ListByGoal returns all tasks that belong to a goal
func (s *Service) ListByGoal(ctx context.Context, goalID string) (*ListResult, ErrorType, error) {
	// TODO: Step 1 - Validate that goalID is not empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Verify the goal exists with s.GoalStore.GetGoal()
	// - Return ErrorGoalNotFound if errors.Is(err, entity.ErrGoalNotFound)
	// - "empty list" and "no such goal" are different answers

	// TODO: Step 3 - Fetch with s.TaskStore.ListTasksByGoal()

	// TODO: Step 4 - Return &ListResult{Tasks: tasks}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// Update modifies an existing task (title and done flag)
func (s *Service) Update(ctx context.Context, req *UpdateRequest) (*Result, ErrorType, error) {
	// TODO: Step 1 - Validate request
	// - ID and Title must not be empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Fetch the existing task with s.TaskStore.GetTask()
	// - Return ErrorNotFound if errors.Is(err, entity.ErrTaskNotFound)

	// TODO: Step 3 - Apply changes
	// - Overwrite Title / Done
	// - Set UpdatedAt to time.Now() (GoalID and CreatedAt stay as is)

	// TODO: Step 4 - Save with s.TaskStore.SaveTask()

	// TODO: Step 5 - Return &Result{Task: task}, ErrorNone, nil

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// Delete removes a task
func (s *Service) Delete(ctx context.Context, id string) (ErrorType, error) {
	// TODO: Step 1 - Validate that id is not empty
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 2 - Delete with s.TaskStore.DeleteTask()
	// - Return ErrorNotFound if errors.Is(err, entity.ErrTaskNotFound)

	// TODO: Step 3 - Return ErrorNone, nil

	// Stub implementation - returns error for now
	return ErrorServerError, nil
}

// Suppress unused import warnings for stub
var (
	_ = uuid.NewString
	_ = time.Now
	_ = errors.Is
	_ = entity.ErrTaskNotFound
)
