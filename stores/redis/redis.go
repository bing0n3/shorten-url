package redis

import (
	"github.com/bing0n3/shorten-url/utils"
	"github.com/go-redis/redis/v7"
	"sync"
)

type RedisClient struct {
	Client *redis.Client
}

var instance *RedisClient
var once sync.Once

// GetRedisClient get a RedisClient instance.
func GetRedisClient() *RedisClient {
	once.Do(func() {
		instance = &RedisClient{redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			PoolSize: 5,
		})}
	})

	// check redis server status
	_, err := instance.Client.Ping().Result()
	if err != nil {
		utils.Error.Printf("Server failed to connect Redis server")
	}

	return instance
}
