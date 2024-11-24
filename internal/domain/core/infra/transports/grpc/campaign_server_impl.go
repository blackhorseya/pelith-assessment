package grpc

import (
	"context"
	"strconv"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/query"
	"github.com/blackhorseya/pelith-assessment/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type campaignServerImpl struct {
	createCampaignHandler *command.CreateCampaignHandler
	addTaskHandler        *command.AddTaskHandler
	startCampaignHandler  *command.StartCampaignHandler

	campaignGetter query.CampaignGetter
}

// NewCampaignServer is used to create a new campaign server.
func NewCampaignServer(
	createCampaignHandler *command.CreateCampaignHandler,
	addTaskHandler *command.AddTaskHandler,
	startCampaignHandler *command.StartCampaignHandler,
	campaignGetter query.CampaignGetter,
) core.CampaignServiceServer {
	return &campaignServerImpl{
		createCampaignHandler: createCampaignHandler,
		addTaskHandler:        addTaskHandler,
		startCampaignHandler:  startCampaignHandler,
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
	id, err := i.startCampaignHandler.Handle(c, command.StartCampaignCommand{
		ID: req.Id,
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

func (i *campaignServerImpl) ListCampaigns(
	req *core.ListCampaignsRequest,
	stream grpc.ServerStreamingServer[core.GetCampaignResponse],
) error {
	items, total, err := i.campaignGetter.List(stream.Context(), query.ListCampaignCondition{
		// TODO: 2024/11/24|sean|pass the condition
	})
	if err != nil {
		return err
	}

	err = stream.SendHeader(metadata.New(map[string]string{"total": strconv.Itoa(total)}))
	if err != nil {
		return err
	}

	for _, item := range items {
		tasks := make([]*model.Task, 0, len(item.Tasks))
		for _, task := range item.Tasks {
			tasks = append(tasks, &task.Task)
		}

		err = stream.Send(&core.GetCampaignResponse{
			Campaign: &item.Campaign,
			Tasks:    tasks,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
