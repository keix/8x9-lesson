package authorize

// ErrorType represents OAuth2/OIDC authorization error types
// These are protocol-level errors, not HTTP errors
type ErrorType int

const (
	ErrorNone ErrorType = iota
	ErrorInvalidRequest
	ErrorInvalidClient
	ErrorInvalidRedirectURI
	ErrorUnsupportedResponse
	ErrorInvalidScope
	ErrorAccessDenied
	ErrorServerError
)

// String returns the OAuth2 error code string
func (e ErrorType) String() string {
	switch e {
	case ErrorInvalidRequest:
		return "invalid_request"
	case ErrorInvalidClient:
		return "invalid_client"
	case ErrorInvalidRedirectURI:
		return "invalid_redirect_uri"
	case ErrorUnsupportedResponse:
		return "unsupported_response_type"
	case ErrorInvalidScope:
		return "invalid_scope"
	case ErrorAccessDenied:
		return "access_denied"
	default:
		return "server_error"
	}
}
