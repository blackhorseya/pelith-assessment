package configx

import (
	"fmt"

	"github.com/blackhorseya/pelith-assessment/pkg/netx"
)

// Application is the application configuration.
type Application struct {
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	HTTP HTTP `json:"http" yaml:"http" mapstructure:"http"`

	Storage struct {
		DSN string `json:"dsn" yaml:"dsn" mapstructure:"dsn"`
	} `json:"storage" yaml:"storage" mapstructure:"storage"`
}

// HTTP defines the http struct.
type HTTP struct {
	URL  string `json:"url" yaml:"url"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GetAddr is used to get the http address.
func (http *HTTP) GetAddr() string {
	if http.Host == "" {
		http.Host = "0.0.0.0"
	}

	if http.Port == 0 {
		http.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", http.Host, http.Port)
}
