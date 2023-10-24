package service

import (
	"context"

	"github.com/kozhamseitova/balance-service/internal/models"
)

func(s *service) GetBalanceByUserID(ctx context.Context, id int) (int, error) {
	return s.repo.GetBalanceByUserID(ctx, id)
}

func(s *service) DepositFunds(ctx context.Context, id int, amount int) error {
	return s.repo.DepositFunds(ctx, id, amount)
}

func(s *service) ReserveFunds(ctx context.Context, userID, serviceID, orderID, price int) error {
	return s.repo.ReserveFunds(ctx, userID, serviceID, orderID, price)
}

func(s *service) RecognizeRevenue(ctx context.Context, userID, serviceID, orderID, price int) error {
	return s.repo.RecognizeRevenue(ctx, userID, serviceID, orderID, price)
}

func(s *service) GetReport(ctx context.Context) ([]*models.Report, error){
	return s.repo.GetReport(ctx)
}

func(s *service) CanselReservation(ctx context.Context, userID, serviceId int) error {
	return s.repo.CanselReservation(ctx, userID, serviceId)
}