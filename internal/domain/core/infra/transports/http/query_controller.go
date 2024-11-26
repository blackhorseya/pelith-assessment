package http

import (
	"net/http"

	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/gin-gonic/gin"
)

// QueryController is the controller for query
type QueryController struct {
	rewardStore *query.RewardQueryStore
	userStore   *query.UserQueryStore
}

// NewQueryController is the constructor for QueryController
func NewQueryController(rewardStore *query.RewardQueryStore, userStore *query.UserQueryStore) *QueryController {
	return &QueryController{
		rewardStore: rewardStore,
		userStore:   userStore,
	}
}

// GetTasksStatusQuery is the query to get tasks status
type GetTasksStatusQuery struct {
	CampaignID string `form:"campaignID" binding:"required"`
	Page       int64  `form:"page" default:"1" minimum:"1"`
	Size       int64  `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// GetTasksStatus is the handler to get tasks status
// @Summary Get tasks status
// @Description Get tasks status by address
// @Tags users
// @Accept json
// @Produce json
// @Param address path string true "User address"
// @Param params query GetTasksStatusQuery false "query string"
// @Router /v1/users/{address}/tasks/status [get]
func (ctrl *QueryController) GetTasksStatus(c *gin.Context) {
	ret, err := ctrl.userStore.GetTasksStatus(c.Request.Context(), c.Param("address"), c.Query("campaignID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ret)
}

// GetPointsHistoryQuery is the query to get points history
type GetPointsHistoryQuery struct {
	Page int64 `form:"page" default:"1" minimum:"1"`
	Size int64 `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// GetPointsHistory is the handler to get points history
// @Summary Get points history
// @Description Get points history by address
// @Tags users
// @Accept json
// @Produce json
// @Param address path string true "User address"
// @Param params query GetPointsHistoryQuery false "query string"
// @Router /v1/users/{address}/points/history [get]
func (ctrl *QueryController) GetPointsHistory(c *gin.Context) {
	rewards, err := ctrl.rewardStore.GetRewardHistoryByWalletAddress(c.Request.Context(), c.Param("address"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rewards": rewards})
}
