package barnsvc_test

import (
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	barnsvc "ptiple/barn-svc"
	"ptiple/barn-svc/mocks"
	"testing"
)

func TestBarn_New(t *testing.T) {
	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongodb := mocks.NewMockIMongoDatabase(ctrl)
	mongodb.
		EXPECT().
		InsertBarn(gomock.Eq(&barnsvc.Barn{
			BelongsToUser: userId,
			Feed:          100,
			AutoFeeder:    false,
			DB:            mongodb,
		})).
		Return(&barnsvc.Barn{
			Id:            primitive.NewObjectID(),
			BelongsToUser: userId,
			Feed:          100,
			AutoFeeder:    false,
			DB:            mongodb,
		}, nil)

	barn, err := barnsvc.New(userId, mongodb)

	if err != nil {
		t.Errorf("Expected no errors got %s", err)
	}

	if barn.Feed != 100 {
		t.Errorf("Expected 100 default feed got %d", barn.Feed)
	}

	if barn.AutoFeeder {
		t.Error("Expected false AutoFeeder by default feed got true")
	}

	if barn.BelongsToUser != userId {
		t.Errorf("Expected owner to be %s got %s", userId, barn.BelongsToUser)
	}
}

func TestBarn_AddFeed(t *testing.T) {
	userId := primitive.NewObjectID()
	barnId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongodb := mocks.NewMockIMongoDatabase(ctrl)
	mongodb.
		EXPECT().
		UpdateBarn(gomock.Eq(barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          1,
			AutoFeeder:    false,
			DB:            mongodb,
		})).
		Return(nil)

	barn := barnsvc.Barn{
		Id:            barnId,
		BelongsToUser: userId,
		Feed:          0,
		AutoFeeder:    false,
		DB:            mongodb,
	}

	err := barn.AddFeed(1)
	if err != nil {
		t.Errorf("Expected no errors got %s", err)
	}

	if barn.Feed != 1 {
		t.Errorf("Expected feed to be 1 got %d", barn.Feed)
	}
}

func TestBarn_RemoveFeed(t *testing.T) {
	userId := primitive.NewObjectID()
	barnId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongodb := mocks.NewMockIMongoDatabase(ctrl)
	mongodb.
		EXPECT().
		UpdateBarn(gomock.Eq(barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    false,
			DB:            mongodb,
		})).
		Return(nil)

	barn := barnsvc.Barn{
		Id:            barnId,
		BelongsToUser: userId,
		Feed:          1,
		AutoFeeder:    false,
		DB:            mongodb,
	}

	err := barn.RemoveFeed(1)
	if err != nil {
		t.Errorf("Expected no errors got %s", err)
	}

	if barn.Feed != 0 {
		t.Errorf("Expected feed to be 0 got %d", barn.Feed)
	}
}

func TestBarn_RemoveFeedMustHaveEnough(t *testing.T) {
	userId := primitive.NewObjectID()
	barnId := primitive.NewObjectID()

	barn := barnsvc.Barn{
		Id:            barnId,
		BelongsToUser: userId,
		Feed:          1,
		AutoFeeder:    false,
		DB:            nil,
	}

	err := barn.RemoveFeed(2)
	if err == nil {
		t.Error("Expected to fail because of not enough feed")
	}
}
