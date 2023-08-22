package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/proxy/public"
)

var suffix = "/api/v1/"

func main() {
	e := gin.Default()
	e.GET("/api/v1/info", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})
	s := public.GetRedis().SuffixChack(suffix)
	public.GetRedis().SetKey(s[:len(s)-1]+"127.0.0.1:2401", "0", 5*time.Second)
	ctx, cl := context.WithCancel(context.Background())
	defer cl()
	go public.GetRedis().WatchDog(ctx, s[:len(s)-1]+"127.0.0.1:2401", 5*time.Second, 4*time.Second)
	e.Run(":2401")
}
