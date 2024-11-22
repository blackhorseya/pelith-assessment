package http

import (
	"net/http"

	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/gin-gonic/gin"
)

// QueryController is the controller for query
type QueryController struct {
	taskQuery *query.TaskQueryService
}

// NewQueryController is the constructor for QueryController
func NewQueryController(taskQuery *query.TaskQueryService) *QueryController {
	return &QueryController{
		taskQuery: taskQuery,
	}
}

// GetTasksStatusQuery is the query to get tasks status
type GetTasksStatusQuery struct {
	CampaignID string `form:"campaignID"`
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
// @Param query query GetTasksStatusQuery false "query string"
// @Router /api/v1/users/{address}/tasks/status [get]
func (ctrl *QueryController) GetTasksStatus(c *gin.Context) {
	tasks, err := ctrl.taskQuery.GetTaskStatus(c.Request.Context(), c.Param("address"), c.Query("campaignID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
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
// @Param query query GetPointsHistoryQuery false "query string"
// @Router /api/v1/users/{address}/points/history [get]
func (ctrl *QueryController) GetPointsHistory(c *gin.Context) {
	// TODO: 2024/11/22|sean|implement the handler
}
