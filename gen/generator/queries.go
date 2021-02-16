package generator

import (
	"io"

	"github.com/jamillosantos/orm/gen/templates"
)

type QueriesGenerator struct {
}

func (*QueriesGenerator) Name() string {
	return "Queries"
}

func (*QueriesGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteQueries(writer, &templates.QueriesInput{
		&gctx.OutputPackage,
		&gctx.ModelsPackage,
		gctx.Document.Records,
	})
	return nil
}
