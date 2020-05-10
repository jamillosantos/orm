package document

type Generators struct {
	Models      bool `yaml:"models"`
	Schema      bool `yaml:"schema"`
	Connections bool `yaml:"connections"`
	ResultSets  bool `yaml:"resultsets"`
	Stores      bool `yaml:"stores"`
	Queries     bool `yaml:"queries"`
}
