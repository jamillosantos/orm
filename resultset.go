package orm

type ResultSet interface {
	Close() error
	Columns() ([]string, error)
	Next() bool
	NextResultSet() bool
	Scan(args ...interface{}) error
	Err() error
}
