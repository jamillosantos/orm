package orm

import (
	"bytes"
	"errors"
	"strings"

	"github.com/jamillosantos/sqlf"
)

type JoinType int

var ErrInvalidJoinType = errors.New("invalid join type")

const (
	JoinNone JoinType = iota
	InnerJoin
	LeftJoin
	RightJoin
	FullJoin
	OuterJoin
)

func (t JoinType) String() string {
	switch t {
	case InnerJoin:
		return "INNER"
	case LeftJoin:
		return "LEFT"
	case RightJoin:
		return "RIGHT"
	case OuterJoin:
		return "OUTER"
	default:
		return ""
	}
}

type Query interface {
	sqlf.FastSqlizer
	sqlf.Sqlizer

	// Select defines what are the fields that this query should return. If this method is called twice, the second time
	// will override all values from the first call. If you want to stack fields, use `AddSelect` instead.
	Select(fields ...SchemaField)
	// AddSelect adds the passed fields to the select list.
	AddSelect(fields ...SchemaField)
	// GetSelect returns the fields that where selected.
	GetSelect() []SchemaField
	// From defines the FROM clause for the Query
	From(schema Schema)
	// Join adds a join to the Query.
	Join(joinType JoinType, schema Schema, condition string, params ...interface{})
	// Join adds a inner join to the Query.
	InnerJoin(schema Schema, condition string, params ...interface{})
	// Join adds a left join to the Query.
	LeftJoin(schema Schema, condition string, params ...interface{})
	// Join adds a right join to the Query.
	RightJoin(schema Schema, condition string, params ...interface{})
	// Join adds a full join to the Query.
	FullJoin(schema Schema, condition string, params ...interface{})
	// Where define the where condition for the Query.
	Where(condition string, params ...interface{})
	// WhereCriteria define the where condition for the Query using criteria.
	WhereCriteria(conditions ...sqlf.FastSqlizer)
	// Skip sets the skip option for the Query.
	Skip(skip int)
	// Limit defines the pagination for the Query.
	Limit(limit int)
	// GroupBy defines the GROUP BY clause for the Query.
	GroupBy(fields ...interface{})
	// GroupByX defines the GROUP BY with the HAVING clause for the Query.
	GroupByX(func(sqlf.GroupBy))
	// OrderBy defines the ORDER BY for the Query.
	OrderBy(fields ...interface{})
	// OrderBy defines the ORDER BY for the Query.
	OrderByX(f func(sqlf.OrderBy))
}

type DefaultScoper interface {
	DefaultScope() error
}

type join struct {
	joinType   JoinType
	schema     Schema
	conditions []sqlf.Sqlizer
}

func newJoin(joinType JoinType, schema Schema, conditions ...sqlf.Sqlizer) *join {
	return &join{
		joinType, schema, conditions,
	}
}

func (j *join) ToSQL() (string, []interface{}, error) {
	buf := bytes.NewBuffer(nil)
	a := j.schema.Alias()
	buf.WriteString(j.schema.Table())
	if a != "" {
		buf.WriteString(" ")
		buf.WriteString(j.schema.Alias())
	}
	if len(j.conditions) > 0 {
		buf.WriteString(" ON ")
		args := make([]interface{}, 0, len(j.conditions))
		for i, condition := range j.conditions {
			if i > 0 {
				buf.WriteString(" AND ")
			}
			s, as, err := condition.ToSQL()
			if err != nil {
				return "", nil, err
			}
			if len(as) > 0 {
				args = append(args, as...)
			}
			buf.WriteString(s)
		}
		return buf.String(), args, nil
	}
	return buf.String(), nil, nil
}

type baseQuery struct {
	_dirty          bool
	sqlQuery        sqlf.Select
	Conn            ConnectionPgx
	selectFields    []SchemaField
	selectFieldsStr []interface{}
	from            Schema
	groupBy         []SchemaField
	groupByHaving   []sqlf.Sqlizer
	orderBy         []SchemaField
	skip            int
	limit           int
}

func NewQuery(conn ConnectionPgx, schema Schema) Query {
	return &baseQuery{
		Conn:     conn,
		from:     schema,
		sqlQuery: conn.Builder().Select().From(schema.Table(), schema.Alias()),
	}
}

func (query *baseQuery) Select(fields ...SchemaField) {
	query.selectFields = fields
	query._dirty = true
}

func (query *baseQuery) GetSelect() []SchemaField {
	return query.selectFields
}

func (query *baseQuery) AddSelect(fields ...SchemaField) {
	if query.selectFields == nil {
		query.selectFields = fields
		return
	}
	query.selectFields = append(query.selectFields, fields...)
	query._dirty = true
}

func (query *baseQuery) From(schema Schema) {
	query.sqlQuery.From(schema.Table(), schema.Alias())
}

func (query *baseQuery) Join(joinType JoinType, schema Schema, condition string, params ...interface{}) {
	query.sqlQuery.JoinClause(joinType.String(), schema.Table(), schema.Alias()).On(condition, params)
}

func (query *baseQuery) InnerJoin(schema Schema, condition string, params ...interface{}) {
	query.sqlQuery.JoinClause(InnerJoin.String(), schema.Table(), schema.Alias()).On(condition, params)
}

func (query *baseQuery) LeftJoin(schema Schema, condition string, params ...interface{}) {
	query.sqlQuery.JoinClause(LeftJoin.String(), schema.Table(), schema.Alias()).On(condition, params)
}

func (query *baseQuery) RightJoin(schema Schema, condition string, params ...interface{}) {
	query.sqlQuery.JoinClause(RightJoin.String(), schema.Table(), schema.Alias()).On(condition, params)
}

func (query *baseQuery) FullJoin(schema Schema, condition string, params ...interface{}) {
	query.sqlQuery.JoinClause(FullJoin.String(), schema.Table(), schema.Alias()).On(condition, params)
}

func (query *baseQuery) Where(condition string, params ...interface{}) {
	query.sqlQuery.Where(condition, params...)
}

func (query *baseQuery) WhereCriteria(criteria ...sqlf.FastSqlizer) {
	query.sqlQuery.WhereCriteria(criteria...)
}

func (query *baseQuery) Skip(skip int) {
	query.sqlQuery.Offset(skip)
}

func (query *baseQuery) Limit(limit int) {
	query.sqlQuery.Limit(limit)
}

func (query *baseQuery) GroupBy(fields ...interface{}) {
	query.sqlQuery.GroupBy(fields...)
}

func (query *baseQuery) GroupByX(f func(groupBy sqlf.GroupBy)) {
	query.sqlQuery.GroupByX(f)
}

func (query *baseQuery) OrderBy(fields ...interface{}) {
	query.sqlQuery.OrderBy(fields...)
}

func (query *baseQuery) OrderByX(f func(orderBy sqlf.OrderBy)) {
	query.sqlQuery.OrderByX(f)
}

func (query *baseQuery) ToSQL() (string, []interface{}, error) {
	args := make([]interface{}, 0, 2)
	var sb strings.Builder
	err := query.ToSQLFast(&sb, &args)
	if err != nil {
		return "", nil, err
	}
	return sb.String(), args, nil
}

func (query *baseQuery) ToSQLFast(sb *strings.Builder, args *[]interface{}) error {
	var selectFields []interface{}
	if query._dirty {
		selectFields = make([]interface{}, len(query.selectFields))
		for i, field := range query.selectFields {
			selectFields[i] = field.String()
		}
		query.selectFieldsStr = selectFields
	} else {
		selectFields = query.selectFieldsStr
	}
	builder := query.sqlQuery.Select(selectFields...)
	return builder.ToSQLFast(sb, args)
}
