package document

import (
	"fmt"
)

// Package will keep the information about some packages that will be needed for
// generation.
type Package struct {
	Name       string
	ImportPath string
	Directory  string
	Alias      string
}

func (pkg *Package) Equals(b *Package) bool {
	return pkg.ImportPath == b.ImportPath
}

func (pkg *Package) Ref(srcPkg *Package, t string) string {
	if !pkg.Equals(srcPkg) {
		return fmt.Sprintf("%s.%s", pkg.Name, t)
	}
	return t
}
