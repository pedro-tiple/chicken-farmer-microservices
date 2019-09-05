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
	BelongsToBarn primitive.ObjectID `json:"belongsToBarn"`
	BirthDay      int                `json:"birthDay"`
	EggsLaid      int                `json:"eggsLaid"`
	GoldEggsLaid  int                `json:"goldEggsLaid"`
	GoldEggChance int                `json:"goldEggChance"`

	DB  DatabaseInterface `json:"-" bson:"-"`
	Rng func(int) int     `json:"-" bson:"-"`
}

func New(
	birthDay int,
	barn primitive.ObjectID,
	_db DatabaseInterface,
	_rng func(int) int,
) (*Chicken, error) {
	chicken := &Chicken{
		BelongsToBarn: barn,
		BirthDay:      birthDay,
		EggsLaid:      0,
		GoldEggsLaid:  0,
		GoldEggChance: _rng(100),
		DB:            _db,
		Rng:           _rng,
	}

	return _db.InsertChicken(chicken)
}

func (chicken *Chicken) Feed() error {
	// TODO consume feed from barn

	chicken.EggsLaid++

	if chicken.Rng(100) < chicken.GoldEggChance {
		chicken.GoldEggsLaid++
		// TODO send message to queue
	}

	return chicken.DB.UpdateChicken(*chicken)
}
