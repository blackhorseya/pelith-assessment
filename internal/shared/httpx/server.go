package httpx

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/blackhorseya/pelith-assessment/pkg/contextx"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// InitRoutes is a function to initialize routes.
type InitRoutes func(router *gin.Engine)

// GinServer is an HTTP server.
type GinServer struct {
	httpserver *http.Server

	// Router is the gin engine.
	Router *gin.Engine
}

// NewGinServer is used to create a new HTTP server.
func NewGinServer(app *configx.Application, init InitRoutes) (*GinServer, error) {
	gin.SetMode(app.HTTP.Mode)

	router := gin.New()
	router.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(zap.L(), true))

	init(router)

	httpserver := &http.Server{
		Addr:              app.HTTP.GetAddr(),
		Handler:           router,
		ReadHeaderTimeout: time.Second,
	}

	return &GinServer{
		httpserver: httpserver,
		Router:     router,
	}, nil
}

// Start starts the server.
func (s *GinServer) Start(ctx contextx.Contextx) error {
	go func() {
		ctx.Info("http server is starting", zap.String("addr", s.httpserver.Addr))

		if err := s.httpserver.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Fatal("start http server error", zap.Error(err))
			return
		}
	}()

	return nil
}

// Stop halts the server.
func (s *GinServer) Stop(ctx contextx.Contextx) error {
	timeout, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	ctx.Info("http server is stopping")

	if err := s.httpserver.Shutdown(timeout); err != nil {
		ctx.Error("stop http server error", zap.Error(err))
		return err
	}

	return nil
}

// GetAddr is used to get the server address.
func (s *GinServer) GetAddr() string {
	return s.httpserver.Addr
}
