package http

import (
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

// GetTasksStatus is the handler to get tasks status
func (ctrl *QueryController) GetTasksStatus(c *gin.Context) {
	// TODO: 2024/11/22|sean|implement the handler
}

// GetPointsHistory is the handler to get points history
func (ctrl *QueryController) GetPointsHistory(c *gin.Context) {
	// TODO: 2024/11/22|sean|implement the handler
}
