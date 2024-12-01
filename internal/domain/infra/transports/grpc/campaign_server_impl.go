package grpc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	"github.com/blackhorseya/pelith-assessment/internal/domain/app/command"
	"github.com/blackhorseya/pelith-assessment/internal/domain/app/query"
	"github.com/blackhorseya/pelith-assessment/internal/domain/repo"
	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"github.com/blackhorseya/pelith-assessment/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// NewCampaignServiceClient is used to create a new campaign service client.
func NewCampaignServiceClient(client *grpcx.Client) (core.CampaignServiceClient, error) {
	const service = "server"
	conn, err := client.Dial(service)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %s: %w", service, err)
	}

	return core.NewCampaignServiceClient(conn), nil
}

type campaignServerImpl struct {
	createCampaignHandler *command.CreateCampaignHandler
	addTaskHandler        *command.AddTaskHandler
	startCampaignHandler  *command.StartCampaignHandler
	runBacktestHandler    *command.RunBacktestHandler

	campaignGetter query.CampaignGetter
	campaignFinder repo.CampaignFinder
}

// NewCampaignServer is used to create a new campaign server.
func NewCampaignServer(
	createCampaignHandler *command.CreateCampaignHandler,
	addTaskHandler *command.AddTaskHandler,
	startCampaignHandler *command.StartCampaignHandler,
	runBacktestHandler *command.RunBacktestHandler,
	campaignGetter query.CampaignGetter,
	campaignFinder repo.CampaignFinder,
) core.CampaignServiceServer {
	return &campaignServerImpl{
		createCampaignHandler: createCampaignHandler,
		addTaskHandler:        addTaskHandler,
		startCampaignHandler:  startCampaignHandler,
		runBacktestHandler:    runBacktestHandler,
		campaignGetter:        campaignGetter,
		campaignFinder:        campaignFinder,
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

	tasks := make([]*model.Task, 0, len(campaign.Tasks()))
	for _, task := range campaign.Tasks() {
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
	items, total, err := i.campaignFinder.List(stream.Context(), repo.ListCampaignCondition{})
	if err != nil {
		return err
	}

	err = stream.SendHeader(metadata.New(map[string]string{"total": strconv.Itoa(int(total))}))
	if err != nil {
		return err
	}

	for _, item := range items {
		tasks := make([]*model.Task, 0, len(item.Tasks()))
		for _, task := range item.Tasks() {
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

func (i *campaignServerImpl) RunBacktestByCampaign(
	req *core.GetCampaignRequest,
	stream grpc.ServerStreamingServer[core.BacktestResultResponse],
) error {
	resultCh := make(chan *model.Reward)
	var err error
	go func() {
		err = i.runBacktestHandler.Handle(stream.Context(), req.Id, resultCh)
		close(resultCh)
	}()

	for result := range resultCh {
		err = stream.Send(&core.BacktestResultResponse{
			Reward: result,
		})
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	return nil
}
