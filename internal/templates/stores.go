package templates

import "github.com/setare/orm/internal/document"

type StoresInput struct {
	Package *document.Package
	Records []*document.Record
}
