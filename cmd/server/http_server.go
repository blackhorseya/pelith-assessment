package server

import (
	v1 "github.com/blackhorseya/pelith-assessment/cmd/server/v1"
	"github.com/blackhorseya/pelith-assessment/cmd/server/wirex"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/gin-gonic/gin"
)

// NewInitRoutesFn is a function to init routes
func NewInitRoutesFn(injector *wirex.Injector) httpx.InitRoutes {
	return func(router *gin.Engine) {
		api := router.Group("/api")
		{
			v1.Handler(api, injector)
		}
	}
}
