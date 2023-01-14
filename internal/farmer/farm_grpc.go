package farmer

import (
	"chicken-farmer/backend/internal/pkg"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"

	"github.com/google/uuid"
)

type FarmGRPCClient struct {
	grpcClient internalGrpc.FarmServiceClient
}

var _ IFarmService = &FarmGRPCClient{}

func ProvideFarmGRPCClient(
	grpcClient internalGrpc.FarmServiceClient,
) *FarmGRPCClient {
	return &FarmGRPCClient{
		grpcClient: grpcClient,
	}
}

func (f *FarmGRPCClient) NewFarm(
	ctx context.Context, ownerID uuid.UUID, name string,
) (uuid.UUID, error) {
	result, err := f.grpcClient.NewFarm(
		ctx, &internalGrpc.NewFarmRequest{
			OwnerId: ownerID.String(),
			Name:    name,
		},
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return pkg.UUIDFromString(result.GetFarmId()), nil
}

func (f *FarmGRPCClient) DeleteFarm(
	ctx context.Context, farmID uuid.UUID,
) error {
	if _, err := f.grpcClient.DeleteFarm(
		ctx, &internalGrpc.DeleteFarmRequest{FarmId: farmID.String()},
	); err != nil {
		return err
	}

	return nil
}
