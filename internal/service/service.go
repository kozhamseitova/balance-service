package service

import (
	"context"

	"github.com/kozhamseitova/balance-service/internal/config"
	"github.com/kozhamseitova/balance-service/internal/models"
	"github.com/kozhamseitova/balance-service/internal/repository"
	"github.com/kozhamseitova/balance-service/pkg/logger"
)

type Service interface {
	GetBalanceByUserID(ctx context.Context, id int) (int, error)
	DepositFunds(ctx context.Context, id int, amount int) error
	// checkBalance(ctx context.Context, id, price int) (bool, error)
	ReserveFunds(ctx context.Context, userID, serviceID, orderID, price int) error
	RecognizeRevenue(ctx context.Context, userID, serviceID, orderID, price int) error
	GetReport(ctx context.Context) ([]*models.Report, error)
	CanselReservation(ctx context.Context, userID, serviceId int) error
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