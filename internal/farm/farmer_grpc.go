package farm

import (
	cfGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
)

type FarmerService struct {
	grpcClient cfGrpc.FarmerServiceClient
}

var _ IFarmerService = &FarmerService{}

func ProvideFarmerService(grpcClient cfGrpc.FarmerServiceClient) *FarmerService {
	return &FarmerService{
		grpcClient: grpcClient,
	}
}

func (f FarmerService) SpendGoldEggs(ctx context.Context, amount uint) error {
	_, err := f.grpcClient.SpendGoldEggs(
		ctx, &cfGrpc.SpendGoldEggsRequest{
			Amount: uint32(amount),
		})
	return err
}

func (f FarmerService) GetGoldEggs(ctx context.Context) (uint, error) {
	result, err := f.grpcClient.GetGoldEggs(ctx, &cfGrpc.GetGoldEggsRequest{})
	if err != nil {
		return 0, err
	}

	return uint(result.GetAmount()), nil
}
