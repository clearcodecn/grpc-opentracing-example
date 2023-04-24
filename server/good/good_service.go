package good

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go-opentracing-example/pb"
	"go-opentracing-example/pkg/logger"
	"go-opentracing-example/pkg/sleep"
	"time"
)

type GoodService struct {
	pb.UnimplementedGoodServiceServer
}

func New() *GoodService {
	return &GoodService{}
}

func (s *GoodService) GetGoodsByID(ctx context.Context, req *pb.GetGoodsByIDsRequest) (*pb.GetGoodsByIDsResponse, error) {
	logger.WithContext(ctx, "GoodService.GetGoodsByID").Infof("获取商品详情: %v", req.Ids)
	goods := s.queryGoods(ctx)
	return &pb.GetGoodsByIDsResponse{
		Goods: goods,
	}, nil
}

func (s *GoodService) UpdateGoodsStoke(ctx context.Context, req *pb.UpdateGoodsStokeRequest) (*pb.UpdateGoodsStokeResponse, error) {
	logger.WithContext(ctx, "GoodService.UpdateGoodsStoke").Infof("更新商品详情: %v", req.Requests)
	sleep.SleepRandom()
	return &pb.UpdateGoodsStokeResponse{}, nil
}

var goods = []*pb.Good{
	{
		Id:    1,
		Name:  "商品1",
		Stoke: 999,
		Price: 80,
	},
	{
		Id:    2,
		Name:  "商品2",
		Stoke: 999,
		Price: 100,
	},
	{
		Id:    3,
		Name:  "商品3",
		Stoke: 999,
		Price: 80,
	},
}

func (s *GoodService) queryGoods(ctx context.Context) []*pb.Good {
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		sp = opentracing.StartSpan("queryGoods")
	} else {
		spContext := sp.Context()
		sp = opentracing.StartSpan("queryGoods",
			opentracing.ChildOf(spContext),
		)

		sp.SetTag("nodeid", "127.0.0.1")

	}

	defer sp.Finish()

	time.Sleep(1 * time.Second)

	return goods
}
