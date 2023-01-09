package farm

import (
	"chicken-farmer/backend/internal/farm/ctxfarm"
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

const (
	PurchaseCostChicken = 1
	PurchaseCostFeedBag = 1
	PurchaseCostBarn    = 10

	FeedChickenCost = 1
	FeedPerBag      = 10

	MaxChickenRestingDays = 5

	EggTypeNormal = iota
	EggTypeGolden
)

var (
	ErrFarmNotYours    = errors.New("farm doesn't belong to you")
	ErrChickenNotYours = errors.New("chicken doesn't belong to you")
	ErrChickenResting  = errors.New("chicken is resting")
	ErrInvalidEggType  = errors.New("invalid egg type")
)

type IDataSource interface {
	GetFarm(ctx context.Context, chickenID uuid.UUID) (Farm, error)
	GetBarnsOfFarm(ctx context.Context, farmID uuid.UUID) ([]Barn, error)
	GetChickensOfBarn(ctx context.Context, barnID uuid.UUID) ([]Chicken, error)
	GetChicken(ctx context.Context, chickenID uuid.UUID) (Chicken, error)

	InsertChicken(ctx context.Context, chicken Chicken) (chickenID uuid.UUID, err error)
	InsertBarn(ctx context.Context, barn Barn) (barnID uuid.UUID, err error)

	UpdateChickenRestingUntil(ctx context.Context, chickenID uuid.UUID, day uint) error

	IncrementBarnFeed(ctx context.Context, barnID uuid.UUID, amount uint) error
	DecrementBarnFeed(ctx context.Context, barnID uuid.UUID, amount uint) error
	IncrementChickenEggLayCount(ctx context.Context, chickenID uuid.UUID, eggType int) error
}

type IFarmerService interface {
	SpendGoldEggs(ctx context.Context, amount uint) error
	GetGoldEggs(ctx context.Context) (uint, error)
}

type Controller struct {
	datasource    IDataSource
	farmerService IFarmerService
	currentDay    uint
}

var _ IController = &Controller{}

func ProvideController(
	datasource IDataSource,
	farmerService IFarmerService,
) *Controller {
	return &Controller{
		datasource:    datasource,
		farmerService: farmerService,
		currentDay:    1,
	}
}

func (c *Controller) GetFarm(ctx context.Context) (GetFarmResult, error) {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return GetFarmResult{}, err
	}

	farm, err := c.datasource.GetFarm(ctx, ctxData.FarmID)
	if err != nil {
		return GetFarmResult{}, err
	}

	// TODO if no farm exists, assume it's a new registration and create a new farm.
	// This avoids creating a Farm dependency on the Farmer service.

	if farm.OwnerID != ctxData.FarmerID {
		return GetFarmResult{}, ErrFarmNotYours
	}

	barns, err := c.datasource.GetBarnsOfFarm(ctx, farm.ID)
	if err != nil {
		return GetFarmResult{}, err
	}

	resultBarns := make([]getFarmResultBarn, len(barns))
	errGrp, errGrpCtx := errgroup.WithContext(ctx)
	errGrp.SetLimit(5)

	for i, barn := range barns {
		i, barn := i, barn //nolint:varnamelen

		errGrp.Go(func() error {
			chickens, err := c.datasource.GetChickensOfBarn(errGrpCtx, barn.ID)
			if err != nil {
				return err
			}
			resultBarns[i] = getFarmResultBarn{
				Barn:     barn,
				Chickens: chickens,
			}

			return nil
		})
	}

	if err := errGrp.Wait(); err != nil {
		return GetFarmResult{}, err
	}

	// Gold egg count lives in another service for this implementation so must
	// go fetch it there.
	goldEggCount, err := c.farmerService.GetGoldEggs(ctx)
	if err != nil {
		return GetFarmResult{}, err
	}

	return GetFarmResult{
		Farm:         farm,
		GoldEggCount: goldEggCount,
		CurrentDay:   c.currentDay,
		Barns:        resultBarns,
	}, nil
}

func (c *Controller) BuyBarn(ctx context.Context) error {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return err
	}

	if err := c.farmerService.SpendGoldEggs(ctx, PurchaseCostBarn); err != nil {
		return err
	}

	_, err = c.datasource.InsertBarn(ctx, Barn{
		FarmID: ctxData.FarmID,
	})
	if err != nil {
		return err
	}

	// TODO send purchase event.

	return nil
}

func (c *Controller) BuyFeedBag(ctx context.Context, barnID uuid.UUID, amount uint) error {
	if err := c.farmerService.SpendGoldEggs(ctx, PurchaseCostFeedBag*amount); err != nil {
		return err
	}

	if err := c.datasource.IncrementBarnFeed(ctx, barnID, amount*FeedPerBag); err != nil {
		return err
	}

	// TODO send purchase event.

	return nil
}

func (c *Controller) BuyChicken(
	ctx context.Context, barnID uuid.UUID,
) error {
	if err := c.farmerService.SpendGoldEggs(ctx, PurchaseCostBarn); err != nil {
		return err
	}

	// Maybe have the gold egg chance be on a bell curve?
	rand.Seed(time.Now().Unix())

	_, err := c.datasource.InsertChicken(ctx, Chicken{
		ID:            uuid.New(),
		BarnID:        barnID,
		DateOfBirth:   c.currentDay,
		GoldEggChance: uint(rand.Intn(99) + 1), // [1,100]
	})
	if err != nil {
		return err
	}

	// TODO send purchase event.

	return nil
}

func (c *Controller) FeedChicken(ctx context.Context, chickenID uuid.UUID) error {
	ctxData, err := ctxfarm.Extract(ctx)
	if err != nil {
		return err
	}

	chicken, err := c.datasource.GetChicken(ctx, chickenID)
	if err != nil {
		return err
	}

	if chicken.OwnerID != ctxData.FarmerID {
		return ErrChickenNotYours
	}

	if chicken.RestingUntil >= c.currentDay {
		return ErrChickenResting
	}

	if err := c.datasource.DecrementBarnFeed(
		ctx, chickenID, FeedChickenCost,
	); err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())

	var eggType = EggTypeNormal
	if rand.Intn(100) <= int(chicken.GoldEggChance) {
		eggType = EggTypeGolden
	}

	if err := c.datasource.IncrementChickenEggLayCount(
		ctx, chickenID, eggType,
	); err != nil {
		return err
	}

	// Must rest at least one day, can rest up to 1 + MaxChickenRestingDays.
	if err := c.datasource.UpdateChickenRestingUntil(
		ctx, chickenID, c.currentDay+1+uint(rand.Intn(MaxChickenRestingDays)),
	); err != nil {
		return err
	}

	// TODO send egg event

	return nil
}

func (c *Controller) FeedChickensOfBarn(
	ctx context.Context, barnID uuid.UUID,
) error {
	// TODO
	return nil
}

func (c *Controller) SetDay(
	ctx context.Context, day uint,
) error {
	c.currentDay = day

	return nil
}

func (c *Controller) GetCurrentDay(ctx context.Context) (uint, error) {
	return c.currentDay, nil
}
