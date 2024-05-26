package Handlers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	RC  *redis.Client
	Ctx context.Context
}

func NewRedis(redisClient *redis.Client) RedisHandler {
	return RedisHandler{redisClient, context.Background()}
}
