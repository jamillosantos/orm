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
	Version    string     `yaml:"version"`
	Output     Output     `yaml:"output"`
	Generators Generators `yaml:"generators"`
	Imports    []*Import  `yaml:"imports"`
	Records    []*Record  `yaml:"records"`
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
			i++
			doc.Imports = make([]*Import, len(value.Content[i].Content))
			for i := 0; i < len(doc.Imports); i++ {
				doc.Imports[i] = &Import{}
			}
			err := value.Content[i].Decode(&doc.Imports)
			if err != nil {
				return err
			}
		case "output":
			i++
			err := value.Content[i].Decode(&doc.Output)
			if err != nil {
				return err
			}
		case "generators":
			i++
			err := value.Content[i].Decode(&doc.Generators)
			if err != nil {
				return err
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
