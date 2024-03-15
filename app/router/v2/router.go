package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 轻量级网关
// @version 1.0
// @description 一个轻量级网关代理，支持多种代理协议，接口转发，数据统计，带有管理界面
// @tag.name example
// @contact.email stone2401@qq.com
// @host 127.0.0.1:2401
// @schemes http
// @BasePath /api/v2
func RegisterRouterV2(v2 *gin.RouterGroup) {
	v2.GET("/ping", handler.Ping())
	v2.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
		c.InstanceName = "v2"
	}))
}
