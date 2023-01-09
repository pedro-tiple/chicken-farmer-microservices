package farmer

import (
	"context"

	"github.com/google/uuid"
)

type IDataSource interface {
	GetFarmer(ctx context.Context, farmerID uuid.UUID) (Farmer, error)

	InsertFarmer(ctx context.Context, farmer Farmer) (farmerID uuid.UUID, err error)
}

type Controller struct {
	datasource IDataSource
	currentDay uint
}

var _ IController = &Controller{}

func ProvideController(
	datasource IDataSource,
) *Controller {
	return &Controller{
		datasource: datasource,
		currentDay: 0,
	}
}

func (c *Controller) Register(
	ctx context.Context, farmerName, farmName, password string,
) (Farmer, error) {
	// TODO implement me
	panic("implement me")
}

func (c *Controller) GetGoldEggs(ctx context.Context) (uint, error) {
	// TODO implement me
	return 10, nil
}

func (c *Controller) SpendGoldEggs(ctx context.Context, amount uint) error {
	// TODO implement me
	return nil
}
