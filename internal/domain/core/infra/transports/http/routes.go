package http

import (
	"net/http"

	docs "github.com/blackhorseya/pelith-assessment/docs/api"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/blackhorseya/pelith-assessment/web"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewInitUserRoutesFn is the function to init user routes
func NewInitUserRoutesFn(queryCtrl *QueryController) httpx.InitRoutes {
	return func(router *gin.Engine) {
		// frontend
		web.SetHTMLTemplate(router)
		router.GET("/", index)
		router.GET("/simulation", simulation)
		router.GET("/tasks/config", tasksConfig)
		router.POST("/tasks/config", saveTaskConfig)

		// restful api
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

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.templ", gin.H{
		"title": "Home Page",
	})
}

func simulation(c *gin.Context) {
}

// TaskConfig represents the structure of a task configuration
type TaskConfig struct {
	TaskName    string `form:"taskName" binding:"required"`
	Threshold   int    `form:"threshold" binding:"required"`
	TotalPoints int    `form:"points" binding:"required"`
}

// In-memory storage for task configurations
var taskConfigs []TaskConfig

// Handle GET request to render task configuration page
func tasksConfig(c *gin.Context) {
	c.HTML(http.StatusOK, "config.templ", gin.H{
		"title": "Task Configuration",
	})
}

// Handle POST request to save a new task configuration
func saveTaskConfig(c *gin.Context) {
	var newConfig TaskConfig
	if err := c.ShouldBind(&newConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the new configuration
	taskConfigs = append(taskConfigs, newConfig)
	c.Redirect(http.StatusSeeOther, "/tasks/config")
}
