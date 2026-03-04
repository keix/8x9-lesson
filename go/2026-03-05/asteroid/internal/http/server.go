package http

import (
	"net/http"

	httpauthorize "asteroid/internal/http/authorize"
	httpjwks "asteroid/internal/http/jwks"
	httptoken "asteroid/internal/http/token"
	httpwellknown "asteroid/internal/http/wellknown"
	"asteroid/internal/oidc/authorize"
	"asteroid/internal/oidc/jwks"
	"asteroid/internal/oidc/token"
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

	authorizeService := authorize.NewService()
	authorizeHandler := httpauthorize.NewHandler(authorizeService)
	s.router.GET("/authorize", authorizeHandler.Authorize)

	tokenService := token.NewService()
	tokenHandler := httptoken.NewHandler(tokenService)
	s.router.POST("/token", tokenHandler.ExchangeToken)
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
