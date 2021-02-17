package orm

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jamillosantos/sqlf"
)

type (
	PgxExecer interface {
		Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	}

	PgxQueryer interface {
		Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
		QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	}

	PgxRowQueryer interface {
		QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	}

	PgxPreparer interface {
		Prepare(ctx context.Context, name, sql string) (sd *pgconn.StatementDescription, err error)
	}

	PgxDBRunner interface {
		PgxExecer
		PgxQueryer
		PgxRowQueryer
		PgxPreparer
	}

	ConnectionPgx interface {
		PgxDBRunner
		DB() PgxDBProxy
		Builder() sqlf.Builder
		Begin(ctx context.Context) (TxPgxProxy, error)
		BeginTx(ctx context.Context, opts pgx.TxOptions) (TxPgxProxy, error)
	}

	PgxDBProxy interface {
		PgxDBRunner
		Begin(ctx context.Context) (pgx.Tx, error)
		BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	}

	baseConnection struct {
		_db      PgxDBProxy
		_builder sqlf.Builder
	}
)

// NewConnectionPgx will create a new instance of the `*BaseConnection`.
func NewConnectionPgx(db PgxDBProxy, builder sqlf.Builder) ConnectionPgx {
	return &baseConnection{
		_db:      db,
		_builder: builder,
	}
}

// DB returns the real connection object for the database connection.
func (conn *baseConnection) DB() PgxDBProxy {
	return conn._db
}

// Builder returns the Statement Builder used to generate the queries for this connection.
func (conn *baseConnection) Builder() sqlf.Builder {
	return conn._builder
}

// Begin starts a transaction.
func (conn *baseConnection) Begin(ctx context.Context) (TxPgxProxy, error) {
	tx, err := conn._db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return NewPgxTx(tx), nil
}

// BeginTx starts a transaction with more options.
func (conn *baseConnection) BeginTx(ctx context.Context, opts pgx.TxOptions) (TxPgxProxy, error) {
	tx, err := conn._db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewPgxTx(tx), nil
}

func (conn *baseConnection) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return conn._db.Exec(ctx, query, args...)
}

func (conn *baseConnection) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return conn._db.Query(ctx, sql, args...)
}

func (conn *baseConnection) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return conn._db.QueryFunc(ctx, sql, args, scans, f)
}

func (conn *baseConnection) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return conn._db.QueryRow(ctx, query, args...)
}

func (conn *baseConnection) Prepare(ctx context.Context, name, query string) (*pgconn.StatementDescription, error) {
	return conn._db.Prepare(ctx, name, query)
}
