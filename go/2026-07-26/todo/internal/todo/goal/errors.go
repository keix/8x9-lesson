package goal

// ErrorType represents business-level error types for goal operations
// These are domain errors, not HTTP errors
type ErrorType int

const (
	ErrorNone ErrorType = iota
	ErrorInvalidRequest
	ErrorNotFound
	ErrorServerError
)

// String returns the error code string
func (e ErrorType) String() string {
	switch e {
	case ErrorInvalidRequest:
		return "invalid_request"
	case ErrorNotFound:
		return "goal_not_found"
	default:
		return "server_error"
	}
}
