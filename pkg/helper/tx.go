package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CommitOrRollback(err error, ctx context.Context, tx pgx.Tx) error {
	if err != nil {
		tx.Rollback(ctx)
		return err
	} else {
		tx.Commit(ctx)
	}
	return nil
}
