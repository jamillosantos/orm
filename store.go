package orm

type Store interface {
	Insert() error
	InsertBatch() error
	Update() error
}

type baseStore struct {
	//
}
