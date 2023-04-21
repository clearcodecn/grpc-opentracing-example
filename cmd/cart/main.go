package main

import (
	"context"
	"go-opentracing-example/pb"
	"go-opentracing-example/pkg/logger"
	"go-opentracing-example/pkg/requestid"
	"go-opentracing-example/pkg/tracing"
	"go-opentracing-example/server/cart"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			requestid.WithRequestID(),
			tracing.Middleware(),
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
