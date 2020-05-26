package generator

import (
	"io"

	"github.com/setare/orm/gen/templates"
)

type ModelGenerator struct {
}

func (*ModelGenerator) Name() string {
	return "Models"
}

func (*ModelGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteRecords(writer, gctx.ModelsPackage, gctx.Document.Records, gctx.Document.Imports)
	return nil
}
