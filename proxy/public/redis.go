package public

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var rc *redisClient
var once sync.Once

type redisClient struct {
	rdb *redis.Client
}

const (
	PREFIX = "__keyspace@0__:"
)

func GetRedis() *redisClient {
	once.Do(func() {
		rc = &redisClient{
			rdb: redis.NewClient(&redis.Options{Password: "!Shizhenfei123", DB: 0}),
		}
	})
	return rc
}

func (r *redisClient) SuffixChack(suffix string) string {
	if suffix[len(suffix)-2:] != ":*" {
		suffix += ":*"
	}
	return suffix
}

func (r *redisClient) PSubscribe(channels string) *redis.PubSub {
	ps := r.rdb.PSubscribe(context.Background(), channels)
	return ps
}

func (r *redisClient) GetKeys(pattern string) []string {
	ssc := r.rdb.Keys(context.Background(), pattern)
	if ssc.Err() != nil {
		return []string{}
	}
	return ssc.Val()
}

func (r *redisClient) GetValue(key string) string {
	sc := r.rdb.Get(context.Background(), key)
	if sc.Err() != nil {
		return ""
	}
	return sc.Val()
}

func (r *redisClient) SetKey(key, value string, dura time.Duration) string {
	sc := r.rdb.SetEx(context.Background(), key, value, dura)
	if sc.Err() != nil {
		os.Exit(-1)
	}
	return sc.Val()
}

// 看门狗， ctx 结束则结束，dura 续时, tickerTime 循环时间
// 建议 tickerTime < dura, 切足够小
func (r *redisClient) WatchDog(ctx context.Context, key string, dura, tickerTime time.Duration) {
	if key == "" {
		return
	}
	if dura == 0 {
		dura = 10 * time.Second
	}
	ticker := time.NewTicker(tickerTime)
	for {
		select {
		case <-ticker.C:
			log.Println(key)
			r.rdb.Expire(ctx, key, dura)
		case <-ctx.Done():
			return
		}
	}
}
