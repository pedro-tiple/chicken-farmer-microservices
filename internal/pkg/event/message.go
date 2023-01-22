package event

import (
	"context"
	"encoding/json"

	messagePkg "github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

const (
	MetadataFieldFarmerID = "farmerID"
	MetadataFieldType     = "message-type"

	MessageTypeDay              = "day"
	MessageTypeNewBarn          = "new-barn"
	MessageTypeNewChicken       = "new-chicken"
	MessageTypeFeedChange       = "feed-change"
	MessageTypeGoldenEggsChange = "golden-eggs-change"
	MessageTypeChickenFed       = "chicken-fed"
)

// Future thing: send these messages encoded with protobuf?

type DayMessage struct {
	Day uint `json:"day"`
}

type NewBarnMessage struct {
	ID            string `json:"barnID"`
	Feed          uint   `json:"feed"`
	HasAutoFeeder bool   `json:"hasAutoFeeder"`
}

type NewChickenMessage struct {
	BarnID         string `json:"barnID"`
	ChickenID      string `json:"chickenID"`
	DateOfBirth    uint   `json:"dateOfBirth"`
	RestingUntil   uint   `json:"restingUntil"`
	NormalEggsLaid uint   `json:"normalEggsLaid"`
	GoldEggsLaid   uint   `json:"goldEggsLaid"`
}

type FeedChangeMessage struct {
	BarnID string `json:"barnID"`
	Count  int    `json:"count"`
}

type GoldenEggChangeMessage struct {
	Count int `json:"count"`
}

type ChickenFedMessage struct {
	ChickenID      string `json:"chickenID"`
	RestingUntil   uint   `json:"restingUntil"`
	NormalEggsLaid uint   `json:"normalEggsLaid"`
	GoldEggsLaid   uint   `json:"goldEggsLaid"`
}

func PublishMessage(
	ctx context.Context,
	publisher messagePkg.Publisher,
	farmerID uuid.UUID,
	topic, messageType string,
	message any,
) error {
	eggLaidMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := messagePkg.NewMessage(uuid.New().String(), eggLaidMessage)
	msg.SetContext(ctx)
	msg.Metadata[MetadataFieldFarmerID] = farmerID.String()
	msg.Metadata[MetadataFieldType] = messageType

	return publisher.Publish(topic, msg)
}
