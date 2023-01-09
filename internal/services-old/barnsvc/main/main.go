package main

import (
	"context"
	"log"
	"ptiple/barnsvc/api"
	"ptiple/barnsvc/mongodatabase"

	"github.com/go-redis/redis"
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

	api.Start(mongodb, redisClient)
}
