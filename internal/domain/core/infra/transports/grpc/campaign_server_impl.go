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

func (i *campaignServerImpl) CreateCampaign(
	c context.Context,
	req *core.CreateCampaignRequest,
) (*core.CreateCampaignResponse, error) {
	id, err := i.createCampaignHandler.Handle(c, command.CreateCampaignCommand{
		Name:       req.Name,
		StartTime:  req.StartTime.AsTime(),
		Mode:       req.Mode,
		TargetPool: req.TargetPool,
		MinAmount:  req.MinAmount,
	})
	if err != nil {
		return nil, err
	}

	return &core.CreateCampaignResponse{
		Id: id,
	}, nil
}

func (i *campaignServerImpl) StartCampaign(
	c context.Context,
	req *core.StartCampaignRequest,
) (*core.StartCampaignResponse, error) {
	id, err := i.createCampaignHandler.Handle(c, command.CreateCampaignCommand{
		Name:       req.Name,
		StartTime:  req.StartTime.AsTime(),
		Mode:       req.Mode,
		TargetPool: req.TargetPool,
		MinAmount:  req.MinAmount,
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
