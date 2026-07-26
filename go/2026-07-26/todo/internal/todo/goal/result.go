package goal

import "todo/internal/store/entity"

// Result carries a single goal across the service boundary
type Result struct {
	Goal *entity.Goal
}

// ListResult carries multiple goals across the service boundary
type ListResult struct {
	Goals []*entity.Goal
}
