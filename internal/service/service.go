package service

import (
	"github.com/kozhamseitova/balance-service/internal/config"
	"github.com/kozhamseitova/balance-service/internal/repository"
	"github.com/kozhamseitova/balance-service/pkg/logger"
)

type Service interface {
}

type service struct {
	config *config.Config
	logger logger.Logger
	repo   repository.Repository
}

func New(config *config.Config, logger logger.Logger, repo repository.Repository) Service {
	return &service{
		config: config,
		logger: logger,
		repo:   repo,
	}
}