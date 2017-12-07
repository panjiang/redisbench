package wares

import (
	"strings"
	"github.com/go-redis/redis"
	"benchmark/config"
)


// A single instance redis client instance
func newRedisClient(addr string, pwd string, db int) (redis.UniversalClient, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{addr},
		Password: pwd,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// A redis cluster client instance
func newRedisClusterClient(nodes_addr string) (redis.UniversalClient, error) {
	addrsArray := strings.Split(nodes_addr, ",")
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: addrsArray,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}


// Create a new universal redis client, no matter single instance or redis cluster
func NewUniversalRedisClient() (redisClient redis.UniversalClient, err error) {
	if config.ClusterMode {
		redisClient, err = newRedisClusterClient(config.RedisAddr)
	} else {
		redisClient, err = newRedisClient(config.RedisAddr, "", 0)
	}
	return
}
