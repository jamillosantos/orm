package templates

import "github.com/jamillosantos/orm/gen/document"

type SchemaInput struct {
	Package *document.Package
	Records []*document.Record
}
