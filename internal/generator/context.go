package generator

import "github.com/setare/orm/internal/document"

type Context struct {
	ModelsPackage document.Package
	Document      *document.Document
}
