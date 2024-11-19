package server

import (
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/gin-gonic/gin"
)

// NewInitRoutesFn is a function to init routes
func NewInitRoutesFn() httpx.InitRoutes {
	return func(router *gin.Engine) {
		// TODO: 2024/11/20|sean|binding routes
	}
}
