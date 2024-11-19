package users

import (
	"github.com/blackhorseya/pelith-assessment/cmd/server/wirex"
	"github.com/gin-gonic/gin"
)

func Handler(g *gin.RouterGroup, injector *wirex.Injector) {
	users := g.Group("/users")
	{
		user := users.Group("/:address")
		{
			// TODO: 2024/11/20|sean|implement the handler
			user.GET("/tasks/status")
			user.GET("/points/history")
		}
	}
}
