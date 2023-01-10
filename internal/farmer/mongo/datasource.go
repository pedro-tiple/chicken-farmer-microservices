package mongo

import (
	farmerPkg "chicken-farmer/backend/internal/farmer"
	"chicken-farmer/backend/internal/pkg"
	internalDatabase "chicken-farmer/backend/internal/pkg/database"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	database          = "farmer-service"
	farmersCollection = "farmers"
)

type Datasource struct {
	client *mongo.Client
	db     *mongo.Database
}

var _ farmerPkg.IDataSource = &Datasource{}

func ProvideDatasource(ctx context.Context) (*Datasource, error) {
	// TODO use env files
	// TODO remove hardcoded authentication
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(
			"mongodb://admin:password@localhost:27017",
		),
	)
	if err != nil {
		return nil, err
	}

	// Ping to make sure the connection is open.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &Datasource{
		client,
		client.Database(database),
	}, nil
}

func (d Datasource) GetFarmer(
	ctx context.Context, farmerID uuid.UUID,
) (farmerPkg.Farmer, error) {
	var farmer Farmer
	if err := d.db.
		Collection(farmersCollection).
		FindOne(ctx, bson.D{{"uuid", farmerID.String()}}).
		Decode(&farmer); err != nil {
		return farmerPkg.Farmer{}, err
	}

	return farmerPkg.Farmer{
		ID:           pkg.UUIDFromString(farmer.UUID),
		FarmID:       pkg.UUIDFromString(farmer.FarmID),
		Name:         farmer.Name,
		PasswordHash: farmer.PasswordHash,
		GoldEggCount: farmer.GoldEggCount,
	}, nil
}

func (d Datasource) GetFarmerByName(
	ctx context.Context, name string,
) (farmerPkg.Farmer, error) {
	var farmer Farmer
	if err := d.db.
		Collection(farmersCollection).
		FindOne(ctx, bson.D{{"name", name}}).
		Decode(&farmer); err != nil {
		return farmerPkg.Farmer{}, err
	}

	return farmerPkg.Farmer{
		ID:           pkg.UUIDFromString(farmer.UUID),
		FarmID:       pkg.UUIDFromString(farmer.FarmID),
		Name:         farmer.Name,
		PasswordHash: farmer.PasswordHash,
		GoldEggCount: farmer.GoldEggCount,
	}, nil
}

func (d Datasource) GetFarmerGoldEggs(
	ctx context.Context, farmerID uuid.UUID,
) (uint, error) {
	var farmer Farmer
	if err := d.db.
		Collection(farmersCollection).
		FindOne(
			ctx,
			bson.D{{"uuid", farmerID.String()}},
			options.FindOne().SetProjection(
				bson.D{{
					"goldEggCount", 1,
				}},
			),
		).
		Decode(&farmer); err != nil {
		return 0, err
	}

	return farmer.GoldEggCount, nil
}

func (d Datasource) InsertFarmer(
	ctx context.Context, farmer farmerPkg.Farmer,
) (uuid.UUID, error) {
	if farmer.ID.String() == (uuid.UUID{}).String() {
		farmer.ID = uuid.New()
	}

	_, err := d.db.
		Collection(farmersCollection).
		InsertOne(
			ctx, Farmer{
				UUID:         farmer.ID.String(),
				FarmID:       farmer.FarmID.String(),
				Name:         farmer.Name,
				PasswordHash: farmer.PasswordHash,
				GoldEggCount: farmer.GoldEggCount,
			},
		)
	if err != nil {
		return uuid.UUID{}, err
	}

	return farmer.FarmID, nil
}

func (d Datasource) DecrementGoldEggCountGreaterEqualThan(
	ctx context.Context, farmerID uuid.UUID, amount uint,
) error {
	result, err := d.db.
		Collection(farmersCollection).
		UpdateOne(
			ctx,
			bson.D{
				{"uuid", farmerID.String()},
				{"goldEggCount", bson.D{{"$gte", amount}}},
			},
			// bson.D{{"$set", bson.D{{"test", "test"}}}},
			bson.D{{"$inc", bson.D{{"goldEggCount", -int(amount)}}}},
		)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return internalDatabase.ErrNotEnoughGoldEggs
	}

	return nil
}
