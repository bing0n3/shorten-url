package redis

import (
	"github.com/bing0n3/shorten-url/utils"
	"github.com/go-redis/redis/v7"
	"strconv"
)

type SignGenerator struct {
	*redis.Client
}

func InitSignGenerator() *SignGenerator {
	client := redisClient.GetRedisClient().Client
	_, err := client.Get("SID").Result()
	if err == redis.Nil {
		utils.Info.Println("SID didn't exist in Redis, and will set it")
		setErr := client.Set("SID", "10000", 0).Err()
		if setErr != nil {
			utils.Error.Panic("Setting SID in redis Failed")
		}
	} else if err != nil {
		utils.Error.Panic("Setting SID in redis Failed")
	}

	return &SignGenerator{client}
}

func (s *SignGenerator) GetSID() (int64, error) {
	client := s.Client
	client.Incr("SID")
	sid, err := client.Get("SID").Result()
	if err != nil {
		utils.Error.Println("Failed to fetch SID")
	}

	n, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		utils.Error.Printf("Failed to Convert String(%s) to ID", sid)
	}

	return n, err
}
