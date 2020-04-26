package query

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func And(conditions ...sq.Sqlizer) sq.Sqlizer {
	return sq.And(conditions)
}

func Or(conditions ...sq.Sqlizer) sq.Sqlizer {
	return sq.Or(conditions)
}

type conditionFormat struct {
	format string
	cond   sq.Sqlizer
}

func (cond *conditionFormat) ToSql() (string, []interface{}, error) {
	sql, args, err := cond.cond.ToSql()
	if err != nil {
		return "", args, err
	}
	return fmt.Sprintf(cond.format, sql), args, err
}

func Not(condition sq.Sqlizer) sq.Sqlizer {
	return &conditionFormat{"NOT %s", condition}
}

func Some(condition sq.Sqlizer) sq.Sqlizer {
	return &conditionFormat{"SOME %s", condition}
}

func Exists(condition sq.Sqlizer) sq.Sqlizer {
	return &conditionFormat{"EXISTS %s", condition}
}

func Any(condition sq.Sqlizer) sq.Sqlizer {
	return &conditionFormat{"ANY %s", condition}
}

func All(condition sq.Sqlizer) sq.Sqlizer {
	return &conditionFormat{"ALL %s", condition}
}

func Using(condition sq.Sqlizer) sq.Sqlizer {
	return &conditionFormat{"USING (%s)", condition}
}
