package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func Middleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		sp := opentracing.SpanFromContext(ctx)
		if sp == nil {
			sp, ctx = opentracing.StartSpanFromContext(ctx, info.FullMethod)
		}

		return handler(ctx, req)
	}
}
