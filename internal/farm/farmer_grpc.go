package farm

import (
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
)

type FarmerService struct {
	grpcClient internalGrpc.FarmerServiceClient
}

var _ IFarmerService = &FarmerService{}

func ProvideFarmerService(grpcClient internalGrpc.FarmerServiceClient) *FarmerService {
	return &FarmerService{
		grpcClient: grpcClient,
	}
}

func (f FarmerService) SpendGoldEggs(ctx context.Context, amount uint) error {
	_, err := f.grpcClient.SpendGoldEggs(
		ctx, &internalGrpc.SpendGoldEggsRequest{
			Amount: uint32(amount),
		})

	return err
}

func (f FarmerService) GetGoldEggs(ctx context.Context) (uint, error) {
	result, err := f.grpcClient.GetGoldEggs(ctx, &internalGrpc.GetGoldEggsRequest{})
	if err != nil {
		return 0, err
	}

	return uint(result.GetAmount()), nil
}
