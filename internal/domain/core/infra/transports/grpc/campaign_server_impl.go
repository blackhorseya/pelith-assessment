package grpc

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/proto/core"
)

type campaignServerImpl struct {
}

// NewCampaignServer is used to create a new campaign server.
func NewCampaignServer() core.CampaignServiceServer {
	return &campaignServerImpl{}
}

func (i *campaignServerImpl) StartCampaign(
	c context.Context,
	req *core.StartCampaignRequest,
) (*core.StartCampaignResponse, error) {
	// TODO: 2024/11/20|sean|implement me
	panic("implement me")
}
