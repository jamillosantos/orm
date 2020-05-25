package generator

import (
	"io"

	"github.com/setare/orm/gen/templates"
)

type StoresGenerator struct {
}

func (*StoresGenerator) Name() string {
	return "Stores"
}

func (*StoresGenerator) Generate(writer io.Writer, gctx *Context) error {
	templates.WriteStores(writer, &templates.StoresInput{
		&gctx.OutputPackage,
		&gctx.ModelsPackage,
		gctx.Document.Records,
	})
	return nil
}
