package wares

import (
	"context"
	"strings"

	"github.com/panjiang/redisbench/config"

	"github.com/go-redis/redis/v8"
)

// A single instance redis client instance
func newRedisClient(addr string, pwd string, db int) (redis.UniversalClient, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{addr},
		Password: pwd,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// A redis cluster client instance
func newRedisClusterClient(nodesAddr string, pwd string) (redis.UniversalClient, error) {
	addrsArray := strings.Split(nodesAddr, ",")
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    addrsArray,
		Password: pwd,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewUniversalRedisClient Creates a new universal redis client, no matter single instance or redis cluster
func NewUniversalRedisClient() (redisClient redis.UniversalClient, err error) {
	if config.ClusterMode {
		redisClient, err = newRedisClusterClient(config.RedisAddr, config.RedisPassword)
	} else {
		redisClient, err = newRedisClient(config.RedisAddr, config.RedisPassword, config.RedisDB)
	}
	return
}
