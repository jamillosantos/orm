package orm

import (
	"context"
	"database/sql"

	"github.com/jamillosantos/sqlf"
)

type (
	Row interface {
		Scan(dest ...interface{}) error
	}

	Rows interface {
		Row
		Next() bool
		Columns() ([]string, error)
		Err() error
		Close() error
	}

	TxProxy interface {
		DBRunner
		Commit(context.Context) error
		Rollback(context.Context) error
	}

	DBExecer interface {
		Exec(ctx context.Context, sql string, arguments ...interface{}) (sql.Result, error)
	}

	DBQueryer interface {
		Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	}

	DBRowQueryer interface {
		QueryRow(ctx context.Context, sql string, args ...interface{}) Row
	}

	DBRunner interface {
		DBExecer
		DBQueryer
		DBRowQueryer
	}

	Connection interface {
		DBRunner
		Builder() sqlf.Builder
		Begin(ctx context.Context) (TxProxy, error)
	}
)
