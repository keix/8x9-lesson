package task

import "todo/internal/store/entity"

// Result carries a single task across the service boundary
type Result struct {
	Task *entity.Task
}

// ListResult carries multiple tasks across the service boundary
type ListResult struct {
	Tasks []*entity.Task
}
