package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/learn-video/streaming-platform/internal/config"
)

func NewPool(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}

func NewQuerier(pool *pgxpool.Pool) *Queries {
	return New(pool)
}
