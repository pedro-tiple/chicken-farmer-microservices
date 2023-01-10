package farmer

import (
	"chicken-farmer/backend/internal/pkg"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"

	"github.com/google/uuid"
)

type FarmService struct {
	grpcClient internalGrpc.FarmServiceClient
}

var _ IFarmService = &FarmService{}

func ProvideFarmService(grpcClient internalGrpc.FarmServiceClient) *FarmService {
	return &FarmService{
		grpcClient: grpcClient,
	}
}

func (f FarmService) NewFarm(
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

func (f FarmService) DeleteFarm(ctx context.Context, farmID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}
