package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx"
	"github.com/kozhamseitova/balance-service/internal/models"
	"github.com/kozhamseitova/balance-service/utils"
)

func (r *repository) ReserveFunds(ctx context.Context, userID, serviceID, orderID, price int) error {
	tx, err := r.pool.Begin(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[ReserveFunds] failed to start transaction: %v", err)
        return utils.ErrInternalError
    }
    defer tx.Rollback(ctx)

    // 1. Withdraw funds from user balance
    query := fmt.Sprintf(`UPDATE %s SET balance = balance - $1 WHERE id = $2 and balance >= $1`, usersTable)
    result, err := tx.Exec(ctx, query, price, userID)
    if err != nil {
        r.logger.Errorf(ctx, "[ReserveFunds] failed to update user balance: %v", err)
        return utils.ErrInternalError
    }

	affectedRows := result.RowsAffected()
	if affectedRows == 0 {
		return utils.ErrInsufficientFunds
	}

    // 2. Create reservation record

	//check if reservation already exists
	checkReservationQuery := fmt.Sprintf(`SELECT id from %s where user_id = $1 and service_id = $2`, reservationsTable)
	resultCheckReserve, err := tx.Exec(ctx, checkReservationQuery, userID, serviceID)
	if err != nil {
        r.logger.Errorf(ctx, "[ReserveFunds] failed to check if reservation exists: %v", err)
        return utils.ErrInternalError
    }

	affectedRows = resultCheckReserve.RowsAffected()
	if affectedRows != 0 {
		return utils.ErrDuplicateKey
	}

    reservationQuery := fmt.Sprintf(`INSERT INTO %s (user_id, service_id, order_id, amount) VALUES ($1, $2, $3, $4) RETURNING id`, reservationsTable)
    _, err = tx.Exec(ctx, reservationQuery, userID, serviceID, orderID, price)
    if err != nil {
        r.logger.Errorf(ctx, "[ReserveFunds] failed to create reservation: %v", err)
        return utils.ErrInternalError
    }

    err = tx.Commit(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[ReserveFunds] failed to commit transaction: %v", err)
        return utils.ErrInternalError
    }

    return nil
}

func (r *repository) RecognizeRevenue(ctx context.Context, userID, serviceID, orderID, price int) error {
	tx, err := r.pool.Begin(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[RecognizeRevenue] failed to start transaction: %v", err)
        return utils.ErrInternalError
    }
    defer tx.Rollback(ctx)

    // 1. Check reservation
    reservationQuery := fmt.Sprintf(`SELECT amount FROM %s WHERE user_id = $1 AND service_id = $2 AND order_id = $3 AND recognized_at IS NULL`, reservationsTable)
    var reservedAmount int
    err = tx.QueryRow(ctx, reservationQuery, userID, serviceID, orderID).Scan(&reservedAmount)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            r.logger.Errorf(ctx, "[RecognizeRevenue] reservation not found")
            return utils.ErrReservationNotFound
        } else {
            r.logger.Errorf(ctx, "[RecognizeRevenue] failed to query reservation: %v", err)
            return utils.ErrInternalError
        }
    }

    if price > reservedAmount {
        r.logger.Errorf(ctx, "[RecognizeRevenue] amount exceeds reserved amount")
        return utils.ErrInvalidAmount
    }

    // 2. Update reservation record
    updateReservationQuery := fmt.Sprintf("UPDATE %s SET amount = amount - $1, recognized_at = NOW()  WHERE user_id = $2 AND service_id = $3 AND order_id = $4", reservationsTable)
    _, err = tx.Exec(ctx, updateReservationQuery, price, userID, serviceID, orderID)
    if err != nil {
        r.logger.Errorf(ctx, "[RecognizeRevenue] failed to update reservation: %v", err)
        return utils.ErrInternalError
    }

    // 3. Add to revenue
    revenueQuery := fmt.Sprintf(`INSERT INTO %s (user_id, service_id, order_id, amount) VALUES ($1, $2, $3, $4)`, revenueTable)
    _, err = tx.Exec(ctx, revenueQuery, userID, serviceID, orderID, price)
    if err != nil {
        r.logger.Errorf(ctx, "[RecognizeRevenue] failed to create revenue record: %v", err)
        return utils.ErrInternalError
    }

    err = tx.Commit(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[RecognizeRevenue] failed to commit transaction: %v", err)
        return utils.ErrInternalError
    }

	return nil
}

func(r *repository) GetReport(ctx context.Context) ([]*models.Report, error) {
	var report []*models.Report

	query := fmt.Sprintf(`SELECT * from %s order by user_id`, revenueTable)
	err := pgxscan.Select(ctx, r.pool, &report, query)
	if err != nil {
		r.logger.Errorf(ctx, "[GetReport] err: %v", err)
		return nil, utils.ErrInternalError
	}

	
	return report, nil
}

func(r *repository) CanselReservation(ctx context.Context, userID, serviceId int) error {
	tx, err := r.pool.Begin(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[CanselReservatioin] failed to start transaction: %v", err)
        return utils.ErrInternalError
    }
    defer tx.Rollback(ctx)

    // 1. Delete reservation record
	var amount int
    reservationQuery := fmt.Sprintf(`DELETE from %s WHERE user_id = $1 and service_id = $2 and recognized_at is null RETURNING amount`, reservationsTable)
    err = tx.QueryRow(ctx, reservationQuery, userID, serviceId).Scan(&amount)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            r.logger.Errorf(ctx, "[RecognizeRevenue] reservation not found")
            return utils.ErrReservationNotFound
        }
        r.logger.Errorf(ctx, "[CanselReservatioin] failed to delete reservation: %v", err)
        return utils.ErrInternalError
    }

	// 2. Deposit funds from user balance
    query := fmt.Sprintf(`UPDATE %s SET balance = balance + $1 WHERE id = $2`, usersTable)
    result, err := tx.Exec(ctx, query, amount, userID)
    if err != nil {
        r.logger.Errorf(ctx, "[CanselReservatioin] failed to update user balance: %v", err)
        return utils.ErrInternalError
    }

	affectedRows := result.RowsAffected()
	if affectedRows == 0 {
		return utils.ErrInternalError
	}

    err = tx.Commit(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[CanselReservatioin] failed to commit transaction: %v", err)
        return utils.ErrInternalError
    }
	return nil
}

func (r *repository) GetNotRecognizedReservations(ctx context.Context) ([]*models.Reservation, error) {
	var reservations []*models.Reservation

	query := fmt.Sprintf(`SELECT id, user_id, service_id, order_id, amount, reserved_at from %s WHERE recognized_at is NULL`, reservationsTable)
	err := pgxscan.Select(ctx, r.pool, &reservations, query)
	if err != nil {
		r.logger.Errorf(ctx, "[GetNotRecognizedReservations] err: %v", err)
		return nil, utils.ErrInternalError
	}
	
	return reservations, nil
}