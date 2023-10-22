package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNoRows = pgx.ErrNoRows
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Database, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Database{
		Pool: pool,
	}, nil
}