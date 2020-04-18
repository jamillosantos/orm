package orm

type Store interface {
	Insert() error
}

type BaseStore struct {
	//
}
