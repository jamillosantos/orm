package orm

import "github.com/jackc/pgproto3/v2"

type ResultSetPgx interface {
	Rows
	FieldDescriptions() []pgproto3.FieldDescription
}
