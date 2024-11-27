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
		router.GET("/tasks/status", instance.getTasksStatus)
		router.GET("/points/history", instance.getPointsHistory)
		router.POST("/campaigns", instance.createCampaign)
		router.GET("/campaigns/new", instance.newCampaigns)

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

	var campaigns []*model.Campaign
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

		campaigns = append(campaigns, resp.Campaign)
	}

	ctx.Debug("get campaigns", zap.Any("campaigns", campaigns))

	c.HTML(http.StatusOK, "includes/campaigns", gin.H{
		"title":     "Home Page",
		"campaigns": campaigns,
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

func (i *routesImpl) newCampaigns(c *gin.Context) {
	c.HTML(http.StatusOK, "includes/new_campaign", nil)
}

func (i *routesImpl) getTasksStatus(c *gin.Context) {
	c.HTML(http.StatusOK, "includes/tasks_status", nil)
}

func (i *routesImpl) getPointsHistory(c *gin.Context) {
	c.HTML(http.StatusOK, "includes/points_history", nil)
}
