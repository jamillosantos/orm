package orm

import "github.com/jackc/pgproto3/v2"

type ResultSet interface {
	Next() bool
	Scan(args ...interface{}) error
	Err() error
}

type ResultSetPgx interface {
	ResultSet
	Close()
	FieldDescriptions() []pgproto3.FieldDescription
}
