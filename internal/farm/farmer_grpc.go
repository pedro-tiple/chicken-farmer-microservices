package farm

import (
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
)

type FarmerGRPCClient struct {
	grpcClient internalGrpc.FarmerPrivateServiceClient
}

var _ IFarmerService = &FarmerGRPCClient{}

func ProvideFarmerPrivateGRPCClient(
	grpcClient internalGrpc.FarmerPrivateServiceClient,
) *FarmerGRPCClient {
	return &FarmerGRPCClient{
		grpcClient: grpcClient,
	}
}

func (f FarmerGRPCClient) GrantGoldEggs(
	ctx context.Context, amount uint,
) error {
	_, err := f.grpcClient.GrantGoldEggs(
		ctx, &internalGrpc.GrantGoldEggsRequest{
			Amount: uint32(amount),
		},
	)

	return err
}

func (f FarmerGRPCClient) SpendGoldEggs(
	ctx context.Context, amount uint,
) error {
	_, err := f.grpcClient.SpendGoldEggs(
		ctx, &internalGrpc.SpendGoldEggsRequest{
			Amount: uint32(amount),
		},
	)

	return err
}

func (f FarmerGRPCClient) GetGoldEggs(ctx context.Context) (uint, error) {
	result, err := f.grpcClient.GetGoldEggs(
		ctx, &internalGrpc.GetGoldEggsRequest{},
	)
	if err != nil {
		return 0, err
	}

	return uint(result.GetAmount()), nil
}
