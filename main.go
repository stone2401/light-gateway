package main

import (
	"io"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/stone2401/light-gateway/app/router/v1"
	v2 "github.com/stone2401/light-gateway/app/router/v2"
	"github.com/stone2401/light-gateway/config"
	_ "github.com/stone2401/light-gateway/docs/v1"
	_ "github.com/stone2401/light-gateway/docs/v2"
)

func main() {
	gin.DefaultWriter = io.MultiWriter(config.GenLogFilename("gin"))
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
	app.Run(":2401")
}
