// Code generated by qtc from "schema.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line schema.qtpl:1
package templates

//line schema.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line schema.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line schema.qtpl:1
func StreamSchema(qw422016 *qt422016.Writer, input *SchemaInput) {
//line schema.qtpl:1
	qw422016.N().S(`
// Code generated by ormgen; DO NOT EDIT.

package `)
//line schema.qtpl:4
	qw422016.E().S(input.Package.Name)
//line schema.qtpl:4
	qw422016.N().S(`

import (
	"github.com/jamillosantos/orm"
)

`)
//line schema.qtpl:10
	for _, record := range input.Records {
//line schema.qtpl:10
		qw422016.N().S(`type `)
//line schema.qtpl:11
		qw422016.E().S(record.Schema.Type)
//line schema.qtpl:11
		qw422016.N().S(` struct {
	orm.Schema
`)
//line schema.qtpl:13
		for _, field := range record.Fields {
//line schema.qtpl:13
			qw422016.N().S(`	`)
//line schema.qtpl:14
			qw422016.E().S(field.GoName)
//line schema.qtpl:14
			qw422016.N().S(` orm.SchemaField
`)
//line schema.qtpl:15
		}
//line schema.qtpl:15
		qw422016.N().S(`}

var (
	`)
//line schema.qtpl:19
		qw422016.E().S(record.Schema.InternalRef)
//line schema.qtpl:19
		qw422016.N().S(`Fields = []orm.SchemaField{
`)
//line schema.qtpl:20
		for _, field := range record.Fields {
//line schema.qtpl:20
			qw422016.N().S(`		orm.NewSchemaField("`)
//line schema.qtpl:21
			qw422016.E().S(field.Name)
//line schema.qtpl:21
			qw422016.N().S(`"),
`)
//line schema.qtpl:22
		}
//line schema.qtpl:22
		qw422016.N().S(`	}
	`)
//line schema.qtpl:24
		qw422016.E().S(record.Schema.InternalRef)
//line schema.qtpl:24
		qw422016.N().S(` = &`)
//line schema.qtpl:24
		qw422016.E().S(record.Schema.Type)
//line schema.qtpl:24
		qw422016.N().S(`{
		Schema: orm.NewSchema("`)
//line schema.qtpl:25
		qw422016.E().J(record.TableName)
//line schema.qtpl:25
		qw422016.N().S(`", `)
//line schema.qtpl:25
		qw422016.E().S(record.Schema.InternalRef)
//line schema.qtpl:25
		qw422016.N().S(`Fields...),
`)
//line schema.qtpl:26
		for fieldIndex, field := range record.Fields {
//line schema.qtpl:26
			qw422016.N().S(`		`)
//line schema.qtpl:27
			qw422016.E().S(field.GoName)
//line schema.qtpl:27
			qw422016.N().S(`: `)
//line schema.qtpl:27
			qw422016.E().S(record.Schema.InternalRef)
//line schema.qtpl:27
			qw422016.N().S(`Fields[`)
//line schema.qtpl:27
			qw422016.N().D(fieldIndex)
//line schema.qtpl:27
			qw422016.N().S(`],
`)
//line schema.qtpl:28
		}
//line schema.qtpl:28
		qw422016.N().S(`	}
)
`)
//line schema.qtpl:31
	}
//line schema.qtpl:31
	qw422016.N().S(`

var Schema  = struct {
`)
//line schema.qtpl:34
	for _, record := range input.Records {
//line schema.qtpl:34
		qw422016.N().S(`	`)
//line schema.qtpl:35
		qw422016.E().S(record.Schema.Name)
//line schema.qtpl:35
		qw422016.N().S(` *`)
//line schema.qtpl:35
		qw422016.E().S(record.Schema.Type)
//line schema.qtpl:35
		qw422016.N().S(`
`)
//line schema.qtpl:36
	}
//line schema.qtpl:36
	qw422016.N().S(`}{
`)
//line schema.qtpl:38
	for _, record := range input.Records {
//line schema.qtpl:38
		qw422016.N().S(`	`)
//line schema.qtpl:39
		qw422016.E().S(record.Schema.Name)
//line schema.qtpl:39
		qw422016.N().S(`: `)
//line schema.qtpl:39
		qw422016.E().S(record.Schema.InternalRef)
//line schema.qtpl:39
		qw422016.N().S(`,
`)
//line schema.qtpl:40
	}
//line schema.qtpl:40
	qw422016.N().S(`}

`)
//line schema.qtpl:43
}

//line schema.qtpl:43
func WriteSchema(qq422016 qtio422016.Writer, input *SchemaInput) {
//line schema.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line schema.qtpl:43
	StreamSchema(qw422016, input)
//line schema.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line schema.qtpl:43
}

//line schema.qtpl:43
func Schema(input *SchemaInput) string {
//line schema.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line schema.qtpl:43
	WriteSchema(qb422016, input)
//line schema.qtpl:43
	qs422016 := string(qb422016.B)
//line schema.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line schema.qtpl:43
	return qs422016
//line schema.qtpl:43
}
