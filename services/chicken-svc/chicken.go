package chickensvc

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseInterface interface {
	InsertChicken(*Chicken) (*Chicken, error)
	UpdateChicken(Chicken) error
}

// TODO find a way to not use ObjectID so this isn't tied to mongoDB
type Chicken struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BelongsToBarn primitive.ObjectID `json:"belongsToBarn" bson:"belongsToBarn"`
	BelongsToUser primitive.ObjectID `json:"belongsToUser" bson:"belongsToUser"`
	PurchaseDay   uint               `json:"purchaseDay" bson:"purchaseDay"`
	EggsLaid      uint               `json:"eggsLaid" bson:"eggsLaid"`
	GoldEggsLaid  uint               `json:"goldEggsLaid" bson:"goldEggsLaid"`
	GoldEggChance uint               `json:"goldEggChance" bson:"goldEggChance"`
	RestingUntil  uint               `json:"restingUntil" bson:"restingUntil"`

	DB  DatabaseInterface `json:"-" bson:"-"`
	Rng func(int) int     `json:"-" bson:"-"`
}

func New(
	_barn primitive.ObjectID,
	_user primitive.ObjectID,
	_purchaseDay uint,
	_db DatabaseInterface,
	_rng func(int) int,
) (*Chicken, error) {
	chicken := &Chicken{
		BelongsToBarn: _barn,
		BelongsToUser: _user,
		PurchaseDay:   _purchaseDay,
		EggsLaid:      0,
		GoldEggsLaid:  0,
		GoldEggChance: uint(_rng(100)),
		DB:            _db,
		Rng:           _rng,
	}

	return _db.InsertChicken(chicken)
}

func (chicken *Chicken) Feed(_currentDay uint) (bool, error) {
	var laidGoldEgg = false
	chicken.EggsLaid++

	rnd := uint(chicken.Rng(100))
	chicken.RestingUntil = _currentDay + rnd/3

	if rnd < chicken.GoldEggChance {
		chicken.GoldEggsLaid++
		laidGoldEgg = true
	}

	return laidGoldEgg, chicken.DB.UpdateChicken(*chicken)
}
