package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	usersvc "ptiple/user-svc"
)

type MongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

func New(_ctx context.Context) (MongoDatabase, error) {
	var mongodb MongoDatabase

	// TODO set this back when building, or use env files
	// TODO remove hardcoded authentication
	// clientOptions := options.Client().ApplyURI("mongodb://admin:password@user-svc-mongodb:27017")
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@192.168.99.100:31412")
	client, err := mongo.Connect(_ctx, clientOptions)
	if err != nil {
		return mongodb, err
	}

	mongodb = MongoDatabase{
		client,
		client.Database("user-svc"),
		_ctx,
	}

	return mongodb, nil
}

func (mongodb MongoDatabase) InsertUser(_user *usersvc.User) (*usersvc.User, error) {
	id, err := mongodb.db.Collection("users").InsertOne(mongodb.ctx, _user)
	if err != nil {
		return _user, err
	}

	_user.Id = id.InsertedID.(primitive.ObjectID)

	return _user, nil
}

func (mongodb MongoDatabase) RemoveUser(_id primitive.ObjectID) error {
	filter := bson.M{"_id": _id}
	_, err := mongodb.db.Collection("users").DeleteOne(mongodb.ctx, filter)

	return err
}

func (mongodb MongoDatabase) UpdateUser(_user usersvc.User) error {
	filter := bson.M{"_id": _user.Id}
	update := bson.D{{"$set", _user}}
	_, err := mongodb.db.Collection("users").UpdateOne(mongodb.ctx, filter, update)

	return err
}

func (mongodb MongoDatabase) GetUser(_id primitive.ObjectID) (*usersvc.User, error) {
	var user *usersvc.User

	filter := bson.M{"_id": _id}
	err := mongodb.db.Collection("users").FindOne(mongodb.ctx, filter).Decode(&user)

	return user, err
}
