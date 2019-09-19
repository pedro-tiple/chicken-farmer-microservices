package main

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"ptiple/chicken-svc/api"
	"ptiple/chicken-svc/mongodatabase"
	"ptiple/util"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mongodb, err := mongodatabase.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// if day isn't set, initialize it as 0
	_, err = mongodb.GetDay()
	if err != nil {
		_ = mongodb.UpdateDay(0)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-svc:6379",
		Password: "password",
		DB:       0,
	})

	go util.ListenToTimeUpdates(redisClient, mongodb.UpdateDay)

	api.Start(&mongodb, redisClient)
}
