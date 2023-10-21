package handler

import (
	"github.com/kozhamseitova/balance-service/internal/service"
	"github.com/kozhamseitova/balance-service/pkg/logger"
)

type Handler struct {
	service service.Service
	logger  logger.Logger
}

func New(logger logger.Logger, service service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}