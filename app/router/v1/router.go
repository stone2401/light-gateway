package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/handler"
	"github.com/stone2401/light-gateway/app/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 轻量级网关
// @version 1.0
// @description 一个轻量级网关代理，支持多种代理协议，接口转发，数据统计，带有管理界面
// @contact.email stone2401@qq.com
// @host localhost:2401
// @schemes http
// @BasePath /api/v1
func RegisterRouterV1(v1 *gin.RouterGroup) {
	v1.GET("/ping", handler.Ping())
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
		c.InstanceName = "v1"
	}))
	authGroup := v1.Group("/auth", middleware.RecoveryMiddleware())
	adminLoginGroup := v1.Group("/admin_login", middleware.RecoveryMiddleware())
	adminGroup := v1.Group("/admin", middleware.RecoveryMiddleware(), middleware.TokenMiddleware())
	serviceGroup := v1.Group("/service", middleware.RecoveryMiddleware(), middleware.TokenMiddleware())
	appGroup := v1.Group("/app", middleware.RecoveryMiddleware(), middleware.TokenMiddleware())
	dashboardGroup := v1.Group("/dashboard", middleware.RecoveryMiddleware(), middleware.TokenMiddleware())
	system := v1.Group("/system", middleware.RecoveryMiddleware(), middleware.TokenMiddleware())
	handler.RegisterAuth(authGroup)
	handler.RegisterAdminLogin(adminLoginGroup)
	handler.RegisterAdmin(adminGroup)
	handler.RegisterService(serviceGroup)
	handler.RegisterApp(appGroup)
	handler.RegisterDashBoard(dashboardGroup)
	handler.RegisterSystem(system)

}
