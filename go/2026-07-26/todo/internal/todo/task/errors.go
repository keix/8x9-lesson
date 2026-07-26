package task

// ErrorType represents business-level error types for task operations
// These are domain errors, not HTTP errors
type ErrorType int

const (
	ErrorNone ErrorType = iota
	ErrorInvalidRequest
	ErrorNotFound
	ErrorGoalNotFound
	ErrorServerError
)

// String returns the error code string
func (e ErrorType) String() string {
	switch e {
	case ErrorInvalidRequest:
		return "invalid_request"
	case ErrorNotFound:
		return "task_not_found"
	case ErrorGoalNotFound:
		return "goal_not_found"
	default:
		return "server_error"
	}
}
