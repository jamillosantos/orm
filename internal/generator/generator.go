package generator

import "io"

type Generator interface {
	Generate(writer io.Writer, gctx *Context) error
}
