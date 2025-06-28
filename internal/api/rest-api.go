package api

import (
	"github.com/gin-gonic/gin"
	auth "github.com/lykeion-org/aegis/internal/auth"
)

type api struct {
	router *gin.Engine
	port string
	aegisHandler auth.AuthHandler
}

func NewApi(authHandler auth.AuthHandler) *api {
	r := gin.Default()

	return &api {
		router: r,
		port: ":40800",
		aegisHandler: authHandler,		
	}
}

func (a *api) InitializeApi(port string) {	
	a.port = port
	dev := a.router.Group("/")	
	restHandler := &RestHandler{
		handler: a.aegisHandler,
	}
	dev.POST("/create", restHandler.GenerateTokens)
	dev.POST("/validate", restHandler.Validate)
	dev.POST("/refresh",restHandler.RefreshTokens)

	a.router.Run(port)
}