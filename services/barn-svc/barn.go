package barnsvc

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseInterface interface {
	InsertBarn(*Barn) (*Barn, error)
	UpdateBarn(Barn) error
}

// TODO find a way to not use ObjectID so this isn't tied to mongoDB
type Barn struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BelongsToUser primitive.ObjectID `json:"belongsToUser" bson:"belongsToUser"`
	Feed          uint               `json:"feed" bson:"feed"`
	AutoFeeder    bool               `json:"autoFeeder" bson:"autoFeeder"`

	DB DatabaseInterface `json:"-" bson:"-"`
}

func New(
	_userId primitive.ObjectID,
	_db DatabaseInterface,
) (*Barn, error) {
	barn := &Barn{
		BelongsToUser: _userId,
		Feed:          100,
		AutoFeeder:    false,
		DB:            _db,
	}

	return _db.InsertBarn(barn)
}

func (barn *Barn) AddFeed(_amount uint) error {
	barn.Feed += _amount
	return barn.DB.UpdateBarn(*barn)
}

func (barn *Barn) RemoveFeed(_amount uint) error {
	if _amount > barn.Feed {
		return errors.New("cant remove that much feed")
	}

	barn.Feed = barn.Feed - _amount
	return barn.DB.UpdateBarn(*barn)
}
