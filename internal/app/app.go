package app

import (
	"context"
	"github.com/NikitaVi/minifier-gateway/internal/config"
	"github.com/NikitaVi/minifier-gateway/internal/logger"
	"github.com/gin-gonic/gin"
	"sync"
)

type App struct {
	serviceProvide *ServiceProvider
	httpServer     *gin.Engine
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initApp(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initApp(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.initHTTP,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvide = NewServiceProvider()
	return nil
}

func (a *App) initHTTP(ctx context.Context) error {
	if a.httpServer == nil {
		a.httpServer = gin.New()
	}
	//TODO: finish when controller would be ready
	return nil
}

func (a *App) initLogger(ctx context.Context) error {
	return logger.Init()
}

func (a *App) Run(ctx context.Context) error {
	wg := &sync.WaitGroup{}

	wg.Add(1)

	func() {
		defer wg.Done()
		if err := a.runHTTP(ctx); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) runHTTP(ctx context.Context) error {
	if err := a.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
