package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ptiple/farmersvc"
)

type MongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

func New(_ctx context.Context) (MongoDatabase, error) {
	var mongodb MongoDatabase

	// TODO use env files
	// TODO remove hardcoded authentication
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@farmer-svc-mongodb:27017")
	client, err := mongo.Connect(_ctx, clientOptions)
	if err != nil {
		return mongodb, err
	}

	mongodb = MongoDatabase{
		client,
		client.Database("farmer-svc"),
		_ctx,
	}

	return mongodb, nil
}

func (mongodb MongoDatabase) InsertFarmer(_farmer *farmersvc.Farmer) (*farmersvc.Farmer, error) {
	id, err := mongodb.db.Collection("farmers").InsertOne(mongodb.ctx, _farmer)
	if err != nil {
		return _farmer, err
	}

	_farmer.Id = id.InsertedID.(primitive.ObjectID)

	return _farmer, nil
}

func (mongodb MongoDatabase) RemoveFarmer(_id primitive.ObjectID) error {
	filter := bson.M{"_id": _id}
	_, err := mongodb.db.Collection("farmers").DeleteOne(mongodb.ctx, filter)

	return err
}

func (mongodb MongoDatabase) UpdateFarmer(_farmer farmersvc.Farmer) error {
	filter := bson.M{"_id": _farmer.Id}
	update := bson.D{{"$set", _farmer}}
	_, err := mongodb.db.Collection("farmers").UpdateOne(mongodb.ctx, filter, update)

	return err
}

func (mongodb MongoDatabase) GetFarmer(_id primitive.ObjectID) (*farmersvc.Farmer, error) {
	var farmer *farmersvc.Farmer

	filter := bson.M{"_id": _id}
	err := mongodb.db.Collection("farmers").FindOne(mongodb.ctx, filter).Decode(&farmer)

	return farmer, err
}
