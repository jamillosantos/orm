package orm

type Query interface {
	Select(fields ...string) Query
}
