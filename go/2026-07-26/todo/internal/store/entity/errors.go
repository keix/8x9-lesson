package entity

import "errors"

var (
	ErrGoalNotFound = errors.New("goal not found")
	ErrTaskNotFound = errors.New("task not found")
)
