package authorize

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"

	"asteroid/internal/store"
	"asteroid/internal/store/entity"
	"asteroid/internal/userinfo"
)

// Service handles authorization business logic
type Service struct {
	ClientStore      store.ClientStore
	UserinfoProvider userinfo.Provider
	AuthCodeStore    store.AuthCodeStore
	NonceStore       store.NonceStore
}

// NewService creates a new authorization service
func NewService(
	clientStore store.ClientStore,
	userinfoProvider userinfo.Provider,
	authCodeStore store.AuthCodeStore,
	nonceStore store.NonceStore,
) *Service {
	return &Service{
		ClientStore:      clientStore,
		UserinfoProvider: userinfoProvider,
		AuthCodeStore:    authCodeStore,
		NonceStore:       nonceStore,
	}
}

// AuthorizeRequest represents the data needed for authorization
type AuthorizeRequest struct {
	ClientID            string
	RedirectURI         string
	ResponseType        string
	Scope               string
	State               string
	Nonce               string
	CodeChallenge       string
	CodeChallengeMethod string
	UserID              string // From X-Authenticated-User header
}

// Authorize processes authorization request (pure business logic)
func (s *Service) Authorize(ctx context.Context, req *AuthorizeRequest) (*Result, ErrorType, error) {
	// TODO: Step 1 - Validate required parameters
	// - client_id, redirect_uri, response_type must not be empty
	// - Return ErrorInvalidRequest if any is missing

	// TODO: Step 2 - Validate state parameter (CSRF protection)
	// - state is mandatory for security
	// - Return ErrorInvalidRequest if missing

	// TODO: Step 3 - Nonce replay protection (optional but recommended)
	// - If nonce is provided, check if it was already used
	// - Use s.NonceStore.MarkNonceSeen()
	// - Return ErrorInvalidRequest if nonce was already seen

	// TODO: Step 4 - Validate response_type
	// - Only "code" is supported (Authorization Code Flow)
	// - Return ErrorUnsupportedResponse otherwise

	// TODO: Step 5 - Validate scope
	// - Only "openid" is supported for now
	// - Return ErrorInvalidScope otherwise

	// TODO: Step 6 - Get and validate client
	// - Use s.ClientStore.GetClient()
	// - Return ErrorInvalidClient if not found

	// TODO: Step 7 - PKCE validation
	// - Public clients MUST use PKCE
	// - Only S256 method is supported
	// - Use s.validatePKCEForClient()

	// TODO: Step 8 - Validate redirect_uri
	// - Must exactly match one of the registered URIs
	// - Use validateExactRedirectURI()
	// - Return ErrorInvalidRedirectURI if not matched

	// TODO: Step 9 - Validate authenticated user
	// - UserID (from X-Authenticated-User header) must not be empty
	// - Use s.UserinfoProvider.Fetch() to verify user exists
	// - Return ErrorAccessDenied if user not found

	// TODO: Step 10 - Generate authorization code
	// - Use uuid.NewString() for code generation
	// - Create entity.AuthCode with all required fields
	// - Set ExpiresAt to 5 minutes from now
	// - Save with s.AuthCodeStore.SaveAuthCode()

	// TODO: Step 11 - Build redirect URL
	// - Format: redirect_uri?code=XXX&state=YYY
	// - Return Result with RedirectURL

	// Stub implementation - returns error for now
	return nil, ErrorServerError, nil
}

// validateExactRedirectURI performs RFC 6749 compliant exact redirect URI validation
// SECURITY: Uses string comparison to prevent URL normalization attacks
func validateExactRedirectURI(registeredURIs []string, requestedURI string) bool {
	return slices.Contains(registeredURIs, requestedURI)
}

// validatePKCEForClient validates PKCE requirements based on client type
func (s *Service) validatePKCEForClient(client *entity.Client, req *AuthorizeRequest) error {
	// For public clients, PKCE is mandatory
	if client.IsPublicClient() {
		if req.CodeChallenge == "" {
			return errors.New("PKCE required for public clients")
		}
		if req.CodeChallengeMethod != "S256" {
			return errors.New("only S256 method supported for PKCE")
		}
		return nil
	}

	// For confidential clients, PKCE is optional
	if req.CodeChallenge != "" && req.CodeChallengeMethod != "S256" {
		return errors.New("only S256 method supported for PKCE")
	}

	return nil
}

// Suppress unused import warnings for stub
var (
	_ = uuid.NewString
	_ = time.Now
	_ = entity.ErrClientNotFound
	_ = userinfo.ErrUserNotFound
)
