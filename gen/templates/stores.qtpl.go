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
// Code generated by ormgen; DO NOT EDIT.

package `)
//line stores.qtpl:4
	qw422016.E().S(input.Package.Name)
//line stores.qtpl:4
	qw422016.N().S(`

import (
	"context"
	"github.com/pkg/errors"
	
	"github.com/setare/orm"
	"github.com/setare/orm/query"

`)
//line stores.qtpl:13
	if !input.ModelsPackage.Equals(input.Package) {
//line stores.qtpl:13
		qw422016.N().S(`
	"`)
//line stores.qtpl:14
		qw422016.E().S(input.ModelsPackage.ImportPath)
//line stores.qtpl:14
		qw422016.N().S(`"
`)
//line stores.qtpl:15
	}
//line stores.qtpl:15
	qw422016.N().S(`
)

`)
//line stores.qtpl:18
	for _, record := range input.Records {
//line stores.qtpl:18
		qw422016.N().S(`type `)
//line stores.qtpl:19
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:19
		qw422016.N().S(` struct {
	conn orm.Connection
}

func New`)
//line stores.qtpl:23
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:23
		qw422016.N().S(`(conn orm.Connection) *`)
//line stores.qtpl:23
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:23
		qw422016.N().S(` {
	return &`)
//line stores.qtpl:24
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:24
		qw422016.N().S(`{
		conn,
	}
}

func (store *`)
//line stores.qtpl:29
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:29
		qw422016.N().S(`) Insert(record *`)
//line stores.qtpl:29
		qw422016.E().S(input.ModelsPackage.Ref(input.Package, record.Name))
//line stores.qtpl:29
		qw422016.N().S(`, fields ...orm.SchemaField) error {
	return store.InsertContext(context.Background(), record, fields...)
}

func (store *`)
//line stores.qtpl:33
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:33
		qw422016.N().S(`) InsertContext(ctx context.Context, record *`)
//line stores.qtpl:33
		qw422016.E().S(input.ModelsPackage.Ref(input.Package, record.Name))
//line stores.qtpl:33
		qw422016.N().S(`, fields ...orm.SchemaField) error {
	if len(fields) == 0 {
		fields = `)
//line stores.qtpl:35
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:35
		qw422016.N().S(`Fields
	}

	if biH, ok := record.(orm.HookBeforeInsert); ok {
		err := biH.BeforeInsert(ctx, fields...)
		if err != nil {
			return err
		}
	}

	if bsH, ok := record.(orm.HookBeforeSave); ok {
		err := bsH.BeforeSave(ctx, fields...)
		if err != nil {
			return err
		}
	}

	aiH, aiHOk := record.(orm.HookAfterInsert)
	asH, asHOk := record.(orm.HookAfterSave)

	columnNames := make([]string, len(fields))
	columnValues := make([]interface{}, len(fields))
	var err error
	for i, field := range fields {
		var fieldAddr interface{}
		`)
//line stores.qtpl:60
		StreamColumnAddresses(qw422016, &ColumnAddressesInput{
			FieldName:  "field.Name()",
			TargetName: "fieldAddr",
			RecordName: "record",
			ErrName:    "err",
			Record:     record,
		})
//line stores.qtpl:66
		qw422016.N().S(`
		if err != nil {
			if aiHOk {
				aiH.AfterInsert(ctx, err, fields...)
			}
			if asHOk {
				asH.AfterSave(ctx, err, fields...)
			}
			return err
		}
		columnNames[i] = field.String()
		columnValues[i] = fieldAddr
	}
	builder := store.conn.Builder().Insert(`)
//line stores.qtpl:79
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:79
		qw422016.N().S(`.Table()).Columns(columnNames...).Values(columnValues...)`)
//line stores.qtpl:79
		if record.FieldAutoInc != nil {
//line stores.qtpl:79
			qw422016.N().S(`.Suffix("RETURNING `)
//line stores.qtpl:79
			qw422016.E().S(record.FieldAutoInc.Name)
//line stores.qtpl:79
			qw422016.N().S(`")`)
//line stores.qtpl:79
		}
//line stores.qtpl:79
		qw422016.N().S(`

	sql, args, err := builder.ToSql()
	if err != nil {
		if aiHOk {
			aiH.AfterInsert(ctx, err, fields...)
		}
		if asHOk {
			asH.AfterSave(ctx, err, fields...)
		}
		return err
	}
`)
//line stores.qtpl:91
		if record.FieldAutoInc == nil {
//line stores.qtpl:91
			qw422016.N().S(`
	_, err = store.conn.ExecContext(ctx, sql, args...)

	if aiHOk {
		aiH.AfterInsert(ctx, err, fields...)
	}
	if asHOk {
		asH.AfterSave(ctx, err, fields...)
	}

	return err
`)
//line stores.qtpl:102
		} else {
//line stores.qtpl:102
			qw422016.N().S(`
	var id `)
//line stores.qtpl:103
			qw422016.E().S(record.FieldAutoInc.Type)
//line stores.qtpl:103
			qw422016.N().S(`
	err = store.conn.QueryRowContext(ctx, sql, args...).Scan(&id)
	if err != nil {
		if aiHOk {
			aiH.AfterInsert(ctx, err, fields...)
		}
		if asHOk {
			asH.AfterSave(ctx, err, fields...)
		}
		return err
	}
	record.`)
//line stores.qtpl:114
			qw422016.E().S(record.FieldAutoInc.GoName)
//line stores.qtpl:114
			qw422016.N().S(` = id
	if aiHOk {
		aiH.AfterInsert(ctx, nil, fields...)
	}
	if asHOk {
		asH.AfterSave(ctx, err, fields...)
	}
	return nil
`)
//line stores.qtpl:122
		}
//line stores.qtpl:122
		qw422016.N().S(`
}

func (store *`)
//line stores.qtpl:125
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:125
		qw422016.N().S(`) Update(record *`)
//line stores.qtpl:125
		qw422016.E().S(input.ModelsPackage.Ref(input.Package, record.Name))
//line stores.qtpl:125
		qw422016.N().S(`, fields ...orm.SchemaField) (int64, error) {
	return store.UpdateContext(context.Background(), record, fields...)
}

func (store *`)
//line stores.qtpl:129
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:129
		qw422016.N().S(`) UpdateContext(ctx context.Context, record *`)
//line stores.qtpl:129
		qw422016.E().S(input.ModelsPackage.Ref(input.Package, record.Name))
//line stores.qtpl:129
		qw422016.N().S(`, fields ...orm.SchemaField) (int64, error) {
	if len(fields) == 0 {
		fields = `)
//line stores.qtpl:131
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:131
		qw422016.N().S(`Fields
	}

	if buH, ok := record.(orm.HookBeforeUpdate); ok {
		err := buH.BeforeInsert(ctx, fields...)
		if err != nil {
			return err
		}
	}
	
	if bsH, ok := record.(orm.HookBeforeSave); ok {
		err := bsH.BeforeSave(ctx, fields...)
		if err != nil {
			return err
		}
	}

	auH, auHOk := record.(orm.HookAfterUpdate)
	asH, asHOk := record.(orm.HookAfterSave)

	builder := store.conn.Builder().Update(`)
//line stores.qtpl:151
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:151
		qw422016.N().S(`.Table())
	var err error
	for _, field := range fields {
		var fieldAddr interface{}
		`)
//line stores.qtpl:155
		StreamColumnAddresses(qw422016, &ColumnAddressesInput{
			FieldName:  "field.Name()",
			TargetName: "fieldAddr",
			RecordName: "record",
			ErrName:    "err",
			Record:     record,
		})
//line stores.qtpl:161
		qw422016.N().S(`
		if err != nil {
			if auHOk {
				auH.AfterUpdate(ctx, err, fields...)
			}
			if asHOk {
				asH.AfterSave(ctx, err, fields...)
			}
			return 0, err
		}
		builder = builder.Set(field.String(), fieldAddr)
	}
`)
//line stores.qtpl:173
		for _, field := range record.PrimaryKey {
//line stores.qtpl:173
			qw422016.N().S(`
	builder.Where(query.Eq(query.Raw("`)
//line stores.qtpl:174
			qw422016.E().J(field.Name)
//line stores.qtpl:174
			qw422016.N().S(`"), record.`)
//line stores.qtpl:174
			qw422016.E().S(field.GoName)
//line stores.qtpl:174
			qw422016.N().S(`))
`)
//line stores.qtpl:175
		}
//line stores.qtpl:175
		qw422016.N().S(`
	r, err := builder.ExecContext(ctx)
	if err != nil {
		if auHOk {
			auH.AfterUpdate(ctx, err, fields...)
		}
		if asHOk {
			asH.AfterSave(ctx, err, fields...)
		}
		return 0, err
	}
	rowsAffected, err := r.RowsAffected()
	if auHOk {
		auH.AfterUpdate(ctx, err, fields...)
	}
	if asHOk {
		asH.AfterSave(ctx, err, fields...)
	}
	return rowsAffected, err
}

func (store *`)
//line stores.qtpl:196
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:196
		qw422016.N().S(`) Delete(records ...*`)
//line stores.qtpl:196
		qw422016.E().S(input.ModelsPackage.Ref(input.Package, record.Name))
//line stores.qtpl:196
		qw422016.N().S(`) (int64, error) {
	return store.DeleteContext(context.Background(), records...)
}

func (store *`)
//line stores.qtpl:200
		qw422016.E().S(record.Store.Type)
//line stores.qtpl:200
		qw422016.N().S(`) DeleteContext(ctx context.Context, records ...*`)
//line stores.qtpl:200
		qw422016.E().S(input.ModelsPackage.Ref(input.Package, record.Name))
//line stores.qtpl:200
		qw422016.N().S(`) (int64, error) {
	builder := store.conn.Builder().Delete(`)
//line stores.qtpl:201
		qw422016.E().S(record.Schema.InternalRef)
//line stores.qtpl:201
		qw422016.N().S(`.Table())

	if len(records) == 0 {
		ids := make([]interface{}, len(records) * `)
//line stores.qtpl:204
		qw422016.N().D(len(record.PrimaryKey))
//line stores.qtpl:204
		qw422016.N().S(`)

`)
//line stores.qtpl:206
		if len(record.PrimaryKey) == 1 {
//line stores.qtpl:206
			qw422016.N().S(`
		for i, record := range records {
			if bdH, bdHOk := record.(orm.HookBeforeDelete); bdHOk {
				err := bdH.BeforeDelete(ctx)
				if err != nil {
					return err
				}
			}
			ids[i] = record.`)
//line stores.qtpl:214
			qw422016.E().S(record.PrimaryKey[0].GoName)
//line stores.qtpl:214
			qw422016.N().S(`
		}
		builder = builder.Where(query.In(query.Raw("`)
//line stores.qtpl:216
			qw422016.E().J(record.PrimaryKey[0].Name)
//line stores.qtpl:216
			qw422016.N().S(`"), ids))
`)
//line stores.qtpl:217
		} else {
//line stores.qtpl:217
			qw422016.N().S(`
		ors := make([]sq.Sqlizer, 0, len(reocrds))
`)
//line stores.qtpl:219
			for i, field := range record.PrimaryKey {
//line stores.qtpl:219
				qw422016.N().S(`		field`)
//line stores.qtpl:220
				qw422016.N().D(i)
//line stores.qtpl:220
				qw422016.N().S(` := query.Raw("`)
//line stores.qtpl:220
				qw422016.E().J(field.Name)
//line stores.qtpl:220
				qw422016.N().S(`")
`)
//line stores.qtpl:221
			}
//line stores.qtpl:221
			qw422016.N().S(`		for i, record := range records {
			if bdH, bdHOk := record.(orm.HookBeforeDelete); bdHOk {
				err := bdH.BeforeDelete(ctx)
				if err != nil {
					return err
				}
			}

			ors[i] = query.And(
`)
//line stores.qtpl:231
			for i, field := range record.PrimaryKey {
//line stores.qtpl:231
				qw422016.N().S(`				query.Eq(field`)
//line stores.qtpl:232
				qw422016.N().D(i)
//line stores.qtpl:232
				qw422016.N().S(`, records[0].`)
//line stores.qtpl:232
				qw422016.E().S(field.GoName)
//line stores.qtpl:232
				qw422016.N().S(`),
`)
//line stores.qtpl:233
			}
//line stores.qtpl:233
			qw422016.N().S(`			)
		}
		builder = builder.Where(ors)
`)
//line stores.qtpl:237
		}
//line stores.qtpl:237
		qw422016.N().S(`
	} else if len(records) == 1 {
		if bdH, bdHOk := records[0].(orm.HookBeforeDelete); bdHOk {
			err := bdH.BeforeDelete(ctx)
			if err != nil {
				return err
			}
		}

`)
//line stores.qtpl:246
		if len(record.PrimaryKey) == 1 {
//line stores.qtpl:246
			qw422016.N().S(`		builder = builder.Where(query.Eq(query.Raw("`)
//line stores.qtpl:247
			qw422016.E().J(record.PrimaryKey[0].Name)
//line stores.qtpl:247
			qw422016.N().S(`"), records[0].`)
//line stores.qtpl:247
			qw422016.E().S(record.PrimaryKey[0].GoName)
//line stores.qtpl:247
			qw422016.N().S(`))
`)
//line stores.qtpl:248
		} else {
//line stores.qtpl:248
			qw422016.N().S(`
		builder = builder
`)
//line stores.qtpl:250
			for _, field := range record.PrimaryKey {
//line stores.qtpl:250
				qw422016.N().S(`	.Where(query.In(query.Raw("`)
//line stores.qtpl:251
				qw422016.E().J(field.Name)
//line stores.qtpl:251
				qw422016.N().S(`"), records[0].`)
//line stores.qtpl:251
				qw422016.E().S(field.GoName)
//line stores.qtpl:251
				qw422016.N().S(`))
`)
//line stores.qtpl:252
			}
//line stores.qtpl:253
		}
//line stores.qtpl:253
		qw422016.N().S(`	} else {
		return 0, nil
	}

	r, err := builder.ExecContext(ctx)
	if err != nil {
		for _, record := range records {
			if adH, adHOk := records[0].(orm.HookAfterDelete); adHOk {
				adH.AfterDelete(ctx, err)
			}
		}
		return 0, err
	}
	return r.RowsAffected()
}

`)
//line stores.qtpl:270
	}
//line stores.qtpl:271
}

//line stores.qtpl:271
func WriteStores(qq422016 qtio422016.Writer, input *StoresInput) {
//line stores.qtpl:271
	qw422016 := qt422016.AcquireWriter(qq422016)
//line stores.qtpl:271
	StreamStores(qw422016, input)
//line stores.qtpl:271
	qt422016.ReleaseWriter(qw422016)
//line stores.qtpl:271
}

//line stores.qtpl:271
func Stores(input *StoresInput) string {
//line stores.qtpl:271
	qb422016 := qt422016.AcquireByteBuffer()
//line stores.qtpl:271
	WriteStores(qb422016, input)
//line stores.qtpl:271
	qs422016 := string(qb422016.B)
//line stores.qtpl:271
	qt422016.ReleaseByteBuffer(qb422016)
//line stores.qtpl:271
	return qs422016
//line stores.qtpl:271
}
