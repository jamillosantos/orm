package generator

import (
	"io"

	"github.com/setare/orm/gen/templates"
)

type ResultSetGenerator struct {
}

func (*ResultSetGenerator) Name() string {
	return "ResultSet"
}

func (*ResultSetGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteResultSet(writer, &templates.ResultSetInput{
		&gctx.OutputPackage,
		&gctx.ModelsPackage,
		gctx.Document.Records,
	})
	return nil
}
