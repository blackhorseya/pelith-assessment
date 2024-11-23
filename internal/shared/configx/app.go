package configx

import (
	"fmt"

	"github.com/blackhorseya/pelith-assessment/pkg/netx"
)

// Application is the application configuration.
type Application struct {
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	HTTP HTTP `json:"http" yaml:"http" mapstructure:"http"`
	GRPC GRPC `json:"grpc" yaml:"grpc" mapstructure:"grpc"`

	Storage struct {
		DSN string `json:"dsn" yaml:"dsn" mapstructure:"dsn"`
	} `json:"storage" yaml:"storage" mapstructure:"storage"`

	Etherscan struct {
		APIKey string `json:"api_key" yaml:"apiKey" mapstructure:"apiKey"`
	} `json:"etherscan" yaml:"etherscan" mapstructure:"etherscan"`

	Infura struct {
		ProjectID string `json:"project_id" yaml:"projectID" mapstructure:"projectID"`
	} `json:"infura" yaml:"infura" mapstructure:"infura"`
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

// GRPC is the gRPC configuration.
type GRPC struct {
	URL  string `json:"url" yaml:"url"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

// GetAddr is used to get the gRPC address.
func (x *GRPC) GetAddr() string {
	if x.Host == "" {
		x.Host = "0.0.0.0"
	}

	if x.Port == 0 {
		x.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", x.Host, x.Port)
}
