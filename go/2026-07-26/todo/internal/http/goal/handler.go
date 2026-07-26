package goal

import (
	"net/http"

	"todo/internal/todo/goal"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *goal.Service
}

func NewHandler(service *goal.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// createBody is the JSON shape of POST /goals
type createBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// updateBody is the JSON shape of PUT /goals/:id
type updateBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Create handles POST /goals
func (h *Handler) Create(c *gin.Context) {
	var body createBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	req := &goal.CreateRequest{
		Title:       body.Title,
		Description: body.Description,
	}

	result, errType, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != goal.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusCreated, result.Goal)
}

// List handles GET /goals
func (h *Handler) List(c *gin.Context) {
	result, errType, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != goal.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"goals": result.Goals})
}

// Get handles GET /goals/:id
func (h *Handler) Get(c *gin.Context) {
	result, errType, err := h.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != goal.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusOK, result.Goal)
}

// Update handles PUT /goals/:id
func (h *Handler) Update(c *gin.Context) {
	var body updateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	req := &goal.UpdateRequest{
		ID:          c.Param("id"),
		Title:       body.Title,
		Description: body.Description,
	}

	result, errType, err := h.service.Update(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != goal.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusOK, result.Goal)
}

// Delete handles DELETE /goals/:id
func (h *Handler) Delete(c *gin.Context) {
	errType, err := h.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != goal.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.Status(http.StatusNoContent)
}

// statusFor maps domain error types to HTTP status codes
func statusFor(e goal.ErrorType) int {
	switch e {
	case goal.ErrorInvalidRequest:
		return http.StatusBadRequest
	case goal.ErrorNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
