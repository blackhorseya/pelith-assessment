package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	docs "github.com/blackhorseya/pelith-assessment/docs/api"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/shared/httpx"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/blackhorseya/pelith-assessment/proto/core"
	"github.com/blackhorseya/pelith-assessment/web"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type routesImpl struct {
	campaignClient core.CampaignServiceClient
}

// NewInitUserRoutesFn is the function to init user routes
func NewInitUserRoutesFn(queryCtrl *QueryController, campaignClient core.CampaignServiceClient) httpx.InitRoutes {
	instance := &routesImpl{
		campaignClient: campaignClient,
	}

	return func(router *gin.Engine) {
		// frontend
		web.SetHTMLTemplate(router)
		router.GET("/", instance.index)
		router.GET("/simulation", simulation)
		router.GET("/tasks/config", tasksConfig)
		router.POST("/campaigns", instance.createCampaign)

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

func (i *routesImpl) index(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	// Get all campaigns
	stream, err := i.campaignClient.ListCampaigns(ctx, &core.ListCampaignsRequest{})
	if err != nil {
		ctx.Error("failed to get campaigns", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get campaigns"})
		return
	}

	var tasks []*model.Task
	for {
		resp, err2 := stream.Recv()
		if err2 != nil {
			if errors.Is(err2, io.EOF) {
				break
			}
			ctx.Error("failed to get campaign", zap.Error(err2))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get campaign"})
			return
		}

		tasks = append(tasks, resp.Tasks...)
	}

	ctx.Debug("get all tasks", zap.Any("tasks", tasks))

	c.HTML(http.StatusOK, "includes/tasks", gin.H{
		"title": "Home Page",
		"tasks": tasks,
	})
}

// CampaignRequest represents the structure of a campaign request
type CampaignRequest struct {
	Name    string `form:"name" binding:"required"`
	StartAt string `form:"startAt" binding:"required"`
	PoolID  string `form:"poolID" binding:"required"`
}

func (i *routesImpl) createCampaign(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	var req CampaignRequest
	err := c.ShouldBind(&req)
	if err != nil {
		ctx.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 解析前端提交的时间格式
	startAt, err := time.Parse("2006-01-02T15:04", req.StartAt)
	if err != nil {
		ctx.Error("failed to parse start time", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to parse start time: %s", req.StartAt)})
		return
	}

	campaign, err := i.campaignClient.CreateCampaign(ctx, &core.CreateCampaignRequest{
		Name:       req.Name,
		StartTime:  timestamppb.New(startAt),
		Mode:       model.CampaignMode_CAMPAIGN_MODE_BACKTEST,
		TargetPool: req.PoolID,
		MinAmount:  1000,
	})
	if err != nil {
		ctx.Error("failed to create campaign", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create campaign"})
		return
	}

	ctx.Debug("create campaign", zap.Any("campaign", campaign))
	c.Redirect(http.StatusSeeOther, "/")
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
	c.HTML(http.StatusOK, "includes/config", gin.H{
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
