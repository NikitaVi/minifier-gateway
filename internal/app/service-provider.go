package app

import (
	"context"
	"github.com/NikitaVi/minifier-gateway/internal/api"
	"github.com/NikitaVi/minifier-gateway/internal/api/auth"
	"github.com/NikitaVi/minifier-gateway/internal/api/uploads"
	"github.com/NikitaVi/minifier-gateway/internal/api/user"
	"github.com/NikitaVi/minifier-gateway/internal/clients/auth/grpc"
	"github.com/NikitaVi/minifier-gateway/internal/config"
	"log"
)

type ServiceProvider struct {
	grpcClient grpc.GRPCClient

	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	uploadsController *uploads.Controller
	userController    *user.Controller
	authController    *auth.Controller
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

func (s *ServiceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *ServiceProvider) GRPCClient(ctx context.Context) *grpc.GRPCClient {
	cc, err := grpc.NewGRPCClient(ctx, s.GRPCConfig().Address(), s.GRPCConfig().Max(), s.GRPCConfig().Timeout())
	if err != nil {
		log.Fatal(err)
	}

	return cc
}

func (s *ServiceProvider) UploadsController(ctx context.Context) *uploads.Controller {
	if s.uploadsController == nil {
		s.uploadsController = uploads.NewController()
	}

	return s.uploadsController
}

func (s *ServiceProvider) UserController(ctx context.Context) *user.Controller {
	if s.userController == nil {
		s.userController = user.NewController()
	}

	return s.userController
}

func (s *ServiceProvider) AuthController(ctx context.Context) *auth.Controller {
	if s.authController == nil {
		s.authController = auth.NewController()
	}

	return s.authController
}

func (s *ServiceProvider) AllControllers(ctx context.Context) []api.Router {
	return []api.Router{
		s.AuthController(ctx),
		s.UploadsController(ctx),
		s.UserController(ctx),
	}
}
