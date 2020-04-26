package query

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type condition struct {
	a  interface{}
	op string
	b  interface{}
}

// resolveConditionPart will return the SQL
func resolveConditionPart(part interface{}) (string, []interface{}, error) {
	switch p := part.(type) {
	case sq.Sqlizer:
		return p.ToSql()
	default:
		return "?", []interface{}{p}, nil
	}
}

// ToSql
func (cond *condition) ToSql() (string, []interface{}, error) {
	var (
		sqlA, sqlB   string
		argsA, argsB []interface{}
		err          error
	)

	sqlA, argsA, err = resolveConditionPart(cond.a)
	if err != nil {
		return "", nil, err
	}
	sqlB, argsB, err = resolveConditionPart(cond.b)
	if err != nil {
		return "", nil, err
	}

	var args []interface{}

	if argsA != nil && argsB != nil {
		args = append(argsA, argsB...)
	} else if argsA != nil {
		args = argsA
	} else if argsB != nil {
		args = argsB
	}

	return fmt.Sprintf("%s %s %s", sqlA, cond.op, sqlB), args, nil
}

func Eq(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, "=", b,
	}
}

func NotEq(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, "<>", b,
	}
}

func Like(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, "LIKE", b,
	}
}

func In(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, "IN", b,
	}
}

func GT(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, ">", b,
	}
}

func GTE(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, ">=", b,
	}
}

func LT(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, "<", b,
	}
}

func LTE(a, b interface{}) sq.Sqlizer {
	return &condition{
		a, "<=", b,
	}
}

type conditionBetween struct {
	a,
	b,
	c interface{}
}

func (cond *conditionBetween) ToSql() (string, []interface{}, error) {
	sqlA, argsA, err := resolveConditionPart(cond.a)
	if err != nil {
		return "", nil, err
	}
	sqlB, argsB, err := resolveConditionPart(cond.b)
	if err != nil {
		return "", nil, err
	}
	sqlC, argsC, err := resolveConditionPart(cond.c)
	if err != nil {
		return "", nil, err
	}
	var args []interface{}
	if argsA != nil || argsB != nil || argsC != nil {
		args = make([]interface{}, 0, len(argsA)+len(argsB)+len(argsC))
	}
	if argsA != nil {
		args = append(args, argsA...)
	}
	if argsB != nil {
		args = append(args, argsB...)

	}
	if argsC != nil {
		args = append(args, argsC...)
	}
	return fmt.Sprintf("%s BETWEEN %s AND %s", sqlA, sqlB, sqlC), args, nil
}

func Between(a, b, c interface{}) sq.Sqlizer {
	return &conditionBetween{
		a, b, c,
	}
}
