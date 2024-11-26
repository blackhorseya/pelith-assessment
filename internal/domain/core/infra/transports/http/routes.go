package http

import (
	docs "github.com/blackhorseya/pelith-assessment/docs/api"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewInitUserRoutesFn is the function to init user routes
func NewInitUserRoutesFn(queryCtrl *QueryController) httpx.InitRoutes {
	return func(router *gin.Engine) {
		docs.SwaggerInfo.BasePath = "/api"
		api := router.Group("/api")
		{
			api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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
