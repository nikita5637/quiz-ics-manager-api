package servicename

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ServiceName ...
func (m *Middleware) ServiceName() grpc.UnaryClientInterceptor {
	return func(ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-service-name", "ics-manager")
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
