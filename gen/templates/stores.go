package templates

import "github.com/setare/orm/gen/document"

type StoresInput struct {
	Package       *document.Package
	ModelsPackage *document.Package
	Records       []*document.Record
}
