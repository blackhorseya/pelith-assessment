package grpc

import (
	"github.com/blackhorseya/pelith-assessment/internal/shared/grpcx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// NewInitServersFn is used to create a new init servers function.
func NewInitServersFn() grpcx.InitServers {
	return func(s *grpc.Server) {
		// TODO: 2024/11/20|sean|register grpc server
	}
}

// NewHealthServer is used to create a new health server.
func NewHealthServer() grpc_health_v1.HealthServer {
	return health.NewServer()
}
