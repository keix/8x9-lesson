package authorize

import (
	"net/http"

	"asteroid/internal/oidc/authorize"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *authorize.Service
}

func NewHandler(service *authorize.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Authorize(c *gin.Context) {
	// Build request from query parameters
	req := &authorize.AuthorizeRequest{
		ClientID:            c.Query("client_id"),
		RedirectURI:         c.Query("redirect_uri"),
		ResponseType:        c.Query("response_type"),
		Scope:               c.Query("scope"),
		State:               c.Query("state"),
		Nonce:               c.Query("nonce"),
		CodeChallenge:       c.Query("code_challenge"),
		CodeChallengeMethod: c.Query("code_challenge_method"),
		UserID:              c.GetHeader("X-Authenticated-User"),
	}

	result, errType, err := h.service.Authorize(c.Request.Context(), req)
	if err != nil {
		// Internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}

	if errType != authorize.ErrorNone {
		// OAuth2 error - return as JSON for now
		// In production, some errors should redirect with error params
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errType.String(),
		})
		return
	}

	// Success - redirect to client
	c.Redirect(http.StatusFound, result.RedirectURL)
}
