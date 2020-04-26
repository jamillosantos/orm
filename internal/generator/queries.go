package generator

import (
	"io"

	"github.com/setare/orm/internal/templates"
)

type QueriesGenerator struct {
}

func (*QueriesGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteQueries(writer, &templates.QueriesInput{
		&gctx.ModelsPackage,
		gctx.Document.Records,
	})
	return nil
}
