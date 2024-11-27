package http

import (
	"context"
	"errors"
	"fmt"
	"io"
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
	Name      string  `form:"name" binding:"required"`
	StartAt   string  `form:"startAt" binding:"required"`
	PoolID    string  `form:"poolID" binding:"required"`
	Mode      int32   `form:"mode" default:"2"`
	MinAmount float64 `form:"minAmount" default:"1000"`
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
		Mode:       model.CampaignMode(req.Mode),
		TargetPool: req.PoolID,
		MinAmount:  req.MinAmount,
	})
	if err != nil {
		ctx.Error("failed to create campaign", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create campaign"})
		return
	}

	ctx.Debug("create campaign", zap.Any("campaign", campaign))
	c.Redirect(http.StatusSeeOther, "/")
}

func (ctrl *CommandController) RunBacktest(c *gin.Context) {
	// 获取请求的上下文
	ctx := contextx.WithContext(c.Request.Context())

	campaignID := c.Param("id")
	if campaignID == "" {
		ctx.Error("campaign id is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign id is empty"})
		return
	}

	// 启动后台任务
	go func(c context.Context, campaignID string) {
		ctrl.runBacktestTask(c, campaignID)
	}(context.Background(), campaignID)

	// 返回 HTTP 响应，表示任务已启动
	c.JSON(http.StatusAccepted, gin.H{"message": "backtest has been started"})
}

func (ctrl *CommandController) runBacktestTask(c context.Context, campaignID string) {
	ctx := contextx.WithContext(c)

	// 调用 gRPC 客户端获取流
	stream, err := ctrl.campaignClient.RunBacktestByCampaign(ctx, &core.GetCampaignRequest{Id: campaignID})
	if err != nil {
		ctx.Error("failed to start backtest stream", zap.Error(err))
		return
	}
	defer func() {
		_ = stream.CloseSend()
	}()

	ctx.Info("backtest stream started")

	// 处理流数据
	for {
		select {
		case <-ctx.Done():
			ctx.Warn("backtest stream canceled")
			return
		default:
			// 从 gRPC 流中接收消息
			message, err2 := stream.Recv()
			if errors.Is(err2, io.EOF) {
				ctx.Info("backtest stream completed")
				return
			}
			if err2 != nil {
				ctx.Error("error while receiving from backtest stream", zap.Error(err2))
				return
			}

			// 处理接收到的消息
			ctx.Info("received backtest message", zap.Any("message", &message))
		}
	}
}
