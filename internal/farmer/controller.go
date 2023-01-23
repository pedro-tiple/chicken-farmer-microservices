package farmer

import (
	"chicken-farmer/backend/internal/pkg/event"
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IDataSource interface {
	GetFarmer(ctx context.Context, farmerID uuid.UUID) (Farmer, error)
	GetFarmerByName(ctx context.Context, name string) (Farmer, error)
	GetFarmerGoldEggs(ctx context.Context, farmerID uuid.UUID) (uint, error)

	InsertFarmer(
		ctx context.Context, farmer Farmer,
	) (farmerID uuid.UUID, err error)

	IncrementGoldEggCount(
		ctx context.Context, farmerID uuid.UUID, amount uint,
	) error

	// DecrementGoldEggCountGreaterEqualThan should atomically check that the value
	// of gold egg count is greater or equal than the passed amount.
	// Should return database.ErrNotEnoughGoldEggs when not greater or equal.
	DecrementGoldEggCountGreaterEqualThan(
		ctx context.Context, farmerID uuid.UUID, amount uint,
	) error
}

type IFarmService interface {
	NewFarm(
		ctx context.Context, ownerID uuid.UUID, name string,
	) (farmID uuid.UUID, err error)
	DeleteFarm(ctx context.Context, farmID uuid.UUID) error
}

type Controller struct {
	logger      *zap.SugaredLogger
	datasource  IDataSource
	farmService IFarmService
	publisher   message.Publisher
}

var _ IController = &Controller{}

func ProvideController(
	logger *zap.SugaredLogger,
	datasource IDataSource,
	farmService IFarmService,
	publisher message.Publisher,
) *Controller {
	return &Controller{
		logger:      logger,
		datasource:  datasource,
		farmService: farmService,
		publisher:   publisher,
	}
}

func (c *Controller) Register(
	ctx context.Context, farmerName, farmName, password string,
) (Farmer, error) {
	farmerID := uuid.New()
	farmID, err := c.farmService.NewFarm(ctx, farmerID, farmName)
	if err != nil {
		return Farmer{}, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return Farmer{}, err
	}

	farmer := Farmer{
		ID:           farmerID,
		Name:         farmerName,
		FarmID:       farmID,
		PasswordHash: string(passwordHash),
		GoldEggCount: 1,
	}
	_, err = c.datasource.InsertFarmer(ctx, farmer)
	if err != nil {
		// Delete the farm as it would be left dangling uselessly.
		if delErr := c.farmService.DeleteFarm(ctx, farmID); delErr != nil {
			c.logger.Errorw(
				"deleting dangling farm",
				"farmID", farmID,
				"err", err,
				"delErr", err,
			)
		}

		return Farmer{}, err
	}

	return farmer, nil
}

func (c *Controller) Login(
	ctx context.Context, name, password string,
) (string, error) {
	farmer, err := c.datasource.GetFarmerByName(ctx, name)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(farmer.PasswordHash), []byte(password),
	); err != nil {
		return "", err
	}

	var jwt string
	// TODO generate JWT

	return jwt, nil
}

func (c *Controller) GetGoldEggs(
	ctx context.Context, farmerID uuid.UUID,
) (uint, error) {
	farmer, err := c.datasource.GetFarmer(ctx, farmerID)
	if err != nil {
		return 0, err
	}

	return farmer.GoldEggCount, nil
}

func (c *Controller) GrantGoldEggs(
	ctx context.Context, farmerID uuid.UUID, amount uint,
) error {
	if err := c.datasource.IncrementGoldEggCount(
		ctx, farmerID, amount,
	); err != nil {
		return err
	}

	if err := event.PublishMessage(
		ctx,
		c.publisher,
		farmerID,
		event.FarmTopic,
		event.MessageTypeGoldenEggsChange,
		event.GoldenEggChangeMessage{
			Count: int(amount),
		},
	); err != nil {
		return err
	}

	return nil
}

func (c *Controller) SpendGoldEggs(
	ctx context.Context, farmerID uuid.UUID, amount uint,
) error {
	if err := c.datasource.DecrementGoldEggCountGreaterEqualThan(
		ctx, farmerID, amount,
	); err != nil {
		return err
	}

	if err := event.PublishMessage(
		ctx,
		c.publisher,
		farmerID,
		event.FarmTopic,
		event.MessageTypeGoldenEggsChange,
		event.GoldenEggChangeMessage{
			Count: -int(amount),
		},
	); err != nil {
		return err
	}
	return nil
}
