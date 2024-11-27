package grpcx

import (
	"errors"
	"fmt"

	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is the grpc client
type Client struct {
	config *configx.Configx
}

// NewClient is used to create a new grpc client
func NewClient(config *configx.Configx) (*Client, error) {
	return &Client{
		config: config,
	}, nil
}

// Dial is used to dial the grpc service
func (x *Client) Dial(service string) (*grpc.ClientConn, error) {
	app, err := x.config.GetService(service)
	if err != nil {
		return nil, err
	}

	if app.GRPC.URL == "" || app.GRPC.Port == 0 {
		return nil, errors.New("grpc url or port is empty")
	}

	target := fmt.Sprintf("%s:%d", app.GRPC.URL, app.GRPC.Port)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_prometheus.UnaryClientInterceptor,
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_prometheus.StreamClientInterceptor,
		)),
	}

	return grpc.NewClient(target, options...)
}
