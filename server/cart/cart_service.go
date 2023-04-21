package cart

import (
	context "context"
	"go-opentracing-example/pb"
	"go-opentracing-example/pkg/logger"
	"go-opentracing-example/pkg/sleep"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
}

func New() *CartService {
	return &CartService{}
}

func (c *CartService) GetCardDetailByIDs(ctx context.Context, request *pb.GetCardDetailByIDsRequest) (*pb.GetCardDetailByIDsResponse, error) {

	logger.WithContext(ctx, "CartService.GetCardDetailByIDs").Infof("获取cart: %v", request.Ids)

	sleep.SleepRandom()

	return &pb.GetCardDetailByIDsResponse{
		Carts: []*pb.Cart{
			{
				Id:       1,
				GoodId:   1,
				Selected: true,
			},
			{
				Id:       2,
				GoodId:   2,
				Selected: true,
			},
		},
	}, nil
}

func (c *CartService) UpdateCartByIDs(ctx context.Context, request *pb.UpdateCartByIDsRequest) (*pb.UpdateCartByIDsResponse, error) {
	logger.WithContext(ctx, "CartService.UpdateCartByIDs").Infof("更新 cart: %v, ids=%v", request.Ids, request.Selected)

	sleep.SleepRandom()

	return &pb.UpdateCartByIDsResponse{}, nil
}
