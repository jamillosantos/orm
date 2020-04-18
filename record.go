package orm

type Record interface {
	GetID() interface{}
	TableName() string
	IsPersisted() bool
	IsWritable() bool
}

// BaseRecord implements some basic methods to be used by the store.
type BaseRecord struct {
	tableName string
	persisted bool
	writable  bool
}

// TableName returns the name of the table that this record repesents.
func (record *BaseRecord) TableName() string {
	return record.tableName
}

// IsPersisted returns if the record was persisted.
func (record *BaseRecord) IsPersisted() bool {
	return record.persisted
}

// IsWritable returns if the record can be saved to the database.
func (record *BaseRecord) IsWritable() bool {
	return record.writable
}
