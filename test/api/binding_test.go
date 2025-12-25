package test

import (
	"context"
	"testing"

	"github.com/KhoalaS/thrifter/test/api/binding_test"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/stretchr/testify/require"
)

func Test_binding(t *testing.T) {
	ctx := context.TODO()

	should := require.New(t)
	buf := thrift.NewTMemoryBuffer()
	transport := thrift.NewTFramedTransport(buf)
	proto := thrift.NewTBinaryProtocol(transport, true, true)
	proto.WriteStructBegin(ctx, "hello")
	proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
	proto.WriteI64(ctx, 1024)
	proto.WriteFieldEnd(ctx)
	proto.WriteFieldStop(ctx)
	proto.WriteStructEnd(ctx)
	transport.Flush(ctx)
	var val binding_test.TestObject
	should.NoError(api.Unmarshal(buf.Bytes()[4:], &val))
	should.Equal(int64(1024), val.Field1)
}
