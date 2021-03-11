package orm

import (
	"context"
	"database/sql"

	"github.com/jamillosantos/sqlf"
)

// SQLConnection is the `Connection` interface implemented using the standard
// "sql" package.
type SQLConnection struct {
	_db      *sql.DB
	_builder sqlf.Builder
}

// NewConnectionPgx will create a new instance of the `*BaseConnection`.
func NewConnectionSQL(db *sql.DB, builder sqlf.Builder) Connection {
	return &SQLConnection{
		_db:      db,
		_builder: builder,
	}
}

// SQLTx is the concrete implementation of `Connection` for sql standard
// implementation.
type SQLTx struct {
	tx *sql.Tx
}

// NewSQLTx returns a new transaction abstraction.
func NewSQLTx(tx *sql.Tx) TxProxy {
	return &SQLTx{
		tx: tx,
	}
}

// DB returns the real connection object for the database connection.
func (conn *SQLConnection) DB() *sql.DB {
	return conn._db
}

// Builder returns the Statement Builder used to generate the queries for this connection.
func (conn *SQLConnection) Builder() sqlf.Builder {
	return conn._builder
}

// Begin starts a transaction.
func (conn *SQLConnection) Begin(ctx context.Context) (TxProxy, error) {
	tx, err := conn._db.Begin()
	if err != nil {
		return nil, err
	}
	return NewSQLTx(tx), nil
}

func (conn *SQLConnection) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return conn._db.ExecContext(ctx, query, args...)
}

func (conn *SQLConnection) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	return conn._db.QueryContext(ctx, sql, args...)
}

func (conn *SQLConnection) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return conn._db.QueryRowContext(ctx, query, args...)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tx *SQLTx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, query, args...)
}

func (tx *SQLTx) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	return tx.tx.QueryContext(ctx, sql, args...)
}

func (tx *SQLTx) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return tx.tx.QueryRowContext(ctx, query, args...)
}

func (tx *SQLTx) Commit(ctx context.Context) error {
	return tx.tx.Commit()
}

func (tx *SQLTx) Rollback(ctx context.Context) error {
	return tx.tx.Rollback()
}
