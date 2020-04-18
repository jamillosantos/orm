package generator

import (
	"io"

	"github.com/setare/orm/internal/templates"
)

type ModelGenerator struct {
}

func (*ModelGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteRecords(writer, gctx.ModelsPackage, gctx.Document.Records)
	return nil
}
