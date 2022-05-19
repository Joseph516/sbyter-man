package cache

import "github.com/go-redis/redis"

// redis嵌入

type Redis struct {
	redis *redis.Client
}

func NewRedis(redis *redis.Client) *Redis {
	return &Redis{redis: redis}
}
