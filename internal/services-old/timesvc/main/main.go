package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

const channelName = "universe-updates"

func main() {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     "redis-svc:6379",
			Password: "password",
			DB:       0,
		},
	)
	defer redisClient.Close()

	tickTime(redisClient)
}

func tickTime(_redisClient *redis.Client) {
	// TODO persist universe between restarts?
	var currentTime = 0
	for {
		time.Sleep(1 * time.Second)

		currentTime++

		err := _redisClient.Publish(
			channelName,
			fmt.Sprintf(`{"universe": "%d"}`, currentTime),
		).Err()
		if err != nil {
			log.Println("failed publishing to redis", err)
		}
	}
}
