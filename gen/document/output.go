package document

type Output struct {
	Package    string `yaml:"package"`
	Model      string `yaml:"model`
	Directory  string `yaml:"directory"`
	Schema     string `yaml:"schema"`
	Connection string `yaml:"connection"`
	ResultSet  string `yaml:"resultset"`
	Store      string `yaml:"store"`
	Query      string `yaml:"query"`
}
