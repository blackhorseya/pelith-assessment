package grpc

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/internal/domain/core/app/command"
	"github.com/blackhorseya/pelith-assessment/proto/core"
)

type campaignServerImpl struct {
	createCampaignHandler *command.CreateCampaignHandler
}

// NewCampaignServer is used to create a new campaign server.
func NewCampaignServer(createCampaignHandler *command.CreateCampaignHandler) core.CampaignServiceServer {
	return &campaignServerImpl{
		createCampaignHandler: createCampaignHandler,
	}
}

func (i *campaignServerImpl) StartCampaign(
	c context.Context,
	req *core.StartCampaignRequest,
) (*core.StartCampaignResponse, error) {
	err := i.createCampaignHandler.Handle(c, command.CreateCampaignCommand{
		Name:      req.Name,
		StartTime: req.StartTime.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	return &core.StartCampaignResponse{
		// TODO: 2024/11/20|sean|return campaign id
		Id: "you need to fill the campaign id",
	}, nil
}
