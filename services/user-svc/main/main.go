package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"ptiple/user-svc/api"
	"ptiple/user-svc/mongodatabase"
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
		Addr:     "192.168.99.100:30444",
		Password: "password",
		DB:       0,
	})
	defer redisClient.Close()

	go listenToChickenUpdates(redisClient, mongodb)

	api.Start(&mongodb, redisClient)
}

func listenToChickenUpdates(_redisClient *redis.Client, _db mongodatabase.MongoDatabase) {
	sub := _redisClient.Subscribe("chicken-updates")
	defer sub.Close()

	subChannel := sub.Channel()
	for message := range subChannel {
		log.Println("[ChickenUpdate]", message.Payload)

		messageBody := struct {
			UserId    string `json:"userId"`
			ChickenId string `json:"chickenId"`
			Event     string `json:"event"`
			Data      string `json:"data"`
		}{}
		if err := json.Unmarshal([]byte(message.Payload), &messageBody); err != nil {
			continue
		}

		userId, err := primitive.ObjectIDFromHex(messageBody.UserId)
		if err != nil {
			continue
		}

		user, err := _db.GetUser(userId)
		if err != nil {
			continue
		}

		user.DB = _db

		switch messageBody.Event {
		case "laidGoldEgg":
			if err := user.AddGoldEggs(1); err != nil {
				log.Println("failed adding gold eggs to user", err)
			}
		}
	}
}
