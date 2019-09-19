package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

const channelName = "time-updates"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-svc:6379",
		Password: "password",
		DB:       0,
	})
	defer redisClient.Close()

	tickTime(redisClient)
}

func tickTime(_redisClient *redis.Client) {
	// TODO persist time between restarts?
	var currentTime = 0
	for {
		time.Sleep(1 * time.Second)

		currentTime++

		err := _redisClient.Publish(
			channelName,
			fmt.Sprintf(`{"time": "%d"}`, currentTime),
		).Err()
		if err != nil {
			log.Println("failed publishing to redis", err)
		}
	}
}
