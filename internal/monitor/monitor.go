package monitor

import (
	"context"

	"github.com/kozhamseitova/balance-service/internal/models"
	"github.com/kozhamseitova/balance-service/internal/repository"
	"github.com/kozhamseitova/balance-service/pkg/logger"
)

type Monitor interface {
	StartReservationMonitor(ctx context.Context)
	checkReservations(ctx context.Context)
	cancelReservation(ctx context.Context, reservation *models.Reservation)
}

type monitor struct {
	logger logger.Logger
	repo   repository.Repository
}

func New(logger logger.Logger, repo repository.Repository) Monitor {
	return &monitor{
		logger: logger,
		repo:   repo,
	}
}