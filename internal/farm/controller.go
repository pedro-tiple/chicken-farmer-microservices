package farm

import (
	"chicken-farmer/backend/internal/pkg/event"
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	PurchaseCostChicken = 1
	PurchaseCostFeedBag = 1
	PurchaseCostBarn    = 10

	FeedUsedPerFeeding = 1
	FeedPerBag         = 10

	MinChickenRestingDays = 5
	MaxChickenRestingDays = 15

	NormalEggsPerFeed = 1
	GoldEggsPerFeed   = 1
)

const (
	EggTypeNormal = iota
	EggTypeGolden
)

var (
	ErrFarmNotYours    = errors.New("farm doesn't belong to you")
	ErrBarnNotYours    = errors.New("barn doesn't belong to you")
	ErrChickenNotYours = errors.New("chicken doesn't belong to you")
	ErrChickenResting  = errors.New("chicken is resting")
)

type IDataSource interface {
	GetFarm(ctx context.Context, farmID uuid.UUID) (Farm, error)
	GetBarnsOfFarm(ctx context.Context, farmID uuid.UUID) ([]Barn, error)
	GetBarn(ctx context.Context, barnID uuid.UUID) (Barn, error)
	GetChickensOfBarn(ctx context.Context, barnID uuid.UUID) ([]Chicken, error)
	GetChicken(ctx context.Context, chickenID uuid.UUID) (Chicken, error)

	InsertFarm(ctx context.Context, farm Farm) (farmID uuid.UUID, err error)
	InsertChicken(
		ctx context.Context, chicken Chicken,
	) (chickenID uuid.UUID, err error)
	InsertBarn(ctx context.Context, barn Barn) (barnID uuid.UUID, err error)

	UpdateChickenRestingUntil(
		ctx context.Context, chickenID uuid.UUID, day uint,
	) error

	IncrementBarnFeed(ctx context.Context, barnID uuid.UUID, amount uint) error
	// DecrementBarnFeedGreaterEqualThan should atomically check that the value
	// of feed egg count is greater or equal than the passed amount.
	// Should return database.ErrNotEnoughFeed when not greater or equal.
	DecrementBarnFeedGreaterEqualThan(
		ctx context.Context, barnID uuid.UUID, amount uint,
	) error

	IncrementChickenEggLayCount(
		ctx context.Context, chickenID uuid.UUID, normalEggCount, goldEggCount int64,
	) error
}

type IFarmerService interface {
	GrantGoldEggs(ctx context.Context, amount uint) error
	SpendGoldEggs(ctx context.Context, amount uint) error
	GetGoldEggs(ctx context.Context) (uint, error)
}

type Controller struct {
	logger        *zap.SugaredLogger
	datasource    IDataSource
	farmerService IFarmerService
	subscriber    message.Subscriber
	publisher     message.Publisher
	currentDay    uint
}

var _ IController = &Controller{}

func ProvideController(
	ctx context.Context,
	logger *zap.SugaredLogger,
	datasource IDataSource,
	farmerService IFarmerService,
	subscriber message.Subscriber,
	publisher message.Publisher,
) (*Controller, error) {
	c := Controller{
		logger:        logger,
		datasource:    datasource,
		farmerService: farmerService,
		subscriber:    subscriber,
		publisher:     publisher,
		currentDay:    1,
	}

	messages, err := subscriber.Subscribe(ctx, event.UniverseTopic)
	if err != nil {
		return nil, err
	}

	go c.processTimeMessages(messages)

	return &c, nil
}

func (c *Controller) processTimeMessages(messages <-chan *message.Message) {
	for msg := range messages {
		var dayMessage event.DayMessage
		if err := json.Unmarshal(msg.Payload, &dayMessage); err != nil {
			// TODO move to DLQ
			c.logger.Error(err)
			continue
		}

		// Ignore outdated messages.
		newDay := dayMessage.Day
		if newDay > c.currentDay {
			c.currentDay = dayMessage.Day
		}

		msg.Ack()
	}
}

func (c *Controller) NewFarm(
	ctx context.Context, ownerID uuid.UUID, name string,
) (uuid.UUID, error) {
	farmID, err := c.datasource.InsertFarm(
		ctx, Farm{
			ID:      uuid.New(),
			OwnerID: ownerID,
			Name:    name,
		},
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return farmID, nil
}

func (c *Controller) FarmDetails(
	ctx context.Context, farmerID, farmID uuid.UUID,
) (FarmDetailsResult, error) {
	farm, err := c.datasource.GetFarm(ctx, farmID)
	if err != nil {
		return FarmDetailsResult{}, err
	}

	if farm.OwnerID != farmerID {
		return FarmDetailsResult{}, ErrFarmNotYours
	}

	barns, err := c.datasource.GetBarnsOfFarm(ctx, farm.ID)
	if err != nil {
		return FarmDetailsResult{}, err
	}

	resultBarns := make([]farmDetailsResultBarn, len(barns))
	errGrp, errGrpCtx := errgroup.WithContext(ctx)
	errGrp.SetLimit(5)

	for i, barn := range barns {
		i, barn := i, barn //nolint:varnamelen

		errGrp.Go(
			func() error {
				chickens, err := c.datasource.GetChickensOfBarn(
					errGrpCtx, barn.ID,
				)
				if err != nil {
					return err
				}
				resultBarns[i] = farmDetailsResultBarn{
					Barn:     barn,
					Chickens: chickens,
				}

				return nil
			},
		)
	}

	if err := errGrp.Wait(); err != nil {
		return FarmDetailsResult{}, err
	}

	// Gold egg count lives in another service for this implementation so must
	// go fetch it there.
	goldEggCount, err := c.farmerService.GetGoldEggs(ctx)
	if err != nil {
		return FarmDetailsResult{}, err
	}

	return FarmDetailsResult{
		Farm:         farm,
		GoldEggCount: goldEggCount,
		CurrentDay:   c.currentDay,
		Barns:        resultBarns,
	}, nil
}

func (c *Controller) BuyBarn(
	ctx context.Context, farmerID, farmID uuid.UUID,
) error {
	farm, err := c.datasource.GetFarm(ctx, farmID)
	if err != nil {
		return err
	}

	if farm.OwnerID != farmerID {
		return ErrBarnNotYours
	}

	if err := c.farmerService.SpendGoldEggs(ctx, PurchaseCostBarn); err != nil {
		return err
	}

	barnID, err := c.datasource.InsertBarn(ctx, Barn{FarmID: farmID})
	if err != nil {
		// TODO unspend gold eggs?
		return err
	}

	if err := event.PublishMessage(
		ctx,
		c.publisher,
		farmerID,
		event.FarmTopic,
		event.MessageTypeNewBarn,
		event.NewBarnMessage{
			ID:            barnID.String(),
			Feed:          0,
			HasAutoFeeder: false,
		},
	); err != nil {
		return err
	}
	return nil
}

func (c *Controller) BuyFeedBags(
	ctx context.Context, farmerID, barnID uuid.UUID, amount uint,
) error {
	barn, err := c.datasource.GetBarn(ctx, barnID)
	if err != nil {
		return err
	}

	if barn.OwnerID != farmerID {
		return ErrBarnNotYours
	}

	if err := c.farmerService.SpendGoldEggs(
		ctx, PurchaseCostFeedBag*amount,
	); err != nil {
		return err
	}

	totalFeed := amount * FeedPerBag
	if err := c.datasource.IncrementBarnFeed(
		ctx, barnID, totalFeed,
	); err != nil {
		return err
	}

	if err := event.PublishMessage(
		ctx,
		c.publisher,
		farmerID,
		event.FarmTopic,
		event.MessageTypeFeedChange,
		event.FeedChangeMessage{
			BarnID: barnID.String(),
			Count:  int(totalFeed),
		},
	); err != nil {
		return err
	}

	return nil
}

func (c *Controller) BuyChicken(
	ctx context.Context, farmerID, barnID uuid.UUID,
) error {
	barn, err := c.datasource.GetBarn(ctx, barnID)
	if err != nil {
		return err
	}

	if barn.OwnerID != farmerID {
		return ErrBarnNotYours
	}

	if err := c.farmerService.SpendGoldEggs(
		ctx, PurchaseCostChicken,
	); err != nil {
		return err
	}

	// Maybe have the gold egg chance be on a normal distribution?
	rand.Seed(int64(time.Now().Nanosecond()))

	chickenID := uuid.New()
	_, err = c.datasource.InsertChicken(
		ctx, Chicken{
			ID:            chickenID,
			BarnID:        barnID,
			DateOfBirth:   c.currentDay,
			GoldEggChance: uint(rand.Intn(99) + 1), // [1,100]
		},
	)
	if err != nil {
		return err
	}

	if err := event.PublishMessage(
		ctx,
		c.publisher,
		farmerID,
		event.FarmTopic,
		event.MessageTypeNewChicken,
		event.NewChickenMessage{
			BarnID:      barn.ID.String(),
			ChickenID:   chickenID.String(),
			DateOfBirth: c.currentDay,
		},
	); err != nil {
		return err
	}

	return nil
}

func (c *Controller) FeedChicken(
	ctx context.Context, farmerID, chickenID uuid.UUID,
) error {
	chicken, err := c.datasource.GetChicken(ctx, chickenID)
	if err != nil {
		return err
	}

	if chicken.OwnerID != farmerID {
		return ErrChickenNotYours
	}

	if chicken.RestingUntil >= c.currentDay {
		return ErrChickenResting
	}

	if err := c.datasource.DecrementBarnFeedGreaterEqualThan(
		ctx, chicken.BarnID, FeedUsedPerFeeding,
	); err != nil {
		return err
	}

	rand.Seed(int64(time.Now().Nanosecond()))

	var normalEggsLaid, goldEggsLaid int64
	if rand.Intn(100) <= int(chicken.GoldEggChance) {
		normalEggsLaid = GoldEggsPerFeed
	} else {
		goldEggsLaid = NormalEggsPerFeed
	}

	g, errGrpCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if goldEggsLaid == 0 {
			return nil
		}

		return c.farmerService.GrantGoldEggs(
			errGrpCtx, GoldEggsPerFeed,
		)
	})

	g.Go(func() error {
		return c.datasource.IncrementChickenEggLayCount(
			errGrpCtx, chickenID, normalEggsLaid, goldEggsLaid,
		)
	})

	g.Go(func() error {
		// Must rest at least one day, can rest up to 1 + MaxChickenRestingDays.
		chicken.RestingUntil = c.currentDay + MinChickenRestingDays +
			uint(rand.Intn(MaxChickenRestingDays-MinChickenRestingDays))

		return c.datasource.UpdateChickenRestingUntil(
			errGrpCtx, chickenID, chicken.RestingUntil,
		)
	})

	g.Go(func() error {
		return event.PublishMessage(
			errGrpCtx,
			c.publisher,
			farmerID,
			event.FarmTopic,
			event.MessageTypeChickenFed,
			event.ChickenFedMessage{
				ChickenID:      chickenID.String(),
				RestingUntil:   chicken.RestingUntil,
				NormalEggsLaid: uint(normalEggsLaid),
				GoldEggsLaid:   uint(goldEggsLaid),
			},
		)
	})

	g.Go(func() error {
		return event.PublishMessage(
			errGrpCtx,
			c.publisher,
			farmerID,
			event.FarmTopic,
			event.MessageTypeFeedChange,
			event.FeedChangeMessage{
				BarnID: chicken.BarnID.String(),
				Count:  -FeedUsedPerFeeding,
			},
		)
	})

	return g.Wait()
}

func (c *Controller) FeedChickensOfBarn(
	ctx context.Context, farmerID, barnID uuid.UUID,
) error {
	// TODO
	return nil
}

func (c *Controller) SetDay(ctx context.Context, day uint) error {
	c.currentDay = day

	return nil
}

func (c *Controller) GetCurrentDay(ctx context.Context) (uint, error) {
	return c.currentDay, nil
}
