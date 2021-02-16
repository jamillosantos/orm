package orm

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type (
	TxPgxProxy interface {
		PgxDBRunner
		Commit(context.Context) error
		Rollback(context.Context) error
	}

	BaseTxRunner struct {
		tx pgx.Tx
	}
)

// NewPgxTx returns a new transaction abstraction.
func NewPgxTx(tx pgx.Tx) TxPgxProxy {
	return &BaseTxRunner{
		tx: tx,
	}
}

func (tx *BaseTxRunner) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return tx.tx.Exec(ctx, query, args...)
}

func (tx *BaseTxRunner) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return tx.tx.Query(ctx, sql, args...)
}

func (tx *BaseTxRunner) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return tx.tx.QueryFunc(ctx, sql, args, scans, f)
}

func (tx *BaseTxRunner) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return tx.tx.QueryRow(ctx, query, args...)
}

func (tx *BaseTxRunner) Prepare(ctx context.Context, name, query string) (*pgconn.StatementDescription, error) {
	return tx.tx.Prepare(ctx, name, query)
}

func (tx *BaseTxRunner) Commit(ctx context.Context) error {
	return tx.tx.Commit(ctx)
}

func (tx *BaseTxRunner) Rollback(ctx context.Context) error {
	return tx.tx.Rollback(ctx)
}
