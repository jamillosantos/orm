package document

// Package will keep the information about some packages that will be needed for
// generation.
type Package struct {
	Name       string
	ImportPath string
	Directory  string
}
