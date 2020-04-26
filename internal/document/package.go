package document

import "fmt"

// Package will keep the information about some packages that will be needed for
// generation.
type Package struct {
	Name       string
	ImportPath string
	Directory  string
}

func (pkg *Package) Ref(srcPkg *Package, t string) string {
	if pkg.ImportPath != srcPkg.ImportPath {
		return fmt.Sprintf("%s.%s", pkg.Name, t)
	}
	return t
}
