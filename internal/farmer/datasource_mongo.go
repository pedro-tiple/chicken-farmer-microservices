package farmer

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
}

var _ IDataSource = &MongoDatabase{}

func ProvideMongoDatabase(ctx context.Context) (*MongoDatabase, error) {
	// TODO use env files
	// TODO remove hardcoded authentication
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://admin:password@barnsvc-mongodb:27017",
	))
	if err != nil {
		return nil, err
	}

	return &MongoDatabase{
		client,
		client.Database("barnsvc"),
	}, nil
}

func (m MongoDatabase) GetFarmer(ctx context.Context, farmerID uuid.UUID) (Farmer, error) {
	// TODO implement me
	panic("implement me")
}

func (m MongoDatabase) InsertFarmer(ctx context.Context, farmer Farmer) (uuid.UUID, error) {
	// TODO implement me
	panic("implement me")
}

/*
func (mongodb *MongoDatabase) InsertBarn(_barn *barnsvc.Barn) (*barnsvc.Barn, error) {
	id, err := mongodb.db.Collection("barns").InsertOne(mongodb.ctx, _barn)
	if err != nil {
		return _barn, err
	}

	_barn.Id = id.InsertedID.(primitive.ObjectID)

	return _barn, nil
}

func (mongodb *MongoDatabase) RemoveBarn(_id primitive.ObjectID) error {
	filter := bson.M{"_id": _id}
	_, err := mongodb.db.Collection("barns").DeleteOne(mongodb.ctx, filter)

	return err
}

func (mongodb *MongoDatabase) UpdateBarn(_barn barnsvc.Barn) error {
	filter := bson.M{"_id": _barn.Id}
	update := bson.D{{"$set", _barn}}
	_, err := mongodb.db.Collection("barns").UpdateOne(mongodb.ctx, filter, update)

	return err
}

func (mongodb *MongoDatabase) GetBarn(_id primitive.ObjectID) (*barnsvc.Barn, error) {
	var barn *barnsvc.Barn

	filter := bson.M{"_id": _id}
	err := mongodb.db.Collection("barns").FindOne(mongodb.ctx, filter).Decode(&barn)

	return barn, err
}

func (mongodb *MongoDatabase) GetBarnsOfFarmer(_farmerId primitive.ObjectID) ([]*barnsvc.Barn, error) {
	// must be this way so it returns [] instead of nil
	var barns = []*barnsvc.Barn{}

	filter := bson.M{"belongsToFarmer": _farmerId}
	opts := options.Find().SetSort(bson.D{{"purchaseDay", 1}})
	cursor, err := mongodb.db.Collection("barns").Find(mongodb.ctx, filter, opts)
	if err != nil {
		return barns, err
	}

	for cursor.Next(mongodb.ctx) {
		// create a value into which the single document can be decoded
		var barn barnsvc.Barn
		err := cursor.Decode(&barn)
		if err != nil {
			continue
		}
		barns = append(barns, &barn)
	}

	return barns, err
}
*/
