package orm

type ResultSet struct {
	Close() error
	Columns() ([]string, error)
	Next() bool
	NextResultSet() bool
	Scan(args...interface{}) error
	Err() error
}