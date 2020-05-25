package templates

import "github.com/setare/orm/gen/document"

type SchemaInput struct {
	Package *document.Package
	Records []*document.Record
}
