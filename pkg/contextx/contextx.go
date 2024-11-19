package contextx

import (
	"context"

	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

// Contextx extends google's context to support logging methods.
type Contextx struct {
	context.Context
	*zap.Logger
}

// WithContext returns a copy of parent in which the context is set to ctx.
func WithContext(c context.Context) Contextx {
	return Contextx{
		Context: c,
		Logger:  zap.L(),
	}
}
