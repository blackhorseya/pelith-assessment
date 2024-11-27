package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	"github.com/blackhorseya/pelith-assessment/proto/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CommandController is the command controller.
type CommandController struct {
	campaignClient core.CampaignServiceClient
}

// NewCommandController is used to create a new command controller.
func NewCommandController(campaignClient core.CampaignServiceClient) *CommandController {
	return &CommandController{
		campaignClient: campaignClient,
	}
}

// CampaignRequest represents the structure of a campaign request
type CampaignRequest struct {
	Name    string `form:"name" binding:"required"`
	StartAt string `form:"startAt" binding:"required"`
	PoolID  string `form:"poolID" binding:"required"`
}

func (ctrl *CommandController) CreateCampaign(c *gin.Context) {
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

	campaign, err := ctrl.campaignClient.CreateCampaign(ctx, &core.CreateCampaignRequest{
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
