package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kozhamseitova/balance-service/internal/models"
	"github.com/kozhamseitova/balance-service/pkg/logger"
)

type Repository interface {
	GetBalanceByUserID(ctx context.Context, id int) (int, error)
	DepositFunds(ctx context.Context, id int, amount int) error
	ReserveFunds(ctx context.Context, userID, serviceID, orderID, price int) error
	RecognizeRevenue(ctx context.Context, userID, serviceID, orderID, price int) error
	GetReport(ctx context.Context) ([]*models.Report, error)
	CanselReservation(ctx context.Context, userID, serviceId int) error
 }

type repository struct {
	pool *pgxpool.Pool
	logger logger.Logger
}

func New(pool *pgxpool.Pool, logger logger.Logger) Repository {
	return &repository{
		pool: pool,
		logger: logger,
	}
}