package order

import (
	context "context"
	"go-opentracing-example/pb"
	"go-opentracing-example/pkg/event"
	"go-opentracing-example/pkg/logger"
	"time"
)

type OrderService struct {
	cartClient pb.CartServiceClient
	goodClient pb.GoodServiceClient
	publisher  event.CreateOrderPublisher

	pb.UnimplementedOrderServiceServer
}

func NewOrderService(
	cartClient pb.CartServiceClient,
	goodClient pb.GoodServiceClient,
	publisher event.CreateOrderPublisher,
) *OrderService {
	return &OrderService{
		cartClient: cartClient,
		goodClient: goodClient,
		publisher:  publisher,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	log := logger.WithContext(ctx, "CartService.CreateOrder")

	log.Infof("创建订单: %v", request.CartIds)

	// 1. 获取购物车详情.
	cartDetail, err := s.cartClient.GetCardDetailByIDs(ctx, &pb.GetCardDetailByIDsRequest{
		Ids: request.CartIds,
	})

	if err != nil {
		log.WithError(err).Errorf("获取购物车详情失败:%v", request.CartIds)
		return nil, err
	}

	// 2. 获取商品详情
	var (
		goodIds            []int64
		cartIds            []int64
		updateStokeRequest []*pb.StokeRequest
	)
	for _, c := range cartDetail.Carts {
		goodIds = append(goodIds, c.GoodId)
		cartIds = append(cartIds, c.Id)
		updateStokeRequest = append(updateStokeRequest, &pb.StokeRequest{
			GoodId: c.GoodId,
			Stoke:  1, // 忘了写库存字段了
		})
	}

	// 2. 获取商品详情
	goodsResp, err := s.goodClient.GetGoodsByID(ctx, &pb.GetGoodsByIDsRequest{
		Ids: goodIds,
	})
	if err != nil {
		log.WithError(err).Errorf("获取商品详情失败:%v", goodIds)
		return nil, err
	}

	// 创建订单.
	realOrder := s.createRealOrder(ctx, goodsResp.Goods)

	log.Infof("创建订单: %v", realOrder)

	// 更新库存
	_, err = s.goodClient.UpdateGoodsStoke(ctx, &pb.UpdateGoodsStokeRequest{
		Requests: updateStokeRequest,
	})
	if err != nil {
		log.WithError(err).Errorf("更新商品库存失败:%v", goodIds)
		return nil, err
	}

	// 更新购物车
	_, err = s.cartClient.UpdateCartByIDs(ctx, &pb.UpdateCartByIDsRequest{
		Ids: cartIds,
	})

	if err != nil {
		log.WithError(err).Errorf("更新商品库存失败:%v", goodIds)
		return nil, err
	}

	// 发送下单事件
	s.publisher.Notify(realOrder)

	// 返回
	return &pb.CreateOrderResponse{
		OrderId: realOrder.Id,
	}, nil
}

func (s *OrderService) createRealOrder(ctx context.Context, goods []*pb.Good) *pb.Order {

	logger.WithContext(ctx, "OrderService.createRealOrder").Infof("商品是: %+v", goods)

	// 创建订单

	return &pb.Order{
		Id:        1,
		Uid:       1,
		Price:     310,
		Status:    pb.Order_Created,
		CreatedAt: time.Now().Unix(),
	}
}
