package wares

import (
	"context"
	"runtime"
	"strings"
	"time"

	"github.com/panjiang/redisbench/config"

	"github.com/go-redis/redis/v8"
)

// NewUniversalRedisClient Creates a new universal redis client, no matter single instance or redis cluster
func NewUniversalRedisClient() (redis.UniversalClient, error) {
	addrsArray := strings.Split(config.RedisAddr, ",")
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:       addrsArray,
		Password:    config.RedisPassword,
		DB:          config.RedisDB,
		DialTimeout: time.Second * 3,
		PoolSize:    100 * runtime.NumCPU(),
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
