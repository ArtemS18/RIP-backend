package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"failiverCheck/internal/app/config"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	cfg    *config.RedisConfig
	client *redis.Client
}

func New(cfg *config.RedisConfig) (Redis, error) {
	client := Redis{}

	client.cfg = cfg

	redisClient := redis.NewClient(&redis.Options{
		Password:    cfg.Password,
		Username:    cfg.User,
		Addr:        cfg.Host + ":" + strconv.Itoa(cfg.Port),
		DB:          0,
		DialTimeout: time.Duration(cfg.DialTimeoutSec) * time.Second,
		ReadTimeout: time.Duration(cfg.ReadTimeoutSec) * time.Second,
	})

	client.client = redisClient
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return Redis{}, fmt.Errorf("cant ping redis: %w", err)
	}

	return client, nil
}

func (c *Redis) Close() error {
	return c.client.Close()
}
