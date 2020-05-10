package document

func NewDocument() *Document {
	doc := &Document{
		Output: Output{
			Model:      "models/models_gen.go",
			Directory:  "db",
			Connection: "connections_gen.go",
			Query:      "queries_gen.go",
			ResultSet:  "resultset_gen.go",
			Schema:     "schema_gen.go",
			Store:      "store_gen.go",
		},
		Generators: Generators{
			Models:      true,
			Schema:      true,
			Connections: true,
			ResultSets:  true,
			Stores:      true,
			Queries:     true,
		},
	}
	return doc
}
