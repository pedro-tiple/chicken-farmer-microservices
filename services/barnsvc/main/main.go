package main

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"ptiple/barnsvc/api"
	"ptiple/barnsvc/mongodatabase"
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
