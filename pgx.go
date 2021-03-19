package orm

import (
	"context"
	"database/sql"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jamillosantos/sqlf"
)

// PgxConnection is the `Connection` abstraction implementing the pgx library.
type PgxConnection struct {
	_pool    *pgxpool.Pool
	_builder sqlf.Builder
}

// NewPgxConnection will create a new instance of the `*BaseConnection`.
func NewPgxConnection(db *pgxpool.Pool, builder sqlf.Builder) Connection {
	return &PgxConnection{
		_pool:    db,
		_builder: builder,
	}
}

// PgxTx is the concrete implementation of a `Tx`.
type PgxTx struct {
	tx pgx.Tx
}

// NewPgxTx returns a new transaction abstraction.
func NewPgxTx(tx pgx.Tx) TxProxy {
	return &PgxTx{
		tx: tx,
	}
}

type PgxRows struct {
	pgx.Rows
	columns []string
}

func newPgxRows(rs pgx.Rows, err error) (Rows, error) {
	if err != nil {
		return nil, err
	}

	fd := rs.FieldDescriptions()
	columns := make([]string, len(fd))
	for i, field := range fd {
		columns[i] = string(field.Name)
	}
	return &PgxRows{
		Rows:    rs,
		columns: columns,
	}, nil
}

// Builder returns the Statement Builder used to generate the queries for this connection.
func (conn *PgxConnection) Builder() sqlf.Builder {
	return conn._builder
}

// Begin starts a transaction.
func (conn *PgxConnection) Begin(ctx context.Context) (TxProxy, error) {
	tx, err := conn._pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return NewPgxTx(tx), nil
}

// BeginTx starts a transaction with more options.
func (conn *PgxConnection) BeginTx(ctx context.Context, opts pgx.TxOptions) (TxProxy, error) {
	tx, err := conn._pool.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewPgxTx(tx), nil
}

func (conn *PgxConnection) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	_, err := conn._pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	// TODO(Jota): To check how is this going to work.
	return nil, nil
}

func (conn *PgxConnection) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	return newPgxRows(conn._pool.Query(ctx, sql, args...))
}

func (conn *PgxConnection) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return conn._pool.QueryFunc(ctx, sql, args, scans, f)
}

func (conn *PgxConnection) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return conn._pool.QueryRow(ctx, query, args...)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tx *PgxTx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	_, err := tx.tx.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	// TODO(Jota): Ensure what is happening here.
	return nil, nil
}

func (tx *PgxTx) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	return newPgxRows(tx.tx.Query(ctx, sql, args...))
}

func (tx *PgxTx) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return tx.tx.QueryFunc(ctx, sql, args, scans, f)
}

func (tx *PgxTx) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return tx.tx.QueryRow(ctx, query, args...)
}

func (tx *PgxTx) Prepare(ctx context.Context, name, query string) (*pgconn.StatementDescription, error) {
	return tx.tx.Prepare(ctx, name, query)
}

func (tx *PgxTx) Commit(ctx context.Context) error {
	return tx.tx.Commit(ctx)
}

func (tx *PgxTx) Rollback(ctx context.Context) error {
	return tx.tx.Rollback(ctx)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (rows *PgxRows) Scan(dest ...interface{}) error {
	return rows.Rows.Scan(dest...)
}

func (rows *PgxRows) Next() bool {
	return rows.Rows.Next()
}

func (rows *PgxRows) Columns() ([]string, error) {
	return rows.columns, nil
}

func (rows *PgxRows) Close() error {
	rows.Rows.Close()
	return nil
}
