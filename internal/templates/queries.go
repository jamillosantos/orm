package templates

import "github.com/setare/orm/internal/document"

type QueriesInput struct {
	Package       *document.Package
	ModelsPackage *document.Package
	Records       []*document.Record
}
