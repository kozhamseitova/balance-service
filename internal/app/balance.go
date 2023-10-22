package app

import (
	"context"
	"fmt"

	"github.com/kozhamseitova/balance-service/internal/config"
	"github.com/kozhamseitova/balance-service/internal/handler"
	"github.com/kozhamseitova/balance-service/internal/repository"
	"github.com/kozhamseitova/balance-service/internal/server"
	"github.com/kozhamseitova/balance-service/internal/service"
	"github.com/kozhamseitova/balance-service/pkg/db"
	"github.com/kozhamseitova/balance-service/pkg/logger"
)

type App struct {
	server *server.Server
	config *config.Config
}

func New(ctx context.Context) (*App, error) {
	config, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to init config: %w", err)
	}
	logger, err := logger.New(config.App.LogLevel, config.App.Environment)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}

	database, err := db.New(ctx, config.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}
	fmt.Println(database.Pool)

	repo := repository.New(database.Pool)
	service := service.New(config, logger, repo)
	handler := handler.New(logger, service)

	server := server.New(config, logger, handler)

	return &App{
		server: server,
		config: config,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.server.Run(a.config.App.Port)
}

func (a *App) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}