package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/kozhamseitova/balance-service/utils"
)

const (
	usersTable = "users"
	reservationsTable = "reservations"
	revenueTable = "revenuereports"
)

func (r *repository) DepositFunds(ctx context.Context, id int, amount int) error {
	tx, err := r.pool.Begin(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[CreditBalance] failed to start transaction: %v", err)
        return utils.ErrInternalError
    }
    defer tx.Rollback(ctx) 

    query := fmt.Sprintf("UPDATE %s SET balance = balance + $1 WHERE id = $2", usersTable)
    _, err = tx.Exec(ctx, query, amount, id)
    if err != nil {
        r.logger.Errorf(ctx, "[CreditBalance] failed to update balance: %v", err)
        return utils.ErrInternalError
    }

    err = tx.Commit(ctx)
    if err != nil {
        r.logger.Errorf(ctx, "[CreditBalance] failed to commit transaction: %v", err)
        return utils.ErrInternalError
    }

	return nil
}

func (r *repository) GetBalanceByUserID(ctx context.Context, id int) (int, error) {
	var balance int

	query := fmt.Sprintf(`SELECT balance from %s WHERE id = $1 LIMIT 1`, usersTable)

	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(&balance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, utils.ErrNotFound
		}
		r.logger.Errorf(ctx, "[GetUserBalanceByID] err: %v", err)
		return 0, utils.ErrInternalError
	}

	
	return balance, nil
}
