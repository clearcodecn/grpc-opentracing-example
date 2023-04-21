package main

import (
	"context"
	"fmt"
	"go-opentracing-example/pb"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewOrderServiceClient(cc)

	resp, err := client.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CartIds: []int64{1, 2, 3},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.OrderId)
}
