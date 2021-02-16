package templates

import "github.com/jamillosantos/orm/gen/document"

type ConnectionsInput struct {
	Package *document.Package
	Records []*document.Record
}
