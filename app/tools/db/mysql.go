package db

import (
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stone2401/light-gateway/config"
	"xorm.io/xorm"

	xormLog "xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine
var dbonce sync.Once

// xorm Engine 单例
func GetDBDriver() *xorm.Engine {
	// 只执行一次，切确保线程安全
	dbonce.Do(func() {
		var err error
		engine, err = xorm.NewEngine(config.Config.DriverName, config.Config.GetDatabaseConfig())
		if err != nil {
			log.Println("[!] 数据库连接失败：", err)
		}
		if err2 := engine.Ping(); err2 != nil {
			log.Println("[!] 数据库 ping 失败", err2)
		}
		// 最大连接数设置
		engine.SetMaxOpenConns(30)
		engine.SetMaxIdleConns(10)
		engine.SetConnMaxLifetime(30 * time.Minute)
		// 日志打印设置
		// engine.SetLogger(xormLog.NewSimpleLogger(config.GenLogFilename("xorm")))
		engine.Logger().SetLevel(xormLog.LOG_DEBUG)
		// 设置前缀
		tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, "gateway_")
		engine.SetTableMapper(tbMapper)
	})
	return engine
}
