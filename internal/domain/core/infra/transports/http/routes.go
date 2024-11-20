package http

import (
	v1 "github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/transports/http/v1"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/gin-gonic/gin"
)

// NewInitUserRoutesFn is the function to init user routes
func NewInitUserRoutesFn() httpx.InitRoutes {
	return func(router *gin.Engine) {
		api := router.Group("/api")
		{
			v1.Handler(api)
		}
	}
}
