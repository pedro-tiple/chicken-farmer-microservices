package farmer

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IDataSource interface {
	GetFarmer(ctx context.Context, farmerID uuid.UUID) (Farmer, error)
	GetFarmerByName(ctx context.Context, name string) (Farmer, error)

	InsertFarmer(ctx context.Context, farmer Farmer) (farmerID uuid.UUID, err error)
}

type IFarmService interface {
	NewFarm(ctx context.Context, ownerID uuid.UUID, name string) (farmID uuid.UUID, err error)
}

type Controller struct {
	datasource  IDataSource
	farmService IFarmService
	currentDay  uint
	logger      *zap.SugaredLogger
}

var _ IController = &Controller{}

func ProvideController(
	datasource IDataSource,
	farmService IFarmService,
	logger *zap.SugaredLogger,
) *Controller {
	return &Controller{
		datasource:  datasource,
		farmService: farmService,
		currentDay:  0,
		logger:      logger,
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
	}
	_, err = c.datasource.InsertFarmer(ctx, farmer)
	if err != nil {
		c.logger.Errorw(
			"should delete dangling farm",
			"err", err,
			"farmID", farmID,
		)
		// TODO delete the farm as it's gonna be left dangling uselessly
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
	// TODO implement me
	return 10, nil
}

func (c *Controller) SpendGoldEggs(
	ctx context.Context, farmerID uuid.UUID, amount uint,
) error {
	// TODO implement me
	return nil
}
