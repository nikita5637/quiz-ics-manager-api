package log

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/logger"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"google.golang.org/grpc"
)

// Log ...
func (m *Middleware) Log() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := timeutils.TimeNow()
		err := invoker(ctx, method, req, reply, cc, opts...)
		logger.Debugf(ctx, "Invoked RPC method=%s; Duration=%s; Error=%v", method, time.Since(start), err)
		return err
	}
}
