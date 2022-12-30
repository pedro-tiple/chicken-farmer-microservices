package main

import (
	"context"
	"encoding/json"
	"log"
	"ptiple/farmersvc/api"
	"ptiple/farmersvc/mongodatabase"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mongodb, err := mongodatabase.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-svc:6379",
		Password: "password",
		DB:       0,
	})
	defer redisClient.Close()

	go listenToChickenUpdates(redisClient, mongodb)

	api.Start(&mongodb, redisClient)
}

func listenToChickenUpdates(_redisClient *redis.Client, _db mongodatabase.MongoDatabase) {
	sub := _redisClient.Subscribe("farm-old-updates")
	defer sub.Close()

	subChannel := sub.Channel()
	for message := range subChannel {
		log.Println("[ChickenUpdate]", message.Payload)

		messageBody := struct {
			FarmerId  string `json:"farmerId"`
			ChickenId string `json:"chickenId"`
			Event     string `json:"event"`
			Data      string `json:"data"`
		}{}
		if err := json.Unmarshal([]byte(message.Payload), &messageBody); err != nil {
			continue
		}

		farmerId, err := primitive.ObjectIDFromHex(messageBody.FarmerId)
		if err != nil {
			continue
		}

		farmer, err := _db.GetFarmer(farmerId)
		if err != nil {
			continue
		}

		farmer.DB = _db

		switch messageBody.Event {
		case "laidGoldEgg":
			if err := farmer.AddGoldEggs(1); err != nil {
				log.Println("failed adding gold eggs to farmer", err)
			}
		}
	}
}
