package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func CreateRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	log.Print("Connecting to redis at " + host + ":" + port)

	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		DB:   0,
	})

	if client == nil {
		panic("failed to connect to redis")
	} else {
		log.Print("Connected to redis")
	}

	ctx := context.Background()

	info, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		log.Print("Ping response: " + info)
	}

	return client
}
