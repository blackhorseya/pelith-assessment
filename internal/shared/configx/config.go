package configx

import (
	"github.com/blackhorseya/pelith-assessment/pkg/logger"
	"github.com/spf13/viper"
)

// Configx is the application configuration.
type Configx struct {
	viper *viper.Viper

	// Logger is the logger configuration.
	Logger logger.Options `json:"logger" yaml:"logger" mapstructure:"logger"`
}
