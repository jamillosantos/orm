package generator

import "io"

type Generator interface {
	Name() string
	Generate(writer io.Writer, gctx *Context) error
}
