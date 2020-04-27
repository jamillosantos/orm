package document

import (
	"strconv"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var (
	ErrInvalidField        = errors.New("invalid field")
	ErrOnlyOneAutoIncField = errors.New("only one autoinc field is allowed")

	plrzClient = pluralize.NewClient()
)

// Record represents a record.
type Record struct {
	Document            *Document  `yaml:"-"`
	Schema              *Schema    `yaml:"-"`
	Query               *Query     `yaml:"-"`
	Store               *Store     `yaml:"-"`
	ResultSet           *ResultSet `yaml:"-"`
	Name                string     `yaml:"name"`
	TableName           string     `yaml:"table_name"`
	Documentation       []string   `yaml:"-"`
	Fields              []*Field   `yaml:"fields"`
	PrimaryKey          []*Field   `yaml:"-"`
	FieldAutoInc        *Field     `yaml:"-"`
	FieldsNameMaxLength int        `yaml:"-"`
	FieldsTypeMaxLength int        `yaml:"-"`
}

// UnmarshalYAML
func (record *Record) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return ErrNotDocument
	}
	if len(value.Content) > 0 && value.Content[0].HeadComment != "" {
		docs := strings.Split(value.Content[0].HeadComment, "\n")
		record.Documentation = make([]string, 0, len(docs))
		for _, d := range docs {
			if d[0] == '#' {
				d = d[1:]
			}
			if len(d) == 0 {
				record.Documentation = append(record.Documentation, "")
				continue
			}
			if d[0] == ' ' {
				d = d[1:]
			}
			record.Documentation = append(record.Documentation, d)
		}
	}
	for i := 0; i < len(value.Content); i++ {
		content := value.Content[i]
		switch content.Value {
		case "name":
			i++
			record.Name = value.Content[i].Value
			if record.TableName == "" {
				record.TableName = plrzClient.Plural(record.Name)
			}
		case "table_name":
			i++
			record.TableName = value.Content[i].Value
		case "fields":
			i++
			content := value.Content[i]
			record.Fields = make([]*Field, len(content.Content))
			for i := 0; i < len(record.Fields); i++ {
				record.Fields[i] = &Field{
					Record: record,
				}
				err := content.Content[i].Decode(record.Fields[i])
				if err != nil {
					return err
				}
			}
		default:
			// TODO(Jota): Add the line information on this error.
			return errors.Wrap(ErrUnknownProperty, content.Value)
		}
	}
	record.Schema = &Schema{
		Name:        record.Name,
		Type:        "schema" + record.Name,
		InternalRef: "defaultSchema" + record.Name,
	}
	record.Query = &Query{
		Type: record.Name + "Query",
	}
	record.Store = &Store{
		Type: record.Name + "Store",
	}
	record.ResultSet = &ResultSet{
		Type: record.Name + "ResultSet",
	}
	return nil
}

// Field represents a field in the yaml file.
type Field struct {
	Record     *Record `yaml:"-"`
	GoName     string  `yaml:"go_name"`
	Name       string  `yaml:"name"`
	Type       string  `yaml:"type"`
	AutoInc    bool    `yaml:"auto_inc"`
	PrimaryKey bool    `yaml:"pk"`
}

// UnmarshalYAML
func (field *Field) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.MappingNode:
		for i := 0; i < len(value.Content); i++ {
			content := value.Content[i]
			switch content.Value {
			case "name":
				i++
				field.Name = value.Content[i].Value
				if field.GoName == "" {
					field.GoName = GoNamePublic(field.Name)
					if len(field.GoName) > field.Record.FieldsNameMaxLength {
						field.Record.FieldsNameMaxLength = len(field.GoName)
					}
				}
			case "gname", "go_name":
				i++
				field.GoName = value.Content[i].Value
				if len(field.GoName) > field.Record.FieldsNameMaxLength {
					field.Record.FieldsNameMaxLength = len(field.GoName)
				}
			case "type":
				i++
				field.Type = value.Content[i].Value
				if len(field.Type) > field.Record.FieldsTypeMaxLength {
					field.Record.FieldsTypeMaxLength = len(field.Type)
				}
			case "autoinc":
				i++
				ai, err := strconv.ParseBool(value.Content[i].Value)
				if err != nil {
					return err // TODO(Jota): Wrap this with an proper error (with location too).
				}
				field.AutoInc = ai
				if ai {
					if field.Record.FieldAutoInc != nil {
						// TODO(Jota): Add the line information on this error.
						return errors.Wrap(ErrOnlyOneAutoIncField, field.Name)
					}
					field.Record.FieldAutoInc = field
				}
			case "pk":
				i++
				pk, err := strconv.ParseBool(value.Content[i].Value)
				if err != nil {
					return err // TODO(Jota): Wrap this with an proper error (with location too).
				}
				field.PrimaryKey = pk
				if pk {
					field.Record.PrimaryKey = append(field.Record.PrimaryKey, field)
				}
			default:
				// TODO(Jota): Add the line information on this error.
				return errors.Wrap(ErrUnknownProperty, content.Value)
			}
		}
	case yaml.ScalarNode:
		if value.Tag == "!!str" {
			toks := strings.Split(value.Value, ":")
			if len(toks) != 2 {
				// TODO(Jota): Add the line information on this error.
				return ErrInvalidField
			}
			field.Name = toks[0]
			field.GoName = GoNamePublic(toks[0])
			field.Type = toks[1]
		} else {
			// TODO(Jota): Add the line information on this error.
			return ErrInvalidField
		}
	}
	return nil
}
