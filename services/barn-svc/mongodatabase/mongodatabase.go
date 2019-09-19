package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	barnsvc "ptiple/barn-svc"
)

type IMongoDatabase interface {
	InsertBarn(_barn *barnsvc.Barn) (*barnsvc.Barn, error)
	RemoveBarn(_id primitive.ObjectID) error
	UpdateBarn(_barn barnsvc.Barn) error
	GetBarn(_id primitive.ObjectID) (*barnsvc.Barn, error)
	GetBarnsOfUser(_userId primitive.ObjectID) ([]*barnsvc.Barn, error)
}

type MongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

func New(_ctx context.Context) (IMongoDatabase, error) {
	var mongodb IMongoDatabase

	// TODO use env files
	// TODO remove hardcoded authentication
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@barn-svc-mongodb:27017")
	client, err := mongo.Connect(_ctx, clientOptions)
	if err != nil {
		return mongodb, err
	}

	mongodb = &MongoDatabase{
		client,
		client.Database("barn-svc"),
		_ctx,
	}

	return mongodb, nil
}

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

func (mongodb *MongoDatabase) GetBarnsOfUser(_userId primitive.ObjectID) ([]*barnsvc.Barn, error) {
	// must be this way so it returns [] instead of nil
	var barns = []*barnsvc.Barn{}

	filter := bson.M{"belongsToUser": _userId}
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
