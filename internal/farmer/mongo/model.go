package mongo

type Farmer struct {
	ID           string `bson:"_id,omitempty"`
	UUID         string `bson:"uuid,omitempty"`
	FarmID       string `bson:"farmID,omitempty"`
	Name         string `bson:"name,omitempty"`
	PasswordHash string `bson:"passwordHash,omitempty"`
	GoldEggCount uint   `bson:"goldEggCount,omitempty"`
}
