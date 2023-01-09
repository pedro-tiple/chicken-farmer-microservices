package farmer

import (
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

func (f FarmService) NewFarm(ctx context.Context, ownerID uuid.UUID, name string) (farmID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}
