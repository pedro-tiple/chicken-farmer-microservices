package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	chickensvc "ptiple/chicken-svc"
	"strings"
)

type MongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

func New(_ctx context.Context) (MongoDatabase, error) {
	var mongodb MongoDatabase

	// TODO set this back when building, or use env files, remove hardcoded authentication
	// clientOptions := options.Client().ApplyURI("mongodb://admin:password@chicken-svc-mongodb:27017")
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@192.168.99.100:30335")
	client, err := mongo.Connect(_ctx, clientOptions)
	if err != nil {
		return mongodb, err
	}

	mongodb = MongoDatabase{
		client,
		client.Database("chicken-svc"),
		_ctx,
	}

	return mongodb, nil
}

func (mongodb MongoDatabase) InsertChicken(_chicken *chickensvc.Chicken) (*chickensvc.Chicken, error) {
	id, err := mongodb.db.Collection("chickens").InsertOne(mongodb.ctx, _chicken)
	if err != nil {
		return _chicken, err
	}

	_chicken.Id = id.InsertedID.(primitive.ObjectID)

	return _chicken, nil
}

func (mongodb MongoDatabase) RemoveChicken(_id primitive.ObjectID) error {
	filter := bson.M{"_id": _id}
	_, err := mongodb.db.Collection("chickens").DeleteOne(mongodb.ctx, filter)

	return err
}

func (mongodb MongoDatabase) UpdateChicken(_chicken chickensvc.Chicken) error {
	filter := bson.M{"_id": _chicken.Id}
	update := bson.D{{"$set", _chicken}}
	_, err := mongodb.db.Collection("chickens").UpdateOne(mongodb.ctx, filter, update)

	return err
}

func (mongodb MongoDatabase) GetChicken(_id primitive.ObjectID) (*chickensvc.Chicken, error) {
	var chicken *chickensvc.Chicken

	filter := bson.M{"_id": _id}
	err := mongodb.db.Collection("chickens").FindOne(mongodb.ctx, filter).Decode(&chicken)

	return chicken, err
}

func (mongodb MongoDatabase) GetChickensOfBarn(_barnId primitive.ObjectID) ([]*chickensvc.Chicken, error) {
	// must be this way so it returns [] instead of nil
	var chickens = []*chickensvc.Chicken{}

	filter := bson.M{strings.ToLower("BelongsToBarn"): _barnId}
	opts := options.Find().SetSort(bson.D{{strings.ToLower("BirthDay"), 1}})
	cursor, err := mongodb.db.Collection("chickens").Find(mongodb.ctx, filter, opts)
	if err != nil {
		return chickens, err
	}

	for cursor.Next(mongodb.ctx) {
		// create a value into which the single document can be decoded
		var chicken chickensvc.Chicken
		err := cursor.Decode(&chicken)
		if err != nil {
			log.Println("[GetChickenOfBarn] error decoding chicken", err)
			continue
		}
		chickens = append(chickens, &chicken)
	}

	return chickens, err
}
