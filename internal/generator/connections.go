package generator

import (
	"io"

	"github.com/setare/orm/internal/templates"
)

type ConnectionsGenerator struct {
}

func (*ConnectionsGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteConnections(writer, &templates.ConnectionsInput{
		&gctx.ModelsPackage,
		gctx.Document.Records,
	})
	return nil
}
