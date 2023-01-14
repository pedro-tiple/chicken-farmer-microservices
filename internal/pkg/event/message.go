package event

import (
	"encoding/json"

	messagePkg "github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

const (
	MetadataFieldFarmerID = "farmerID"
	MetadataFieldType     = "message-type"

	MessageTypeDay     = "day"
	MessageTypeEggLaid = "egg-laid"
)

type DayMessage struct {
	Day uint `json:"day"`
}

type EggLaidMessage struct {
	Day          uint   `json:"day"`
	ChickenID    string `json:"chickenID"`
	EggType      int    `json:"eggType"`
	RestingUntil uint   `json:"restingUntil"`
}

func PublishMessage(
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
	msg.Metadata[MetadataFieldFarmerID] = farmerID.String()
	msg.Metadata[MetadataFieldType] = messageType

	return publisher.Publish(topic, msg)
}
