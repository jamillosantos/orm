package templates

import "github.com/setare/orm/internal/document"

type ConnectionsInput struct {
	Package *document.Package
	Records []*document.Record
}
