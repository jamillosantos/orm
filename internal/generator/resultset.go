package generator

import (
	"io"

	"github.com/setare/orm/internal/templates"
)

type ResultSetGenerator struct {
}

func (*ResultSetGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteResultSet(writer, &templates.ResultSetInput{
		&gctx.ModelsPackage,
		gctx.Document.Records,
	})
	return nil
}
