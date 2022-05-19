package dao

import "github.com/go-redis/redis"

type Redis struct {
	redis *redis.Client
}

func NewRedis(redis *redis.Client) *Redis {
	return &Redis{redis: redis}
}
