package main

import (
	"context"
	"github.com/NikitaVi/minifier-gateway/internal/app"
	"github.com/NikitaVi/minifier-gateway/internal/logger"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	logger.Info("Starting server")

	err = a.Run(ctx)
	if err != nil {
		panic(err)
	}
}
