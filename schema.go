package orm

import "fmt"

// Schema represents a table schema.
type Schema interface {
	// Table returns the table name for this schema.
	Table() string
	// Alias returns the alias name for the table. By default, its implementation
	// should return the table name.
	Alias() string
	// As returns a new Schema with an SQL alias set.
	As(alias string) Schema
	// Columns list all columns from this schema.
	Columns() []SchemaField
}

// SchemaField represents a field in the schema.
type SchemaField interface {
	fmt.Stringer
	// Name returns the name of the field
	Name() string
	// QualifiedName returns the field name together with the Schema table name.
	QualifiedName(Schema) string
}

type baseSchema struct {
	tableName string
	columnArr []SchemaField
}

type aliasSchema struct {
	schema Schema
	alias  string
}

func (schema *baseSchema) Table() string {
	return schema.tableName
}

func (schema *baseSchema) Alias() string {
	return schema.tableName
}

func (schema *baseSchema) As(alias string) Schema {
	return &aliasSchema{
		schema: schema,
		alias:  alias,
	}
}

func (schema *baseSchema) Columns() []SchemaField {
	return schema.columnArr
}

func (schema *aliasSchema) Table() string {
	return schema.schema.Table()
}

func (schema *aliasSchema) Alias() string {
	return schema.alias
}

func (schema *aliasSchema) As(alias string) Schema {
	return &aliasSchema{
		schema: schema,
		alias:  alias,
	}
}

func (schema *aliasSchema) Columns() []SchemaField {
	return schema.schema.Columns()
}
