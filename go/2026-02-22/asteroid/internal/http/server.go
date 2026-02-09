package http

import (
	"net/http"

	httpjwks "asteroid/internal/http/jwks"
	httpwellknown "asteroid/internal/http/wellknown"
	"asteroid/internal/oidc/jwks"
	"asteroid/internal/oidc/wellknown"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(issuer string) *Server {
	r := gin.Default()

	s := &Server{
		router: r,
	}

	s.setupRoutes(issuer)
	return s
}

func (s *Server) setupRoutes(issuer string) {
	s.router.GET("/hello", s.handleHello)
	s.router.GET("/user/:name", s.handleUser)

	wellknownService := wellknown.NewService(issuer)
	wellknownHandler := httpwellknown.NewHandler(wellknownService)
	s.router.GET("/.well-known/openid-configuration", wellknownHandler.GetConfiguration)

	jwksService := jwks.NewService()
	jwksHandler := httpjwks.NewHandler(jwksService)
	s.router.GET("/jwks.json", jwksHandler.GetKeySet)
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
