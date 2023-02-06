package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	redisPassword := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: redisPassword,
		DB:       0,
	})

	topic := os.Getenv("TOPIC")
	err := client.Publish(topic, "Hello, World!").Err()
	if err != nil {
		log.Fatalf("failed to publish message: %v", err)
	}
}
