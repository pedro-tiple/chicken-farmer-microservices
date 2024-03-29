package farm

import "github.com/google/uuid"

type Farm struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
}

type Barn struct {
	ID            uuid.UUID
	OwnerID       uuid.UUID
	FarmID        uuid.UUID
	Feed          uint
	HasAutoFeeder bool
}

type Chicken struct {
	ID             uuid.UUID
	OwnerID        uuid.UUID
	BarnID         uuid.UUID
	DateOfBirth    uint
	RestingUntil   uint
	NormalEggsLaid uint
	GoldEggsLaid   uint
	GoldEggChance  uint
}

type FarmDetailsResult struct {
	Farm
	GoldEggCount uint
	CurrentDay   uint
	Barns        []farmDetailsResultBarn
}

type farmDetailsResultBarn struct {
	Barn
	Chickens []Chicken
}
