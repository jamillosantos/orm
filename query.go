package orm

import (
	"bytes"
	"errors"

	sq "github.com/Masterminds/squirrel"
)

type JoinType int

var ErrInvalidJoinType = errors.New("invalid join type")

const (
	JoinNone JoinType = iota
	InnerJoin
	LeftJoin
	RightJoin
	FullJoin
)

type Query interface {
	sq.Sqlizer
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
	Join(joinType JoinType, schema Schema, conditions ...sq.Sqlizer)
	// Join adds a inner join to the Query.
	InnerJoin(schema Schema, conditions ...sq.Sqlizer)
	// Join adds a left join to the Query.
	LeftJoin(schema Schema, conditions ...sq.Sqlizer)
	// Join adds a right join to the Query.
	RightJoin(schema Schema, conditions ...sq.Sqlizer)
	// Join adds a full join to the Query.
	FullJoin(schema Schema, conditions ...sq.Sqlizer)
	// Where define the where condition for the Query.
	Where(conditions ...sq.Sqlizer)
	// Skip sets the skip option for the Query.
	Skip(skip int)
	// Limit defines the pagination for the Query.
	Limit(limit int)
	// GroupBy defines the GROUP BY clause for the Query.
	GroupBy(fields ...SchemaField)
	// GroupByHaving defines the GROUP BY with the HAVING clause for the Query.
	GroupByHaving(fields []SchemaField, conditions ...sq.Sqlizer)
	// OrderBy defines the ORDER BY for the Query.
	OrderBy(fields ...SchemaField)
}

type DefaultScoper interface {
	DefaultScope() error
}

type join struct {
	joinType   JoinType
	schema     Schema
	conditions []sq.Sqlizer
}

func newJoin(joinType JoinType, schema Schema, conditions ...sq.Sqlizer) *join {
	return &join{
		joinType, schema, conditions,
	}
}

func (j *join) ToSql() (string, []interface{}, error) {
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
			s, as, err := condition.ToSql()
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
	Conn            Connection
	selectFields    []SchemaField
	selectFieldsStr []string
	from            Schema
	joins           []*join
	where           []sq.Sqlizer
	groupBy         []SchemaField
	groupByHaving   []sq.Sqlizer
	orderBy         []SchemaField
	skip            int
	limit           int
}

func NewQuery(conn Connection, schema Schema) Query {
	return &baseQuery{
		Conn: conn,
		from: schema,
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
	query.from = schema
	query._dirty = true
}

func (query *baseQuery) Join(joinType JoinType, schema Schema, conditions ...sq.Sqlizer) {
	if query.joins == nil {
		query.joins = make([]*join, 0, 2)
	}
	query.joins = append(query.joins, newJoin(joinType, schema, conditions...))
	query._dirty = true
}

func (query *baseQuery) InnerJoin(schema Schema, conditions ...sq.Sqlizer) {
	query.Join(LeftJoin, schema, conditions...)
}

func (query *baseQuery) LeftJoin(schema Schema, conditions ...sq.Sqlizer) {
	query.Join(LeftJoin, schema, conditions...)
}

func (query *baseQuery) RightJoin(schema Schema, conditions ...sq.Sqlizer) {
	query.Join(RightJoin, schema, conditions...)
}

func (query *baseQuery) FullJoin(schema Schema, conditions ...sq.Sqlizer) {
	query.Join(FullJoin, schema, conditions...)
}

func (query *baseQuery) Where(conditions ...sq.Sqlizer) {
	if query.where == nil {
		query.where = make([]sq.Sqlizer, 0, 3)
	}
	query.where = append(query.where, conditions...)
	query._dirty = true
}

func (query *baseQuery) Skip(skip int) {
	if (query.skip == 0 && query.skip != 0) || (query.skip != 0 && skip == 0) {
		// If skip WAS NOT defined and now it is, the query should be regenerated.
		// OR
		// If skip WAS defined and now it isn't, the query should be regenerated.
		query._dirty = true
	}
	query.skip = skip
}

func (query *baseQuery) Limit(limit int) {
	if (query.limit == 0 && query.limit != 0) || (query.limit != 0 && limit == 0) {
		// If limit WAS NOT defined and now it is, the query should be regenerated.
		// OR
		// If limit WAS defined and now it isn't, the query should be regenerated.
		query._dirty = true
	}
	query.limit = limit
}

func (query *baseQuery) GroupBy(fields ...SchemaField) {
	query.groupBy = fields
	query._dirty = true
}

func (query *baseQuery) GroupByHaving(fields []SchemaField, conditions ...sq.Sqlizer) {
	query.groupBy = fields
	query.groupByHaving = conditions
	query._dirty = true
}

func (query *baseQuery) OrderBy(fields ...SchemaField) {
	query.orderBy = fields
	query._dirty = true
}

func (query *baseQuery) ToSql() (string, []interface{}, error) {
	var selectFields []string
	if query._dirty {
		selectFields = make([]string, len(query.selectFields))
		for i, field := range query.selectFields {
			selectFields[i] = field.String()
		}
		query.selectFieldsStr = selectFields
	} else {
		selectFields = query.selectFieldsStr
	}
	builder := query.Conn.Builder().Select(selectFields...).From(query.from.Alias())
	for _, join := range query.joins {
		sqlJoin, argsJoin, err := join.ToSql()
		if err != nil {
			return "", nil, err
		}
		var f func(join string, rest ...interface{}) sq.SelectBuilder
		switch join.joinType {
		case JoinNone:
			f = builder.Join
		case InnerJoin:
			f = func(p string, args ...interface{}) sq.SelectBuilder {
				return builder.JoinClause("INNER JOIN "+p, args)
			}
		case LeftJoin:
			f = builder.LeftJoin
		case RightJoin:
			f = builder.RightJoin
		case FullJoin:
			f = func(p string, args ...interface{}) sq.SelectBuilder {
				return builder.JoinClause("FULL JOIN "+p, args)
			}
		default:
			return "", nil, ErrInvalidJoinType
		}
		builder = f(sqlJoin, argsJoin...)
	}
	return builder.ToSql()
}
