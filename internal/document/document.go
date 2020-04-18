package document

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var (
	ErrNotDocument     = errors.New("not document")
	ErrUnknownProperty = errors.New("unknown property")
)

type Document struct {
	Version string    `yaml:"version"`
	Imports []string  `yaml:"imports"`
	Records []*Record `yaml:"records"`
}

func (doc *Document) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return ErrNotDocument
	}
	for i := 0; i < len(value.Content); i++ {
		content := value.Content[i]
		switch content.Value {
		case "version":
			i++
			doc.Version = value.Content[i].Value
		case "imports":
			doc.Imports = make([]string, 0)
			i++
			for _, imp := range value.Content[i].Content {
				doc.Imports = append(doc.Imports, imp.Value)
			}
		case "records":
			i++
			doc.Records = make([]*Record, len(value.Content[i].Content))
			for i := 0; i < len(doc.Records); i++ {
				doc.Records[i] = &Record{
					Document: doc,
				}
			}
			err := value.Content[i].Decode(&doc.Records)
			if err != nil {
				return err
			}
		default:
			// TODO(Jota): Add the line information on this error.
			return errors.Wrap(ErrUnknownProperty, content.Value)
		}
	}
	return nil
}
