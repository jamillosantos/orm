package query

import (
	sq "github.com/Masterminds/squirrel"
)

type rawSql struct {
	sql  string
	args []interface{}
}

func Raw(sql string, args ...interface{}) sq.Sqlizer {
	return &rawSql{sql, args}
}

func (r *rawSql) ToSql() (string, []interface{}, error) {
	return r.sql, r.args, nil
}

func RawSql(sql string, args ...interface{}) sq.Sqlizer {
	return &rawSql{sql, args}
}
