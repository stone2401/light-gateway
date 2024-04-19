package redis

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/config"
)

var redisOnce sync.Once
var redisClient *redis.Client

func GetRedisConn() *redis.Client {
	redisOnce.Do(func() {
		fmt.Println("redis 连接中...", config.Config.GetRedisConfig())
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.Config.GetRedisConfig(),
			Password: config.Config.Redis.Password,
			DB:       config.Config.Redis.Database,
		})
	})
	return redisClient
}

type Store struct {
	client *redis.Client
}

func NewStore() *Store {
	return &Store{
		client: GetRedisConn(),
	}
}

func (s *Store) Set(id string, digits []byte) {
	if os.Getenv("GIN_MODE") == "dev" || os.Getenv("GIN_MODE") == "" {
		fmt.Println("redis set", id, "digits:", digits, len(digits))
	}
	s.client.Set(context.Background(), public.CAPTCHAKEY+id, digits, 5*60*time.Second).Err()
}

func (s *Store) Get(id string, clear bool) []byte {
	digits, err := s.client.Get(context.Background(), public.CAPTCHAKEY+id).Bytes()
	if err != nil {
		return nil
	}
	if clear {
		s.client.Del(context.Background(), public.CAPTCHAKEY+id)
	}
	return digits
}
