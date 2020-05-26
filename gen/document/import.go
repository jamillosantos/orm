package document

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

type Import struct {
	Name       string
	ImportPath string
}

func (i *Import) String() string {
	if i.Name != "" {
		return fmt.Sprintf(`%s "%s"`, i.Name, i.ImportPath)
	}
	return fmt.Sprintf(`"%s"`, i.ImportPath)
}

var (
	importRegex = regexp.MustCompile("^([a-z][a-zA-Z0-9_]+) ")
)

// UnmarshalYAML
func (i *Import) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		if value.Tag == "!!str" {
			m := importRegex.FindStringSubmatch(value.Value)
			if m == nil {
				i.Name = ""
				i.ImportPath = value.Value
				return nil
			}
			i.Name = m[1]
			i.ImportPath = m[2]
		} else {
			// TODO(Jota): Add the line information on this error.
			return ErrInvalidField
		}
	default:
		return ErrInvalidField
	}
	return nil
}
