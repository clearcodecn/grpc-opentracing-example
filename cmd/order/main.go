package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go-opentracing-example/pb"
	"go-opentracing-example/pkg/event"
	"go-opentracing-example/pkg/logger"
	"go-opentracing-example/pkg/requestid"
	"go-opentracing-example/pkg/tracing"
	"go-opentracing-example/server/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

func main() {

	tracer, closer := tracing.Init("orderService")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			ctx = requestid.NewWithContext(ctx, "")
			resp, err = handler(ctx, req)

			return resp, err
		},
			tracing.Middleware(),
		),
	)

	goodsCC, err := grpc.Dial(
		"localhost:5001",
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			id := requestid.RequestID(ctx)
			ctx = metadata.AppendToOutgoingContext(ctx, "x-request-id", id)
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	goodClient := pb.NewGoodServiceClient(goodsCC)

	cartCC, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	cartClient := pb.NewCartServiceClient(cartCC)

	orderServer := order.NewOrderService(cartClient, goodClient, event.New())
	pb.RegisterOrderServiceServer(s, orderServer)

	ln, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}

	logger.WithContext(context.Background(), "main").Infof("order server start at: 5000")

	s.Serve(ln)
}
