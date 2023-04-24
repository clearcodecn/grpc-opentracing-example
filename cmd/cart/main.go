package main

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-opentracing-example/pb"
	"go-opentracing-example/pkg/logger"
	"go-opentracing-example/pkg/requestid"
	"go-opentracing-example/pkg/tracing"
	"go-opentracing-example/server/cart"
	"google.golang.org/grpc"
	"net"
)

func main() {
	tracer, closer := tracing.Init("cartService")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			requestid.WithRequestID(),
			grpc_opentracing.UnaryServerInterceptor(),
		),
	)

	service := cart.New()
	pb.RegisterCartServiceServer(s, service)

	ln, err := net.Listen("tcp", "127.0.0.1:5002")
	if err != nil {
		panic(err)
	}

	logger.WithContext(context.Background(), "main").Infof("cart server start at: 5001")

	s.Serve(ln)
}
