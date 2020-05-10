package templates

import "github.com/setare/orm/internal/document"

type ResultSetInput struct {
	Package       *document.Package
	ModelsPackage *document.Package
	Records       []*document.Record
}
