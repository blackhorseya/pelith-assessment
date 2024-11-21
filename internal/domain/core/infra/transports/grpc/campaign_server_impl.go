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
	campaignGetter        query.CampaignGetter
}

// NewCampaignServer is used to create a new campaign server.
func NewCampaignServer(
	createCampaignHandler *command.CreateCampaignHandler,
	campaignGetter query.CampaignGetter,
) core.CampaignServiceServer {
	return &campaignServerImpl{
		createCampaignHandler: createCampaignHandler,
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
	// TODO: 2024/11/21|sean|add tasks for campaign
	panic("implement me")
}
