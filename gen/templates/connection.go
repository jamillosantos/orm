package templates

import "github.com/setare/orm/gen/document"

type ConnectionsInput struct {
	Package *document.Package
	Records []*document.Record
}
