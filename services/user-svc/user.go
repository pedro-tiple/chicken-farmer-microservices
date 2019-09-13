package usersvc

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseInterface interface {
	InsertUser(*User) (*User, error)
	UpdateUser(User) error
}

// TODO find a way to not use ObjectID so this isn't tied to mongoDB
type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	GoldEggs uint               `json:"goldEggs" bson:"goldEggs"`

	DB DatabaseInterface `json:"-" bson:"-"`
}

func New(
	_id primitive.ObjectID,
	_db DatabaseInterface,
) (*User, error) {
	user := &User{
		Id:       _id,
		GoldEggs: 0,
		DB:       _db,
	}

	return _db.InsertUser(user)
}

func (user *User) AddGoldEggs(_amount uint) error {
	user.GoldEggs += _amount
	return user.DB.UpdateUser(*user)
}

func (user *User) RemoveGoldEggs(_amount uint) error {
	result := user.GoldEggs - _amount
	if result < 0 {
		return errors.New("cant remove that many gold eggs")
	}

	user.GoldEggs = result
	return user.DB.UpdateUser(*user)
}
