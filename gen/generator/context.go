package generator

import "github.com/jamillosantos/orm/gen/document"

type Context struct {
	ModelsPackage document.Package
	OutputPackage document.Package
	Document      *document.Document
}
