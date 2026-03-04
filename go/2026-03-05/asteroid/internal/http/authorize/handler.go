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
	result, err := h.service.Authorize()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
