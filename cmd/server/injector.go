package server

import (
	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
)

// Injector is the injector for server
type Injector struct {
	C *configx.Configx
	A *configx.Application
}
