package generator

import (
	"io"

	"github.com/setare/orm/internal/templates"
)

type SchemaGenerator struct {
}

func (*SchemaGenerator) Name() string {
	return "Schema"
}

func (*SchemaGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteSchema(writer, &templates.SchemaInput{
		&gctx.OutputPackage,
		gctx.Document.Records,
	})
	return nil
}
