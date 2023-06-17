package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var invoker grpc.UnaryInvoker = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
	return nil
}

func TestMiddleware_Log(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		f := New().Log()
		err := f(context.Background(), "", nil, nil, nil, invoker)
		assert.NoError(t, err)
	})
}
