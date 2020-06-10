// Code generated by qtc from "resultset.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line resultset.qtpl:1
package templates

//line resultset.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line resultset.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line resultset.qtpl:1
func StreamResultSet(qw422016 *qt422016.Writer, input *ResultSetInput) {
//line resultset.qtpl:1
	qw422016.N().S(`
// Code generated by ormgen; DO NOT EDIT.

package `)
//line resultset.qtpl:4
	qw422016.E().S(input.Package.Name)
//line resultset.qtpl:4
	qw422016.N().S(`

import (
	"github.com/pkg/errors"

	"github.com/setare/orm"

`)
//line resultset.qtpl:11
	if !input.ModelsPackage.Equals(input.Package) {
//line resultset.qtpl:11
		qw422016.N().S(`
	"`)
//line resultset.qtpl:12
		qw422016.E().S(input.ModelsPackage.ImportPath)
//line resultset.qtpl:12
		qw422016.N().S(`"
`)
//line resultset.qtpl:13
	}
//line resultset.qtpl:13
	qw422016.N().S(`
)

`)
//line resultset.qtpl:16
	for _, record := range input.Records {
//line resultset.qtpl:16
		qw422016.N().S(`type `)
//line resultset.qtpl:17
		qw422016.E().S(record.ResultSet.Type)
//line resultset.qtpl:17
		qw422016.N().S(` struct {
	orm.ResultSet
	columns []string
	fields []interface{}
}

func New`)
//line resultset.qtpl:23
		qw422016.E().S(record.ResultSet.Type)
//line resultset.qtpl:23
		qw422016.N().S(`(rs orm.ResultSet) (*`)
//line resultset.qtpl:23
		qw422016.E().S(record.ResultSet.Type)
//line resultset.qtpl:23
		qw422016.N().S(`, error) {
	columns, err := rs.Columns()
	if err != nil {
		return nil, err
	}
	return &`)
//line resultset.qtpl:28
		qw422016.E().S(record.ResultSet.Type)
//line resultset.qtpl:28
		qw422016.N().S(`{
		ResultSet: rs,
		columns: columns,
		fields: make([]interface{}, len(columns)),
	}, nil
}

func (rs *`)
//line resultset.qtpl:35
		qw422016.E().S(record.ResultSet.Type)
//line resultset.qtpl:35
		qw422016.N().S(`) Scan(record *`)
//line resultset.qtpl:35
		qw422016.E().S(input.ModelsPackage.Ref(input.Package, record.Name))
//line resultset.qtpl:35
		qw422016.N().S(`) error {
	err := rs.ResultSet.Err()
	if err != nil {
		return err
	}

	for i, column := range rs.columns {
		`)
//line resultset.qtpl:42
		StreamColumnAddresses(qw422016, &ColumnAddressesInput{
			FieldName:  "column",
			TargetName: "rs.fields[i]",
			RecordName: "record",
			ErrName:    "err",
			Record:     record,
		})
//line resultset.qtpl:48
		qw422016.N().S(`
		if err != nil {
			return err
		}
	}
	return rs.ResultSet.Scan(rs.fields...)
}

`)
//line resultset.qtpl:56
	}
//line resultset.qtpl:57
}

//line resultset.qtpl:57
func WriteResultSet(qq422016 qtio422016.Writer, input *ResultSetInput) {
//line resultset.qtpl:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line resultset.qtpl:57
	StreamResultSet(qw422016, input)
//line resultset.qtpl:57
	qt422016.ReleaseWriter(qw422016)
//line resultset.qtpl:57
}

//line resultset.qtpl:57
func ResultSet(input *ResultSetInput) string {
//line resultset.qtpl:57
	qb422016 := qt422016.AcquireByteBuffer()
//line resultset.qtpl:57
	WriteResultSet(qb422016, input)
//line resultset.qtpl:57
	qs422016 := string(qb422016.B)
//line resultset.qtpl:57
	qt422016.ReleaseByteBuffer(qb422016)
//line resultset.qtpl:57
	return qs422016
//line resultset.qtpl:57
}