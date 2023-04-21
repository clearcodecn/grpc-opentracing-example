package main

import (
	"context"
	"go-opentracing-example/pb"
	"go-opentracing-example/pkg/logger"
	"go-opentracing-example/pkg/requestid"
	"go-opentracing-example/pkg/tracing"
	"go-opentracing-example/server/good"
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

	service := good.New()
	pb.RegisterGoodServiceServer(s, service)

	ln, err := net.Listen("tcp", "127.0.0.1:5001")
	if err != nil {
		panic(err)
	}

	logger.WithContext(context.Background(), "main").Infof("good server start at: 5001")

	s.Serve(ln)
}
