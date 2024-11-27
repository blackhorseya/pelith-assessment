package grpc

import (
	"fmt"

	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"github.com/blackhorseya/pelith-assessment/proto/core"
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
