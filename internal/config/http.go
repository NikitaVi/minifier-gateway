package config

import (
	"fmt"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

type HTTPConfig interface {
	Address() string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("environment variable %s must be set", httpHostEnvName)
	}

	port := os.Getenv(httpPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("environment variable %s must be set", httpPortEnvName)
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (c *httpConfig) Address() string {
	return c.host + ":" + c.port
}
