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

// LoadConfig is used to load the configuration.
func LoadConfig(path string) (*Configx, error) {
	v := viper.GetViper()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	// Load the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Configx
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	config.viper = v

	// Set the logger
	if err := logger.Init(config.Logger); err != nil {
		return nil, err
	}

	return &config, nil
}
