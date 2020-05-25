package document

type GoType interface {
	// Package returns the original go package that this type came from.
	Package() string

	// Name is the go type name.
	Name() string
}

type baseGoType struct {
	pkg  string
	name string
}

// NewGoType returns a new instance of a `GoType` interface with the package and name initialized.
//
// The base implementation is a simple struct that keep only the required information for implementing
// `GoType`.
func NewGoType(pkg, name string) GoType {
	return &baseGoType{
		pkg,
		name,
	}
}

// Package returns the original go package that this type came from.
func (t *baseGoType) Package() string {
	return t.pkg
}

// Name is the go type name.
func (t *baseGoType) Name() string {
	return t.name
}
