package farmer

import (
	"chicken-farmer/backend/internal/pkg/database"
	"chicken-farmer/backend/internal/pkg/event"
	"chicken-farmer/backend/internal/pkg/jwt"
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrFarmerNameAlreadyUsed = errors.New("farmer name already used")
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
	validator   *validator.Validate
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
		validator:   validator.New(),
	}
}

func (c *Controller) Register(
	ctx context.Context, farmerName, farmName, password string,
) (Farmer, error) {
	type validate struct {
		FarmerName string `validate:"required"`
		FarmName   string `validate:"required"`
		Password   string `validate:"required"`
	}
	if err := c.validator.Struct(
		validate{
			FarmerName: farmerName,
			FarmName:   farmName,
			Password:   password,
		},
	); err != nil {
		return Farmer{}, err
	}

	_, err := c.datasource.GetFarmerByName(ctx, farmerName)
	if err != nil && err != database.ErrNotFound {
		return Farmer{}, err
	}

	// Error when farmer name already exists, farm name can be duplicate.
	if err == nil {
		return Farmer{}, ErrFarmerNameAlreadyUsed
	}

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
	type validate struct {
		Name     string `validate:"required"`
		Password string `validate:"required"`
	}
	if err := c.validator.Struct(
		validate{
			Name:     name,
			Password: password,
		},
	); err != nil {
		return "", err
	}

	farmer, err := c.datasource.GetFarmerByName(ctx, name)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(farmer.PasswordHash), []byte(password),
	); err != nil {
		return "", err
	}

	token, err := jwt.GenerateUserToken(farmer.ID)
	if err != nil {
		return "", err
	}

	return token, nil
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
	// TODO validate that this is coming from a valid source.

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
	// TODO validate that this is coming from a valid source.

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
