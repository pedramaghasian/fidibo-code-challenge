package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	options := &redis.Options{
		Addr:     REDIS_ADDRESS,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	}

	client := redis.NewClient(options)

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
	} else {
		fmt.Println("Connected to Redis:", pong)
	}

	return client
}

// docker run -d --name my-redis-container -p 6379:6379 redis:latest
