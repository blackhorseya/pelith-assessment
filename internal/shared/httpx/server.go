package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoutes is a function to initialize routes.
type InitRoutes func(router *gin.Engine)

// GinServer is an HTTP server.
type GinServer struct {
	httpserver *http.Server

	// Router is the gin engine.
	Router *gin.Engine
}
