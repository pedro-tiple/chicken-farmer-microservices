package farm

import (
	"chicken-farmer/backend/internal/farm/ctxFarm"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

const (
	PurchaseCostChicken = 1
	PurchaseCostFeed    = 1
	PurchaseCostBarn    = 10

	FeedChickenCost = 1

	MaxChickenRestingDays = 5

	EggTypeNormal = iota
	EggTypeGolden
)

var (
	ErrChickenNotYours = errors.New("chicken doesn't belong to you")
	ErrChickenResting  = errors.New("chicken is resting")
	ErrInvalidEggType  = errors.New("invalid egg type")
)

type IDatabase interface {
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

type IFarmer interface {
	SpendGoldEggs(ctx context.Context, amount uint) error
	GetGoldEggs(ctx context.Context) (uint, error)
}

type Controller struct {
	database   IDatabase
	farmer     IFarmer
	currentDay uint
}

var _ IController = &Controller{}

func ProvideController(
	database IDatabase,
	farmer IFarmer,
) *Controller {
	return &Controller{
		database: database,
		farmer:   farmer,
	}
}

func (c *Controller) GetFarm(ctx context.Context) (GetFarmResult, error) {
	ctxData, err := ctxFarm.Extract(ctx)
	if err != nil {
		return GetFarmResult{}, err
	}

	farm, err := c.database.GetFarm(ctx, ctxData.FarmID)
	if err != nil {
		return GetFarmResult{}, err
	}

	if farm.OwnerID != ctxData.FarmerID {
		return GetFarmResult{}, errors.New("proto doesn't belong to farmer")
	}

	barns, err := c.database.GetBarnsOfFarm(ctx, farm.ID)
	if err != nil {
		return GetFarmResult{}, err
	}
	fmt.Println(barns, farm.ID)

	resultBarns := make([]getFarmResultBarn, len(barns))
	g, errGrpCtx := errgroup.WithContext(ctx)
	g.SetLimit(5)
	for i, barn := range barns {
		i, barn := i, barn
		g.Go(func() error {
			chickens, err := c.database.GetChickensOfBarn(errGrpCtx, barn.ID)
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

	if err := g.Wait(); err != nil {
		return GetFarmResult{}, err
	}

	// Gold egg count lives in another service for this implementation so must
	// go fetch it there.
	goldEggCount, err := c.farmer.GetGoldEggs(ctx)
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
	ctxData, err := ctxFarm.Extract(ctx)
	if err != nil {
		return err
	}

	if err := c.farmer.SpendGoldEggs(ctx, PurchaseCostBarn); err != nil {
		return err
	}

	_, err = c.database.InsertBarn(ctx, Barn{
		FarmID: ctxData.FarmID,
	})
	if err != nil {
		return err
	}

	// TODO send purchase event.

	return nil
}

func (c *Controller) BuyFeed(ctx context.Context, barnID uuid.UUID, amount uint) error {
	if err := c.farmer.SpendGoldEggs(ctx, PurchaseCostBarn); err != nil {
		return err
	}

	if err := c.database.IncrementBarnFeed(ctx, barnID, amount); err != nil {
		return err
	}

	// TODO send purchase event.

	return nil
}

func (c *Controller) BuyChicken(
	ctx context.Context, barnID uuid.UUID,
) error {
	if err := c.farmer.SpendGoldEggs(ctx, PurchaseCostBarn); err != nil {
		return err
	}

	// Maybe have the gold egg chance be on a bell curve?
	rand.Seed(time.Now().Unix())
	_, err := c.database.InsertChicken(ctx, Chicken{
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
	ctxData, err := ctxFarm.Extract(ctx)
	if err != nil {
		return err
	}

	chicken, err := c.database.GetChicken(ctx, chickenID)
	if err != nil {
		return err
	}

	if chicken.OwnerID != ctxData.FarmerID {
		return ErrChickenNotYours
	}

	if chicken.RestingUntil >= c.currentDay {
		return ErrChickenResting
	}

	if err := c.database.DecrementBarnFeed(
		ctx, chickenID, FeedChickenCost,
	); err != nil {
		return err
	}

	var eggType = EggTypeNormal
	rand.Seed(time.Now().Unix())
	if rand.Intn(100) <= int(chicken.GoldEggChance) {
		eggType = EggTypeGolden
	}

	if err := c.database.IncrementChickenEggLayCount(
		ctx, chickenID, eggType,
	); err != nil {
		return err
	}

	// Must rest at least one day, can rest up to 1 + MaxChickenRestingDays.
	if err := c.database.UpdateChickenRestingUntil(
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
