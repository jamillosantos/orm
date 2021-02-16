package generator

import (
	"io"

	"github.com/jamillosantos/orm/gen/templates"
)

type ConnectionsGenerator struct {
}

func (*ConnectionsGenerator) Name() string {
	return "Connections"
}

func (*ConnectionsGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteConnections(writer, &templates.ConnectionsInput{
		&gctx.OutputPackage,
		gctx.Document.Records,
	})
	return nil
}
