package v1

import (
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/infra/transports/http/v1/users"
	"github.com/gin-gonic/gin"
)

// Handler is the handler for v1 api
func Handler(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		users.Handler(v1)
	}
}
