package monitor

import (
	"context"
	"time"

	"github.com/kozhamseitova/balance-service/internal/models"
)

func(m *monitor) StartReservationMonitor(ctx context.Context) {
    ticker := time.NewTicker(1 * time.Minute) 

    for range ticker.C {
		m.checkReservations(ctx)
	}
	
}

func(m *monitor) checkReservations(ctx context.Context) {
	reservations, err := m.repo.GetNotRecognizedReservations(ctx)
	if err != nil {
		m.logger.Errorf(ctx, "[checkReservations] err: %v", err)
	}

    for _, reservation := range reservations {
        if time.Since(reservation.ReservedAt) > 5 * time.Minute {
            m.cancelReservation(ctx, reservation)
        }
    }
}

func(m *monitor) cancelReservation(ctx context.Context, reservation *models.Reservation) {
    err := m.repo.CanselReservation(ctx, reservation.UserID, reservation.ServiceID)
	if err != nil {
		m.logger.Errorf(ctx, "[cancelReservation] err: %v", err)
	}
	m.logger.Infof(ctx, "[cancelReservation] : %v", "success")
}

