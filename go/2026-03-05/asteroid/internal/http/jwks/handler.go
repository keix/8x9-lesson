package jwks

import (
	"net/http"

	"asteroid/internal/oidc/jwks"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *jwks.Service
}

func NewHandler(service *jwks.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetKeySet(c *gin.Context) {
	keySet := h.service.GetKeySet()
	c.JSON(http.StatusOK, keySet)
}
