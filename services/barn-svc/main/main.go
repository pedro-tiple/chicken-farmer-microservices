package main

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"ptiple/barn-svc/api"
	"ptiple/barn-svc/mongodatabase"
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

	api.Start(&mongodb, redisClient)
}
