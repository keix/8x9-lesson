package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()

	s := &Server{
		router: r,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.GET("/hello", s.handleHello)
	s.router.GET("/user/:name", s.handleUser)
}

func (s *Server) handleHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func (s *Server) handleUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome, " + name + "!",
	})
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
