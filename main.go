package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/model/dao"
	v1 "github.com/stone2401/light-gateway/app/router/v1"
	v2 "github.com/stone2401/light-gateway/app/router/v2"
	"github.com/stone2401/light-gateway/app/tools/db"
	"github.com/stone2401/light-gateway/app/tools/redis"
	"github.com/stone2401/light-gateway/config"
	_ "github.com/stone2401/light-gateway/docs/v1"
	_ "github.com/stone2401/light-gateway/docs/v2"
)

func main() {
	config.Init()
	// mysql 连接
	err := db.GetDBDriver().Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.Init()
	// redis 连接
	_, err = redis.GetRedisConn().Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// gin日志打印
	// gin.DefaultWriter = io.MultiWriter(config.GenLogFilename("gin"))
	app := gin.Default()
	app.Use(cors.Default())
	RouterV1 := app.Group("/api/v1")
	{
		v1.RegisterRouterV1(RouterV1)
	}
	RouterV2 := app.Group("/api/v2")
	{
		v2.RegisterRouterV2(RouterV2)
	}
	go app.Run(":2401")

	// 监听系统信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	<-ch
	fmt.Println("Shutdown Server ...")
}
