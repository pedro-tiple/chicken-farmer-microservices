package farmersvc

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseInterface interface {
	InsertFarmer(*Farmer) (*Farmer, error)
	UpdateFarmer(Farmer) error
}

// TODO find a way to not use ObjectID so this isn't tied to mongoDB
type Farmer struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	GoldEggs uint               `json:"goldEggs" bson:"goldEggs"`

	DB DatabaseInterface `json:"-" bson:"-"`
}

func New(
	_id primitive.ObjectID,
	_db DatabaseInterface,
) (*Farmer, error) {
	farmer := &Farmer{
		Id:       _id,
		GoldEggs: 0,
		DB:       _db,
	}

	return _db.InsertFarmer(farmer)
}

func (farmer *Farmer) AddGoldEggs(_amount uint) error {
	farmer.GoldEggs += _amount
	return farmer.DB.UpdateFarmer(*farmer)
}

func (farmer *Farmer) RemoveGoldEggs(_amount uint) error {
	result := farmer.GoldEggs - _amount
	if result < 0 {
		return errors.New("cant remove that many gold eggs")
	}

	farmer.GoldEggs = result
	return farmer.DB.UpdateFarmer(*farmer)
}
