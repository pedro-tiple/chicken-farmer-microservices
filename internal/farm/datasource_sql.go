package farm

import (
	farmSQL "chicken-farmer/backend/internal/farm/sql"
	internalDB "chicken-farmer/backend/internal/pkg/database"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type SQLDatabase struct {
	database *farmSQL.Queries
}

var _ IDataSource = &SQLDatabase{}

func ProvideSQLDatabase(dbConnection *sql.DB) (*SQLDatabase, error) {
	// TODO use multiple connections

	return &SQLDatabase{
		database: farmSQL.New(dbConnection),
	}, nil
}

func (d *SQLDatabase) GetFarm(
	ctx context.Context, farmID uuid.UUID,
) (Farm, error) {
	farm, err := d.database.GetFarm(ctx, farmID)
	if err != nil {
		return Farm{}, internalDB.NormalizeNotFound(err)
	}

	return Farm{
		ID:      farm.ID,
		OwnerID: farm.OwnerID,
		Name:    farm.Name,
	}, nil
}

func (d *SQLDatabase) GetBarnsOfFarm(
	ctx context.Context, farmID uuid.UUID,
) ([]Barn, error) {
	barns, err := d.database.GetBarnsOfFarm(ctx, farmID)
	if err != nil {
		return nil, internalDB.NormalizeNotFound(err)
	}

	result := make([]Barn, len(barns))
	for i, barn := range barns {
		result[i] = Barn{
			ID:            barn.ID,
			OwnerID:       barn.OwnerID,
			FarmID:        barn.FarmID,
			Feed:          uint(barn.Feed),
			HasAutoFeeder: barn.HasAutoFeeder,
		}
	}

	return result, nil
}

func (d *SQLDatabase) GetChickensOfBarn(
	ctx context.Context, barnID uuid.UUID,
) ([]Chicken, error) {
	chickens, err := d.database.GetChickensOfBarn(ctx, barnID)
	if err != nil {
		return nil, internalDB.NormalizeNotFound(err)
	}

	result := make([]Chicken, len(chickens))
	for i, chicken := range chickens {
		result[i] = Chicken{
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

func (d *SQLDatabase) GetChicken(
	ctx context.Context, chickenID uuid.UUID,
) (Chicken, error) {
	chicken, err := d.database.GetChicken(ctx, chickenID)
	if err != nil {
		return Chicken{}, internalDB.NormalizeNotFound(err)
	}

	return Chicken{
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

func (d *SQLDatabase) InsertChicken(
	ctx context.Context, chicken Chicken,
) (chickenID uuid.UUID, err error) {
	insertedChicken, err := d.database.InsertChicken(
		ctx, farmSQL.InsertChickenParams{
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

func (d *SQLDatabase) InsertBarn(
	ctx context.Context, barn Barn,
) (uuid.UUID, error) {
	insertedBarn, err := d.database.InsertBarn(ctx, farmSQL.InsertBarnParams{
		FarmID:        barn.FarmID,
		Feed:          int64(barn.Feed),
		HasAutoFeeder: barn.HasAutoFeeder,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return insertedBarn.ID, nil
}

func (d *SQLDatabase) UpdateChickenRestingUntil(
	ctx context.Context, chickenID uuid.UUID, day uint,
) error {
	return d.database.UpdateChickenRestingUntil(ctx, farmSQL.UpdateChickenRestingUntilParams{
		ID:           chickenID,
		RestingUntil: int64(day),
	})
}

func (d *SQLDatabase) IncrementBarnFeed(
	ctx context.Context, barnID uuid.UUID, amount uint,
) error {
	return d.database.IncrementBarnFeed(ctx, farmSQL.IncrementBarnFeedParams{
		ID:   barnID,
		Feed: int64(amount),
	})
}

func (d *SQLDatabase) DecrementBarnFeed(
	ctx context.Context, barnID uuid.UUID, amount uint,
) error {
	return d.database.DecrementBarnFeed(ctx, farmSQL.DecrementBarnFeedParams{
		ID:   barnID,
		Feed: int64(amount),
	})
}

func (d *SQLDatabase) IncrementChickenEggLayCount(
	ctx context.Context, chickenID uuid.UUID, eggType int,
) error {
	switch eggType {
	case EggTypeGolden:
		return d.database.IncrementChickenGoldEggLayCount(ctx, chickenID)
	case EggTypeNormal:
		return d.database.IncrementChickenNormalEggLayCount(ctx, chickenID)
	}
	return ErrInvalidEggType
}
