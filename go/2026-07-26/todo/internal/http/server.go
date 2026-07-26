package http

import (
	"net/http"

	httpgoal "todo/internal/http/goal"
	httptask "todo/internal/http/task"
	"todo/internal/store/memory"
	"todo/internal/todo/goal"
	"todo/internal/todo/task"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine

	// Stores (exposed for testing/setup)
	GoalStore *memory.GoalStore
	TaskStore *memory.TaskStore
}

func NewServer() *Server {
	r := gin.Default()

	// Initialize stores
	goalStore := memory.NewGoalStore()
	taskStore := memory.NewTaskStore()

	s := &Server{
		router:    r,
		GoalStore: goalStore,
		TaskStore: taskStore,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.GET("/health", s.handleHealth)

	goalService := goal.NewService(s.GoalStore, s.TaskStore)
	goalHandler := httpgoal.NewHandler(goalService)
	s.router.POST("/goals", goalHandler.Create)
	s.router.GET("/goals", goalHandler.List)
	s.router.GET("/goals/:id", goalHandler.Get)
	s.router.PUT("/goals/:id", goalHandler.Update)
	s.router.DELETE("/goals/:id", goalHandler.Delete)

	taskService := task.NewService(s.TaskStore, s.GoalStore)
	taskHandler := httptask.NewHandler(taskService)
	s.router.POST("/goals/:id/tasks", taskHandler.Create)
	s.router.GET("/goals/:id/tasks", taskHandler.ListByGoal)
	s.router.GET("/tasks/:id", taskHandler.Get)
	s.router.PUT("/tasks/:id", taskHandler.Update)
	s.router.DELETE("/tasks/:id", taskHandler.Delete)
}

func (s *Server) handleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
