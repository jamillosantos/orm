package templates

import "github.com/setare/orm/internal/document"

type SchemaInput struct {
	Package *document.Package
	Records []*document.Record
}
