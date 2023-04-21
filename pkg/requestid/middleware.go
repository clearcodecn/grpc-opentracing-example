package requestid

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func WithRequestID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		reqid := md.Get("x-request-id")
		if len(reqid) > 0 {
			ctx = NewWithContext(ctx, reqid[0])
		}
		return handler(ctx, req)
	}
}
