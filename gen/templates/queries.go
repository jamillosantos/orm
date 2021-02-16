package templates

import "github.com/jamillosantos/orm/gen/document"

type QueriesInput struct {
	Package       *document.Package
	ModelsPackage *document.Package
	Records       []*document.Record
}
