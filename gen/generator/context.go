package generator

import "github.com/setare/orm/gen/document"

type Context struct {
	ModelsPackage document.Package
	OutputPackage document.Package
	Document      *document.Document
}
