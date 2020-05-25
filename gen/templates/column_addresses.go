package templates

import (
	"github.com/setare/orm/gen/document"
)

type ColumnAddressesInput struct {
	FieldName  string
	TargetName string
	RecordName string
	ErrName    string
	Record     *document.Record
}