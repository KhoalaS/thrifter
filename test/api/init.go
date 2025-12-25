package test

import (
	"github.com/KhoalaS/thrifter"
	"github.com/KhoalaS/thrifter/test/api/binding_test"
	"github.com/v2pro/wombat/generic"
)

var api = thrifter.Config{
	Protocol: thrifter.ProtocolBinary,
}.Froze()

//go:generate go install github.com/KhoalaS/thrifter/cmd/thrifter
//go:generate $GOPATH/bin/thrifter -pkg  github.com/KhoalaS/thrifter/test/api
func init() {
	generic.Declare(func() {
		api.WillDecodeFromBuffer(
			(*binding_test.TestObject)(nil),
		)
	})
}
