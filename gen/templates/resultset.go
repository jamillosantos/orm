package templates

import "github.com/jamillosantos/orm/gen/document"

type ResultSetInput struct {
	Package       *document.Package
	ModelsPackage *document.Package
	Records       []*document.Record
}
