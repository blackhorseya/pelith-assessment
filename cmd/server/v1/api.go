package v1

import (
	"github.com/blackhorseya/pelith-assessment/cmd/server/v1/users"
	"github.com/blackhorseya/pelith-assessment/cmd/server/wirex"
	"github.com/gin-gonic/gin"
)

// Handler is the handler for v1 api
func Handler(g *gin.RouterGroup, injector *wirex.Injector) {
	v1 := g.Group("/v1")
	{
		users.Handler(v1, injector)
	}
}
