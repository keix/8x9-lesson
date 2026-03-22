package wellknown

import (
	"net/http"

	"asteroid/internal/oidc/wellknown"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *wellknown.Service
}

func NewHandler(service *wellknown.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetConfiguration(c *gin.Context) {
	config := h.service.GetConfiguration()
	c.JSON(http.StatusOK, config)
}
