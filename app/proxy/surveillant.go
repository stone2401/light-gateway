package proxy

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/stone2401/light-gateway-kernel/pcore"
	"github.com/stone2401/light-gateway-kernel/pkg/zlog"
	"github.com/stone2401/light-gateway/app/model/dao"
	redisGo "github.com/stone2401/light-gateway/app/tools/redis"
	"go.uber.org/zap/zapcore"
)

var surveillant *Surveillant
var onecSurveillantOnce sync.Once

type Surveillant struct {
	rw         sync.RWMutex
	serviceMap map[string]*pcore.Counter
}

func GetSurveillant() *Surveillant {
	onecSurveillantOnce.Do(func() {
		surveillant = NewSurveillant()
	})
	return surveillant
}

func NewSurveillant() *Surveillant {
	return &Surveillant{
		rw: sync.RWMutex{},
		serviceMap: map[string]*pcore.Counter{
			"all": {},
		},
	}
}

func (s *Surveillant) Watch() {
	// 每个整点将上一个整点的数据同步到数据库
	// 计算下一个整点
	for {
		previousTwoWholeHours := getPreviousWholeHours(-1 * time.Hour)
		for Key := range s.serviceMap {
			key := Key + "#" + previousTwoWholeHours.Format("2006010215")
			redisGoCmd := redisGo.GetRedisConn().Get(context.Background(), key)
			if redisGoCmd.Err() != nil {
				continue
			}
			// 将redisGo中的数据同步到数据库
			val, err := redisGoCmd.Int64()
			if err != nil {
				zlog.Zlog().Error("redisGoCmd.Int()", zapcore.Field{Key: "err", Type: zapcore.StringType, String: err.Error()})
				continue
			}
			zlog.Zlog().Info("val", zapcore.Field{Key: "val", Type: zapcore.StringType, String: strconv.Itoa(int(val))})
			qps := &dao.Qps{}
			qps.InsertOrUpdate(key, val)
		}
		time.Sleep(time.Until(time.Now().Truncate(time.Hour).Add(time.Hour).Add(1 * time.Minute)))
	}
}

func getPreviousWholeHours(t time.Duration) time.Time {
	now := time.Now()
	lastWholeHour := now.Truncate(time.Hour)
	previousTwoWholeHours := lastWholeHour.Add(t) // Subtract 2 hours
	return previousTwoWholeHours
}

func (s *Surveillant) Register(serviceName string, counter *pcore.Counter) {
	s.rw.Lock()
	s.serviceMap[serviceName] = counter
	s.rw.Unlock()
	// 将数据同步到redisGo
	go func() {
		for qps := range counter.Gain() {
			previousHours := getPreviousWholeHours(0 * time.Hour)
			key := serviceName + "#" + previousHours.Format("2006010215")
			// 过期时间 3 个小时
			exists, err := redisGo.GetRedisConn().Exists(context.Background(), key).Result()
			if err != nil {
				continue
			}
			if exists == 1 {
				redisGo.GetRedisConn().IncrBy(context.Background(), key, qps.Count)
			} else {
				redisGo.GetRedisConn().Set(context.Background(), key, qps.Count, time.Hour*3)
			}
			// all
			keyAll := "all" + "#" + previousHours.Format("2006010215")
			exists2, err := redisGo.GetRedisConn().Exists(context.Background(), keyAll).Result()
			if err != nil {
				continue
			}
			if exists2 == 1 {
				redisGo.GetRedisConn().IncrBy(context.Background(), keyAll, qps.Count)
			} else {
				redisGo.GetRedisConn().Set(context.Background(), keyAll, qps.Count, time.Hour*3)
			}
		}
	}()
}

// 删除
func (s *Surveillant) Remove(name string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	if s.serviceMap[name] != nil {
		s.serviceMap[name].Close()
		delete(s.serviceMap, name)
	}
}

func (s *Surveillant) Rename(oldName, name string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	if s.serviceMap[oldName] != nil {
		s.serviceMap[name] = s.serviceMap[oldName]
		s.serviceMap[name].SetName(name)
		delete(s.serviceMap, oldName)
	}
}
