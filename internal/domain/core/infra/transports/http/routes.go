package http

import (
	"errors"
	"io"
	"net/http"

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
)

type routesImpl struct {
	campaignClient core.CampaignServiceClient
}

// NewInitUserRoutesFn is the function to init user routes
func NewInitUserRoutesFn(
	queryCtrl *QueryController,
	commandCtrl *CommandController,
	campaignClient core.CampaignServiceClient,
) httpx.InitRoutes {
	instance := &routesImpl{
		campaignClient: campaignClient,
	}

	return func(router *gin.Engine) {
		// frontend
		web.SetHTMLTemplate(router)
		router.GET("/", instance.index)
		router.GET("/tasks/status", instance.getTasksStatus)
		router.GET("/points/history", instance.getPointsHistory)
		router.GET("/campaigns/new", instance.newCampaigns)
		router.GET("/campaigns/:id", instance.getCampaignByID)

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

				campaigns := v1.Group("/campaigns")
				{
					campaigns.POST("", commandCtrl.CreateCampaign)
					campaigns.POST("/:id/start", commandCtrl.StartCampaign)
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

func (i *routesImpl) newCampaigns(c *gin.Context) {
	c.HTML(http.StatusOK, "includes/new_campaign", nil)
}

func (i *routesImpl) getTasksStatus(c *gin.Context) {
	c.HTML(http.StatusOK, "includes/tasks_status", nil)
}

func (i *routesImpl) getPointsHistory(c *gin.Context) {
	c.HTML(http.StatusOK, "includes/points_history", nil)
}

func (i *routesImpl) getCampaignByID(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	campaign, err := i.campaignClient.GetCampaign(ctx, &core.GetCampaignRequest{Id: c.Param("id")})
	if err != nil {
		ctx.Error("failed to get campaign", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get campaign"})
		return
	}

	ctx.Debug("get campaign", zap.Any("campaign", &campaign))
	c.HTML(http.StatusOK, "includes/campaign_detail", gin.H{
		"title":    campaign.Campaign.Name,
		"Campaign": campaign.Campaign,
		"Tasks":    campaign.Tasks,
	})
}
