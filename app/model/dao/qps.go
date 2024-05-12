package dao

import (
	"strings"

	"github.com/stone2401/light-gateway-kernel/pkg/zlog"
	"github.com/stone2401/light-gateway/app/tools/db"
	"go.uber.org/zap/zapcore"
)

type Qps struct {
	Id        int64 `xorm:"bigint pk autoincr 'id'"`
	ServiceId int64 `xorm:"bigint 'service_id'"`
	// 时间
	Time string `xorm:"varchar(255) 'time'"`
	// 不同时间 24小时
	// 0 点
	Hour int64 `xorm:"bigint 'hour'"`
	// 1 点
	Hour1 int64 `xorm:"bigint 'hour1'"`
	// 2 点
	Hour2 int64 `xorm:"bigint 'hour2'"`
	// 3 点
	Hour3 int64 `xorm:"bigint 'hour3'"`
	// 4 点
	Hour4 int64 `xorm:"bigint 'hour4'"`
	// 5 点
	Hour5 int64 `xorm:"bigint 'hour5'"`
	// 6 点
	Hour6 int64 `xorm:"bigint 'hour6'"`
	// 7 点
	Hour7 int64 `xorm:"bigint 'hour7'"`
	// 8 点
	Hour8 int64 `xorm:"bigint 'hour8'"`
	// 9 点
	Hour9 int64 `xorm:"bigint 'hour9'"`
	// 10 点
	Hour10 int64 `xorm:"bigint 'hour10'"`
	// 11 点
	Hour11 int64 `xorm:"bigint 'hour11'"`
	// 12 点
	Hour12 int64 `xorm:"bigint 'hour12'"`
	// 13 点
	Hour13 int64 `xorm:"bigint 'hour13'"`
	// 14 点
	Hour14 int64 `xorm:"bigint 'hour14'"`
	// 15 点
	Hour15 int64 `xorm:"bigint 'hour15'"`
	// 16 点
	Hour16 int64 `xorm:"bigint 'hour16'"`
	// 17 点
	Hour17 int64 `xorm:"bigint 'hour17'"`
	// 18 点
	Hour18 int64 `xorm:"bigint 'hour18'"`
	// 19 点
	Hour19 int64 `xorm:"bigint 'hour19'"`
	// 20 点
	Hour20 int64 `xorm:"bigint 'hour20'"`
	// 21 点
	Hour21 int64 `xorm:"bigint 'hour21'"`
	// 22 点
	Hour22 int64 `xorm:"bigint 'hour22'"`
	// 23 点
	Hour23 int64 `xorm:"bigint 'hour23'"`
}

func (q *Qps) InsertOrUpdate(key string, value int64) {
	// key 是通过 # 分割的
	keys := strings.Split(key, "#")
	serviceName := strings.Join(keys[:len(keys)-2], "#")
	timeString := keys[len(keys)-1]
	// timeString 最后2位置是小时，需要截取出来
	hour := timeString[len(timeString)-2:]
	times := timeString[:len(timeString)-2]
	qps := &Qps{
		ServiceId: 0,
		Time:      times,
	}
	qps.SetHour(hour, value)
	session := db.GetDBDriver().NewSession()
	defer session.Close()
	info := &ServiceInfo{
		ServiceName: serviceName,
	}
	if serviceName != "all" {
		err := info.Find()
		if err != nil {
			return
		}
		qps.ServiceId = int64(info.Id)
	}

	if hour == "00" {
		_, err := session.InsertOne(qps)
		if err != nil {
			zlog.Zlog().Error("insert qps error", zapcore.Field{Key: "err", Type: zapcore.StringType, String: err.Error()})
		}
		return
	} else {
		_, err := session.Update(qps)
		if err != nil {
			zlog.Zlog().Error("update qps error", zapcore.Field{Key: "err", Type: zapcore.StringType, String: err.Error()})
		}
		return
	}
}

func (q *Qps) SetHour(hour string, value int64) {
	switch hour {
	case "00":
		q.Hour = value
	case "01":
		q.Hour1 = value
	case "02":
		q.Hour2 = value
	case "03":
		q.Hour3 = value
	case "04":
		q.Hour4 = value
	case "05":
		q.Hour5 = value
	case "06":
		q.Hour6 = value
	case "07":
		q.Hour7 = value
	case "08":
		q.Hour8 = value
	case "09":
		q.Hour9 = value
	case "10":
		q.Hour10 = value
	case "11":
		q.Hour11 = value
	case "12":
		q.Hour12 = value
	case "13":
		q.Hour13 = value
	case "14":
		q.Hour14 = value
	case "15":
		q.Hour15 = value
	case "16":
		q.Hour16 = value
	case "17":
		q.Hour17 = value
	case "18":
		q.Hour18 = value
	case "19":
		q.Hour19 = value
	case "20":
		q.Hour20 = value
	case "21":
		q.Hour21 = value
	case "22":
		q.Hour22 = value
	case "23":
		q.Hour23 = value
	default:
		zlog.Zlog().Error("set hour error", zapcore.Field{Key: "err", Type: zapcore.StringType, String: "set hour error"})
	}
}
