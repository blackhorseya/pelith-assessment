package grpcx

import (
	"google.golang.org/grpc"
)

// InitServers define register handler
type InitServers func(s *grpc.Server)

// Server represents the grpc server.
type Server struct {
	grpcserver *grpc.Server
	addr       string
}
