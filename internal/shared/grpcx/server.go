package grpcx

import (
	"net"

	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// InitServers define register handler
type InitServers func(s *grpc.Server)

// Server represents the grpc server.
type Server struct {
	grpcserver *grpc.Server
	addr       string
}

// NewServer creates a new grpc server.
func NewServer(app *configx.Application, init InitServers, healthServer grpc_health_v1.HealthServer) (*Server, error) {
	logger := zap.L()
	server := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
	)

	init(server)

	// register health server
	grpc_health_v1.RegisterHealthServer(server, healthServer)

	// register reflection service on gRPC server.
	reflection.Register(server)

	return &Server{
		grpcserver: server,
		addr:       app.GRPC.GetAddr(),
	}, nil
}

// Start begins the server.
func (s *Server) Start(ctx contextx.Contextx) error {
	go func() {
		ctx.Info("grpc server start", zap.String("addr", s.addr))

		listen, err := net.Listen("tcp", s.addr)
		if err != nil {
			ctx.Fatal("grpc server listen error", zap.Error(err))
		}

		err = s.grpcserver.Serve(listen)
		if err != nil {
			ctx.Fatal("grpc server serve error", zap.Error(err))
		}
	}()

	return nil
}

// Stop stops the server.
func (s *Server) Stop(ctx contextx.Contextx) error {
	ctx.Info("grpc server stop")

	s.grpcserver.Stop()

	return nil
}
