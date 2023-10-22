package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository interface {
}

type repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) Repository {
	return &repository{
		pool: pool,
	}
}