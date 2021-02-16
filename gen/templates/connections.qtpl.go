// Code generated by qtc from "connections.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line connections.qtpl:1
package templates

//line connections.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line connections.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line connections.qtpl:1
func StreamConnections(qw422016 *qt422016.Writer, input *ConnectionsInput) {
//line connections.qtpl:1
	qw422016.N().S(`
// Code generated by ormgen; DO NOT EDIT.

package `)
//line connections.qtpl:4
	qw422016.E().S(input.Package.Name)
//line connections.qtpl:4
	qw422016.N().S(`

import (
	"github.com/jamillosantos/orm"
)

type connection struct {
	orm.ConnectionPgx
}

`)
//line connections.qtpl:14
	for _, record := range input.Records {
//line connections.qtpl:14
		qw422016.N().S(`
func (conn *connection) `)
//line connections.qtpl:15
		qw422016.E().S(record.Query.Type)
//line connections.qtpl:15
		qw422016.N().S(`() *`)
//line connections.qtpl:15
		qw422016.E().S(record.Query.Type)
//line connections.qtpl:15
		qw422016.N().S(` {
	return New`)
//line connections.qtpl:16
		qw422016.E().S(record.Query.Type)
//line connections.qtpl:16
		qw422016.N().S(`(conn)
}
`)
//line connections.qtpl:18
	}
//line connections.qtpl:18
	qw422016.N().S(`
`)
//line connections.qtpl:19
	for _, record := range input.Records {
//line connections.qtpl:19
		qw422016.N().S(`
func (conn *connection) `)
//line connections.qtpl:20
		qw422016.E().S(record.Store.Type)
//line connections.qtpl:20
		qw422016.N().S(`() *`)
//line connections.qtpl:20
		qw422016.E().S(record.Store.Type)
//line connections.qtpl:20
		qw422016.N().S(` {
	return New`)
//line connections.qtpl:21
		qw422016.E().S(record.Store.Type)
//line connections.qtpl:21
		qw422016.N().S(`(conn)
}
`)
//line connections.qtpl:23
	}
//line connections.qtpl:23
	qw422016.N().S(`

var DefaultConnection connection

`)
//line connections.qtpl:27
}

//line connections.qtpl:27
func WriteConnections(qq422016 qtio422016.Writer, input *ConnectionsInput) {
//line connections.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line connections.qtpl:27
	StreamConnections(qw422016, input)
//line connections.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line connections.qtpl:27
}

//line connections.qtpl:27
func Connections(input *ConnectionsInput) string {
//line connections.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
//line connections.qtpl:27
	WriteConnections(qb422016, input)
//line connections.qtpl:27
	qs422016 := string(qb422016.B)
//line connections.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
//line connections.qtpl:27
	return qs422016
//line connections.qtpl:27
}
