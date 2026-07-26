package task

import (
	"net/http"

	"todo/internal/todo/task"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *task.Service
}

func NewHandler(service *task.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// createBody is the JSON shape of POST /goals/:id/tasks
type createBody struct {
	Title string `json:"title"`
}

// updateBody is the JSON shape of PUT /tasks/:id
type updateBody struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Create handles POST /goals/:id/tasks
func (h *Handler) Create(c *gin.Context) {
	var body createBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	req := &task.CreateRequest{
		GoalID: c.Param("id"),
		Title:  body.Title,
	}

	result, errType, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != task.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusCreated, result.Task)
}

// ListByGoal handles GET /goals/:id/tasks
func (h *Handler) ListByGoal(c *gin.Context) {
	result, errType, err := h.service.ListByGoal(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != task.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": result.Tasks})
}

// Get handles GET /tasks/:id
func (h *Handler) Get(c *gin.Context) {
	result, errType, err := h.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != task.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusOK, result.Task)
}

// Update handles PUT /tasks/:id
func (h *Handler) Update(c *gin.Context) {
	var body updateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	req := &task.UpdateRequest{
		ID:    c.Param("id"),
		Title: body.Title,
		Done:  body.Done,
	}

	result, errType, err := h.service.Update(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != task.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.JSON(http.StatusOK, result.Task)
}

// Delete handles DELETE /tasks/:id
func (h *Handler) Delete(c *gin.Context) {
	errType, err := h.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server_error"})
		return
	}
	if errType != task.ErrorNone {
		c.JSON(statusFor(errType), gin.H{"error": errType.String()})
		return
	}

	c.Status(http.StatusNoContent)
}

// statusFor maps domain error types to HTTP status codes
func statusFor(e task.ErrorType) int {
	switch e {
	case task.ErrorInvalidRequest:
		return http.StatusBadRequest
	case task.ErrorNotFound, task.ErrorGoalNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
