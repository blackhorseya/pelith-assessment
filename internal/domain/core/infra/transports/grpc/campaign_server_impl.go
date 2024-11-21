package grpc

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/proto/core"
)

type campaignServerImpl struct {
	createCampaignHandler *command.CreateCampaignHandler
	addTaskHandler        *command.AddTaskHandler

	campaignGetter query.CampaignGetter
}

// NewCampaignServer is used to create a new campaign server.
func NewCampaignServer(
	createCampaignHandler *command.CreateCampaignHandler,
	addTaskHandler *command.AddTaskHandler,
	campaignGetter query.CampaignGetter,
) core.CampaignServiceServer {
	return &campaignServerImpl{
		createCampaignHandler: createCampaignHandler,
		addTaskHandler:        addTaskHandler,
		campaignGetter:        campaignGetter,
	}
}

func (i *campaignServerImpl) StartCampaign(
	c context.Context,
	req *core.StartCampaignRequest,
) (*core.StartCampaignResponse, error) {
	id, err := i.createCampaignHandler.Handle(c, command.CreateCampaignCommand{
		Name:      req.Name,
		StartTime: req.StartTime.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	return &core.StartCampaignResponse{
		Id: id,
	}, nil
}

func (i *campaignServerImpl) GetCampaign(
	c context.Context,
	req *core.GetCampaignRequest,
) (*core.GetCampaignResponse, error) {
	campaign, err := i.campaignGetter.GetByID(c, req.Id)
	if err != nil {
		return nil, err
	}

	tasks := make([]*model.Task, 0, len(campaign.Tasks))
	for _, task := range campaign.Tasks {
		tasks = append(tasks, &task.Task)
	}

	return &core.GetCampaignResponse{
		Campaign: &campaign.Campaign,
		Tasks:    tasks,
	}, nil
}

func (i *campaignServerImpl) AddTasksForCampaign(
	c context.Context,
	req *core.AddTasksForCampaignRequest,
) (*core.AddTasksForCampaignResponse, error) {
	tasks := make([]command.TaskCommand, 0, len(req.Tasks))
	for _, task := range req.Tasks {
		tasks = append(tasks, command.TaskCommand{
			Name:        task.Name,
			Description: task.Description,
			Type:        int32(task.Type),
			MinAmount:   task.Criteria.MinTransactionAmount,
			PoolID:      task.Criteria.PoolId,
		})
	}

	_, err := i.addTaskHandler.Handle(c, command.AddTaskCommand{
		CampaignID: req.CampaignId,
		Tasks:      tasks,
	})
	if err != nil {
		return nil, err
	}

	return &core.AddTasksForCampaignResponse{}, nil
}
