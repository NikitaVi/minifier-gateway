package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	grpcEnvHost         = "GRPC_HOST"
	grpcEnvPort         = "GRPC_PORT"
	grpcEnvRetries      = "GRPC_RETRIES"
	grpcEnvRetryTimeout = "GRPC_RETRY_TIMEOUT"
)

type GRPCConfig interface {
	Address() string
	Timeout() time.Duration
	Max() int
}

type grpcConfig struct {
	host         string
	port         string
	retries      int
	retryTimeout time.Duration
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcEnvHost)
	if host == "" {
		return nil, fmt.Errorf("environment variable %s must be set", grpcEnvHost)
	}

	port := os.Getenv(grpcEnvPort)
	if port == "" {
		return nil, fmt.Errorf("environment variable %s must be set", grpcEnvPort)
	}

	retries, err := strconv.Atoi(os.Getenv(grpcEnvRetries))
	if err != nil {
		return nil, fmt.Errorf("environment variable %s must be set", grpcEnvRetries)
	}

	retryTimeout, err := strconv.Atoi(os.Getenv(grpcEnvRetryTimeout))
	if err != nil {
		return nil, fmt.Errorf("environment variable %s must be set", grpcEnvRetryTimeout)
	}

	return &grpcConfig{
		host:         host,
		port:         port,
		retries:      retries,
		retryTimeout: time.Duration(retryTimeout),
	}, nil
}

func (c *grpcConfig) Address() string {
	return c.host + ":" + c.port
}

func (c *grpcConfig) Timeout() time.Duration {
	return c.retryTimeout
}

func (c *grpcConfig) Max() int {
	return c.retries
}
