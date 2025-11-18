package app

import (
	"context"
	"github.com/NikitaVi/minifier-gateway/internal/clients/auth /grpc"
	"github.com/NikitaVi/minifier-gateway/internal/config"
	"log"
)

type ServiceProvider struct {
	grpcClient grpc.GRPCClient

	grpcConfig config.GRPCConfig
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *ServiceProvider) GRPCClient(ctx context.Context) *grpc.GRPCClient {
	cc, err := grpc.NewGRPCClient(ctx, s.GRPCConfig().Address(), s.GRPCConfig().Max(), s.GRPCConfig().Timeout())
	if err != nil {
		log.Fatal(err)
	}

	return cc
}
