package event

import (
	"go-opentracing-example/pb"
	"log"
)

type CreateOrderPublisher interface {
	Notify(order *pb.Order)
}

func New() CreateOrderPublisher {
	return &createOrderPublisher{}
}

type createOrderPublisher struct {
}

func (c createOrderPublisher) Notify(order *pb.Order) {
	log.Printf("推送创建订单事件：%v \n", order)
}
