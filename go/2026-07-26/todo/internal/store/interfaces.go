package store

import (
	"context"

	"todo/internal/store/entity"
)

// GoalStore manages goal data
type GoalStore interface {
	SaveGoal(ctx context.Context, goal *entity.Goal) error
	GetGoal(ctx context.Context, id string) (*entity.Goal, error)
	ListGoals(ctx context.Context) ([]*entity.Goal, error)
	DeleteGoal(ctx context.Context, id string) error
}

// TaskStore manages task data
type TaskStore interface {
	SaveTask(ctx context.Context, task *entity.Task) error
	GetTask(ctx context.Context, id string) (*entity.Task, error)
	ListTasksByGoal(ctx context.Context, goalID string) ([]*entity.Task, error)
	DeleteTask(ctx context.Context, id string) error
	DeleteTasksByGoal(ctx context.Context, goalID string) error
}
