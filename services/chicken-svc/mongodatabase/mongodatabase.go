package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	chickensvc "ptiple/chicken-svc"
)

type MongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

type SystemVariable struct {
	Name  string `json:"name"`
	Value uint   `json:"value,string"`
}

const databaseName = "chicken-svc"
const chickenCollectionName = "chickens"
const variablesCollectionName = "variables"

func New(_ctx context.Context) (MongoDatabase, error) {
	var mongodb MongoDatabase

	// TODO use env files
	// TODO remove hardcoded authentication
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@chicken-svc-mongodb:27017")
	client, err := mongo.Connect(_ctx, clientOptions)
	if err != nil {
		return mongodb, err
	}

	mongodb = MongoDatabase{
		client,
		client.Database(databaseName),
		_ctx,
	}

	return mongodb, nil
}

func (mongodb MongoDatabase) InsertChicken(_chicken *chickensvc.Chicken) (*chickensvc.Chicken, error) {
	id, err := mongodb.db.Collection(chickenCollectionName).InsertOne(mongodb.ctx, _chicken)
	if err != nil {
		return _chicken, err
	}

	_chicken.Id = id.InsertedID.(primitive.ObjectID)

	return _chicken, nil
}

func (mongodb MongoDatabase) RemoveChicken(_id primitive.ObjectID) error {
	filter := bson.M{"_id": _id}
	_, err := mongodb.db.Collection(chickenCollectionName).DeleteOne(mongodb.ctx, filter)

	return err
}

func (mongodb MongoDatabase) UpdateChicken(_chicken chickensvc.Chicken) error {
	filter := bson.M{"_id": _chicken.Id}
	update := bson.D{{"$set", _chicken}}
	_, err := mongodb.db.Collection(chickenCollectionName).UpdateOne(mongodb.ctx, filter, update)

	return err
}

func (mongodb MongoDatabase) GetChicken(_id primitive.ObjectID) (*chickensvc.Chicken, error) {
	var chicken *chickensvc.Chicken

	filter := bson.M{"_id": _id}
	err := mongodb.db.Collection(chickenCollectionName).FindOne(mongodb.ctx, filter).Decode(&chicken)

	return chicken, err
}

func (mongodb MongoDatabase) GetChickensOfUser(_userId primitive.ObjectID) ([]*chickensvc.Chicken, error) {
	// must be this way so it returns [] instead of nil
	var chickens = []*chickensvc.Chicken{}

	filter := bson.M{"belongsToUser": _userId}
	opts := options.Find().SetSort(bson.D{{"purchaseDay", 1}})
	cursor, err := mongodb.db.Collection(chickenCollectionName).Find(mongodb.ctx, filter, opts)
	if err != nil {
		return chickens, err
	}

	for cursor.Next(mongodb.ctx) {
		// create a value into which the single document can be decoded
		var chicken chickensvc.Chicken
		err := cursor.Decode(&chicken)
		if err != nil {
			continue
		}
		chickens = append(chickens, &chicken)
	}

	return chickens, err
}

func (mongodb MongoDatabase) UpdateDay(_day uint) error {
	sysvar := SystemVariable{
		Name:  "day",
		Value: _day,
	}
	filter := bson.M{"name": "day"}
	update := bson.D{{"$set", sysvar}}
	updateOptions := options.Update().SetUpsert(true)
	_, err := mongodb.db.Collection(variablesCollectionName).UpdateOne(mongodb.ctx, filter, update, updateOptions)

	return err
}

func (mongodb MongoDatabase) GetDay() (uint, error) {
	var day SystemVariable

	filter := bson.M{"name": "day"}
	err := mongodb.db.Collection(variablesCollectionName).FindOne(mongodb.ctx, filter).Decode(&day)

	return day.Value, err
}
