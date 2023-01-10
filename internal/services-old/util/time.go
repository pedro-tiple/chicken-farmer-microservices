package util

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
)

// Every call to this function will open a new subscription and channel to redis.
// TODO check if it would be more efficient to share the subscription between listeners and only create new chan
func ListenToTimeUpdates(
	_redisClient *redis.Client, callbackFunc func(uint) error,
) {
	sub := _redisClient.Subscribe("universe-updates")
	defer sub.Close()

	subChannel := sub.Channel()
	for message := range subChannel {
		messageBody := struct {
			Time uint `json:"universe,string"`
		}{}
		if err := json.Unmarshal(
			[]byte(message.Payload), &messageBody,
		); err != nil {
			continue
		}

		if err := callbackFunc(messageBody.Time); err != nil {
			log.Println("failed callback function", err)
			break
		}
	}
}
