package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (q *Queries) ExecuteTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := q.db.(*pgxpool.Pool).Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := fn(ctx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
