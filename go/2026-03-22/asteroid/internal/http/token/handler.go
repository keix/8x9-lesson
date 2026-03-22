package token

import (
	"net/http"

	"asteroid/internal/oidc/token"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *token.Service
}

func NewHandler(service *token.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ExchangeToken(c *gin.Context) {
	result, err := h.service.ExchangeToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
