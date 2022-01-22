package mysql

import (
	"context"
	"database/sql"
)

// ExecTransaction template to execute transaction
func ExecTransaction(
	ctx context.Context,
	tx *sql.Tx,
	exec func(context.Context, *sql.Tx) (sql.Result, error)) (res sql.Result, err error) {
	if res, err = exec(ctx, tx); err != nil {
		_ = tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}
