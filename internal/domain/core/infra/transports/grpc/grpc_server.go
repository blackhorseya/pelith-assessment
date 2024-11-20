package grpc

import (
	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"github.com/blackhorseya/pelith-assessment/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// NewInitServersFn is used to create a new init servers function.
func NewInitServersFn(campaignServer core.CampaignServiceServer) grpcx.InitServers {
	return func(s *grpc.Server) {
		core.RegisterCampaignServiceServer(s, campaignServer)
	}
}

// NewHealthServer is used to create a new health server.
func NewHealthServer() grpc_health_v1.HealthServer {
	return health.NewServer()
}
