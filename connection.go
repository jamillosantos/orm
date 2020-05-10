package orm

import (
	"context"
	"database/sql"
	"io"

	sq "github.com/Masterminds/squirrel"
)

type (
	Pinger interface {
		Ping() error
	}

	PingerContext interface {
		PingContext(ctx context.Context) error
	}

	Transactioner interface {
		Begin() (*sql.Tx, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	}

	Statisticer interface {
		Stats() sql.DBStats
	}

	QueryRower interface {
		QueryRow(sql string, args ...interface{}) *sql.Row
	}

	QueryRowerContext interface {
		QueryRowContext(ctx context.Context, sql string, args ...interface{}) *sql.Row
	}

	DBRunner interface {
		sq.Execer
		sq.ExecerContext
		sq.Queryer
		sq.QueryerContext
		sq.Preparer
		QueryRower
		QueryRowerContext
	}

	DBProxy interface {
		DBRunner
		io.Closer
		Statisticer
		Transactioner
	}

	Connection interface {
		DB() DBProxy
		Builder() sq.StatementBuilderType
		Begin() (TxProxy, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (TxProxy, error)
		sq.Execer
		sq.ExecerContext
		sq.Queryer
		sq.QueryerContext
		QueryRower
		QueryRowerContext
	}

	baseConnection struct {
		_db      DBProxy
		_builder sq.StatementBuilderType
	}
)

// NewConnection will create a new instance of the `*BaseConnection`.
func NewConnection(db DBProxy, builder sq.StatementBuilderType) Connection {
	return &baseConnection{
		_db:      db,
		_builder: builder,
	}
}

// DB returns the real connection object for the database connection.
func (conn *baseConnection) DB() DBProxy {
	return conn._db
}

// Builder returns the Statement Builder used to generate the queries for this connection.
func (conn *baseConnection) Builder() sq.StatementBuilderType {
	return conn._builder.RunWith(conn)
}

// Begin starts a transaction.
func (conn *baseConnection) Begin() (TxProxy, error) {
	tx, err := conn._db.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

// BeginTx starts a transaction with more options.
func (conn *baseConnection) BeginTx(ctx context.Context, opts *sql.TxOptions) (TxProxy, error) {
	tx, err := conn._db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (conn *baseConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
	return conn._db.Exec(query, args...)
}

func (conn *baseConnection) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return conn._db.ExecContext(ctx, query, args...)
}

func (conn *baseConnection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return conn._db.Query(query, args...)
}

func (conn *baseConnection) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return conn._db.QueryContext(ctx, query, args...)
}

func (conn *baseConnection) QueryRow(query string, args ...interface{}) *sql.Row {
	return conn._db.QueryRow(query, args...)
}

func (conn *baseConnection) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return conn._db.QueryRowContext(ctx, query, args...)
}

func (conn *baseConnection) Prepare(query string) (*sql.Stmt, error) {
	return conn._db.Prepare(query)
}
