package sql

import (
	farmPkg "chicken-farmer/backend/internal/farm"
	internalDB "chicken-farmer/backend/internal/pkg/database"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidEggType = errors.New("invalid egg type")
)

type Datasource struct {
	database *Queries
}

var _ farmPkg.IDataSource = &Datasource{}

func ProvideDatasource(dbConnection *sql.DB) (*Datasource, error) {
	// TODO use multiple connections
	return &Datasource{
		database: New(dbConnection),
	}, nil
}

func (d *Datasource) GetFarm(
	ctx context.Context, farmID uuid.UUID,
) (farmPkg.Farm, error) {
	farm, err := d.database.GetFarm(ctx, farmID)
	if err != nil {
		return farmPkg.Farm{}, internalDB.NormalizeNotFound(err)
	}

	return farmPkg.Farm{
		ID:      farm.ID,
		OwnerID: farm.OwnerID,
		Name:    farm.Name,
	}, nil
}

func (d *Datasource) GetBarn(
	ctx context.Context, barnID uuid.UUID,
) (farmPkg.Barn, error) {
	barn, err := d.database.GetBarn(ctx, barnID)
	if err != nil {
		return farmPkg.Barn{}, internalDB.NormalizeNotFound(err)
	}

	return farmPkg.Barn{
		ID:            barn.ID,
		OwnerID:       barn.OwnerID,
		FarmID:        barn.FarmID,
		Feed:          uint(barn.Feed),
		HasAutoFeeder: barn.HasAutoFeeder,
	}, nil
}

func (d *Datasource) GetBarnsOfFarm(
	ctx context.Context, farmID uuid.UUID,
) ([]farmPkg.Barn, error) {
	barns, err := d.database.GetBarnsOfFarm(ctx, farmID)
	if err != nil {
		return nil, internalDB.NormalizeNotFound(err)
	}

	result := make([]farmPkg.Barn, len(barns))
	for i, barn := range barns {
		result[i] = farmPkg.Barn{
			ID:            barn.ID,
			OwnerID:       barn.OwnerID,
			FarmID:        barn.FarmID,
			Feed:          uint(barn.Feed),
			HasAutoFeeder: barn.HasAutoFeeder,
		}
	}

	return result, nil
}

func (d *Datasource) GetChickensOfBarn(
	ctx context.Context, barnID uuid.UUID,
) ([]farmPkg.Chicken, error) {
	chickens, err := d.database.GetChickensOfBarn(ctx, barnID)
	if err != nil {
		return nil, internalDB.NormalizeNotFound(err)
	}

	result := make([]farmPkg.Chicken, len(chickens))
	for i, chicken := range chickens {
		result[i] = farmPkg.Chicken{
			ID:             chicken.ID,
			OwnerID:        chicken.OwnerID,
			BarnID:         chicken.BarnID,
			DateOfBirth:    uint(chicken.DateOfBirth),
			RestingUntil:   uint(chicken.RestingUntil),
			NormalEggsLaid: uint(chicken.NormalEggsLaid),
			GoldEggsLaid:   uint(chicken.GoldEggsLaid),
			GoldEggChance:  uint(chicken.GoldEggChance),
		}
	}

	return result, nil
}

func (d *Datasource) GetChicken(
	ctx context.Context, chickenID uuid.UUID,
) (farmPkg.Chicken, error) {
	chicken, err := d.database.GetChicken(ctx, chickenID)
	if err != nil {
		return farmPkg.Chicken{}, internalDB.NormalizeNotFound(err)
	}

	return farmPkg.Chicken{
		ID:             chicken.ID,
		OwnerID:        chicken.OwnerID,
		BarnID:         chicken.BarnID,
		DateOfBirth:    uint(chicken.DateOfBirth),
		RestingUntil:   uint(chicken.RestingUntil),
		NormalEggsLaid: uint(chicken.NormalEggsLaid),
		GoldEggsLaid:   uint(chicken.GoldEggsLaid),
		GoldEggChance:  uint(chicken.GoldEggChance),
	}, nil
}

func (d *Datasource) InsertFarm(
	ctx context.Context, farm farmPkg.Farm,
) (uuid.UUID, error) {
	insertedChicken, err := d.database.InsertFarm(
		ctx, InsertFarmParams{
			ID:      farm.ID,
			OwnerID: farm.OwnerID,
			Name:    farm.Name,
		},
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return insertedChicken.ID, nil
}

func (d *Datasource) InsertChicken(
	ctx context.Context, chicken farmPkg.Chicken,
) (uuid.UUID, error) {
	insertedChicken, err := d.database.InsertChicken(
		ctx, InsertChickenParams{
			ID:             chicken.ID,
			DateOfBirth:    int64(chicken.DateOfBirth),
			RestingUntil:   int64(chicken.RestingUntil),
			NormalEggsLaid: int64(chicken.NormalEggsLaid),
			GoldEggsLaid:   int64(chicken.GoldEggsLaid),
			GoldEggChance:  int64(chicken.GoldEggChance),
			BarnID:         chicken.BarnID,
		},
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return insertedChicken.ID, nil
}

func (d *Datasource) InsertBarn(
	ctx context.Context, barn farmPkg.Barn,
) (uuid.UUID, error) {
	insertedBarn, err := d.database.InsertBarn(
		ctx, InsertBarnParams{
			FarmID:        barn.FarmID,
			Feed:          int64(barn.Feed),
			HasAutoFeeder: barn.HasAutoFeeder,
		},
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return insertedBarn.ID, nil
}

func (d *Datasource) UpdateChickenRestingUntil(
	ctx context.Context, chickenID uuid.UUID, day uint,
) error {
	return d.database.UpdateChickenRestingUntil(
		ctx, UpdateChickenRestingUntilParams{
			ID:           chickenID,
			RestingUntil: int64(day),
		},
	)
}

func (d *Datasource) IncrementBarnFeed(
	ctx context.Context, barnID uuid.UUID, amount uint,
) error {
	return d.database.IncrementBarnFeed(
		ctx, IncrementBarnFeedParams{
			ID:   barnID,
			Feed: int64(amount),
		},
	)
}

func (d *Datasource) DecrementBarnFeed(
	ctx context.Context, barnID uuid.UUID, amount uint,
) error {
	return d.database.DecrementBarnFeed(
		ctx, DecrementBarnFeedParams{
			ID:   barnID,
			Feed: int64(amount),
		},
	)
}

func (d *Datasource) IncrementChickenEggLayCount(
	ctx context.Context, chickenID uuid.UUID, eggType int,
) error {
	// TODO this is business logic and shouldn't be here.
	switch eggType {
	case farmPkg.EggTypeGolden:
		return d.database.IncrementChickenGoldEggLayCount(ctx, chickenID)
	case farmPkg.EggTypeNormal:
		return d.database.IncrementChickenNormalEggLayCount(ctx, chickenID)
	}

	return ErrInvalidEggType
}
