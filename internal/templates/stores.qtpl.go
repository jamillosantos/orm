// Code generated by qtc from "stores.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line stores.qtpl:1
package templates

//line stores.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line stores.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line stores.qtpl:1
func StreamStores(qw422016 *qt422016.Writer, input *StoresInput) {
//line stores.qtpl:1
	qw422016.N().S(`
package `)
//line stores.qtpl:2
	qw422016.E().S(input.Package.Name)
//line stores.qtpl:2
	qw422016.N().S(`

import (
	"github.com/setare/orm"

	"context"
	"github.com/pkg/errors"
)

`)
//line stores.qtpl:11
	for _, record := range input.Records {
//line stores.qtpl:11
		qw422016.N().S(`type `)
//line stores.qtpl:12
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:12
		qw422016.N().S(` struct {
	conn orm.DBProxy
}

func New`)
//line stores.qtpl:16
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:16
		qw422016.N().S(`(conn orm.Connection) *`)
//line stores.qtpl:16
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:16
		qw422016.N().S(` {
	return &`)
//line stores.qtpl:17
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:17
		qw422016.N().S(`{
		conn,
	}
}

func (store *`)
//line stores.qtpl:22
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:22
		qw422016.N().S(`) Insert(record *`)
//line stores.qtpl:22
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:22
		qw422016.N().S(`, fields ...SchemaField) error {
	return store.InsertContext(context.Background(), record, fields...)
}

func (store *`)
//line stores.qtpl:26
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:26
		qw422016.N().S(`) InsertContext(ctx context.Context, record *`)
//line stores.qtpl:26
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:26
		qw422016.N().S(`, fields ...SchemaField) error {
	if len(fields) == 0 {
		fields = `)
//line stores.qtpl:28
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:28
		qw422016.N().S(`Fields
	}
	columnNames := make([]string, len(fields))
	columnValues := make([]interface{}, len(fields))
	var err error
	for i, field := range fields {
		var fieldAddr interface{}
		`)
//line stores.qtpl:35
		StreamColumnAddresses(qw422016, &ColumnAddressesInput{
			FieldName:  "field.Name",
			TargetName: "fieldAddr",
			RecordName: "record",
			ErrName:    "err",
			Record:     record,
		})
//line stores.qtpl:41
		qw422016.N().S(`
		if err != nil {
			return err
		}
		columnNames[i] = field.String()
		columnValues[i] = fieldAddr
	}
	builder := store.conn.Builder().Insert(`)
//line stores.qtpl:48
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:48
		qw422016.N().S(`.Table()).Columns(columnNames...).Values(columnValues...)`)
//line stores.qtpl:48
		if record.FieldAutoInc != nil {
//line stores.qtpl:48
			qw422016.N().S(`.Suffix("RETURNING `)
//line stores.qtpl:48
			qw422016.E().S(record.FieldAutoInc.Name)
//line stores.qtpl:48
			qw422016.N().S(`")`)
//line stores.qtpl:48
		}
//line stores.qtpl:48
		qw422016.N().S(`

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
`)
//line stores.qtpl:54
		if record.FieldAutoInc == nil {
//line stores.qtpl:54
			qw422016.N().S(`
	_, err := store.conn.ExecContext(ctx, sql, args...)
	return err
`)
//line stores.qtpl:57
		} else {
//line stores.qtpl:57
			qw422016.N().S(`
	var id `)
//line stores.qtpl:58
			qw422016.E().S(record.FieldAutoInc.Type)
//line stores.qtpl:58
			qw422016.N().S(`
	err := store.conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return err
	}
	record.`)
//line stores.qtpl:63
			qw422016.E().S(record.FieldAutoInc.GoName)
//line stores.qtpl:63
			qw422016.N().S(` = id
	return nil
`)
//line stores.qtpl:65
		}
//line stores.qtpl:65
		qw422016.N().S(`
}

func (store *`)
//line stores.qtpl:68
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:68
		qw422016.N().S(`) Update(record *`)
//line stores.qtpl:68
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:68
		qw422016.N().S(`, fields ...SchemaField) (int64, error) {
	return store.UpdateContext(context.Background(), record, fields...)
}

func (store *`)
//line stores.qtpl:72
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:72
		qw422016.N().S(`) UpdateContext(ctx context.Context, record *`)
//line stores.qtpl:72
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:72
		qw422016.N().S(`, fields ...SchemaField) (int64, error) {
	if len(fields) == 0 {
		fields = `)
//line stores.qtpl:74
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:74
		qw422016.N().S(`Fields
	}
	columnNames := make([]string, len(fields))
	columnValues := make([]interface{}, len(fields))

	builder := store.conn.Builder().Update(`)
//line stores.qtpl:79
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:79
		qw422016.N().S(`.Table())
	var err error
	for i, field := range fields {
		var fieldAddr interface{}
		`)
//line stores.qtpl:83
		StreamColumnAddresses(qw422016, &ColumnAddressesInput{
			FieldName:  "field.Name",
			TargetName: "fieldAddr",
			RecordName: "record",
			ErrName:    "err",
			Record:     record,
		})
//line stores.qtpl:89
		qw422016.N().S(`
		if err != nil {
			return err
		}
		builder = builder.Set(field.String(), fieldAddr)
	}
`)
//line stores.qtpl:95
		for _, field := range record.PrimaryKey {
//line stores.qtpl:95
			qw422016.N().S(`
	builder.Where(query.Eq(query.Raw("`)
//line stores.qtpl:96
			qw422016.E().J(field.Name)
//line stores.qtpl:96
			qw422016.N().S(`"), record.`)
//line stores.qtpl:96
			qw422016.E().S(field.GoName)
//line stores.qtpl:96
			qw422016.N().S(`))
`)
//line stores.qtpl:97
		}
//line stores.qtpl:97
		qw422016.N().S(`
	r, err := builder.ExecContext(ctx)
	if err != nil {
		return err
	}
	return r.RowsAffected()
}

`)
//line stores.qtpl:105
	}
//line stores.qtpl:106
}

//line stores.qtpl:106
func WriteStores(qq422016 qtio422016.Writer, input *StoresInput) {
//line stores.qtpl:106
	qw422016 := qt422016.AcquireWriter(qq422016)
//line stores.qtpl:106
	StreamStores(qw422016, input)
//line stores.qtpl:106
	qt422016.ReleaseWriter(qw422016)
//line stores.qtpl:106
}

//line stores.qtpl:106
func Stores(input *StoresInput) string {
//line stores.qtpl:106
	qb422016 := qt422016.AcquireByteBuffer()
//line stores.qtpl:106
	WriteStores(qb422016, input)
//line stores.qtpl:106
	qs422016 := string(qb422016.B)
//line stores.qtpl:106
	qt422016.ReleaseByteBuffer(qb422016)
//line stores.qtpl:106
	return qs422016
//line stores.qtpl:106
}
