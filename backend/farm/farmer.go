package farm

import (
	"context"
)

type FarmerService struct {
}

var _ IFarmer = &FarmerService{}

func ProvideFarmerService() *FarmerService {
	// TODO
	return &FarmerService{}
}

func (f FarmerService) SpendGoldEggs(ctx context.Context, amount uint) error {
	//TODO implement me
	return nil
}

func (f FarmerService) GetGoldEggs(ctx context.Context) (uint, error) {
	//TODO implement me
	return 10, nil
}
