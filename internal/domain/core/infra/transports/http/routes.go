package http

import (
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/gin-gonic/gin"
)

// NewInitUserRoutesFn is the function to init user routes
func NewInitUserRoutesFn(queryCtrl *QueryController) httpx.InitRoutes {
	return func(router *gin.Engine) {
		api := router.Group("/api")
		{
			v1 := api.Group("/v1")
			{
				users := v1.Group("/users")
				{
					user := users.Group("/:address")
					{
						user.GET("/tasks/status", queryCtrl.GetTasksStatus)
						user.GET("/points/history", queryCtrl.GetPointsHistory)
					}
				}
			}
		}
	}
}
