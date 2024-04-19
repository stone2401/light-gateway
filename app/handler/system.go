package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/public"
)

func RegisterSystem(router *gin.RouterGroup) {
	router.GET("/menus", systemctlMenu)
	router.GET("/permissions", systemctlPermissions)
}

// @Summary menu 列表
// @Description menu 列表
// @Tags system
// @ID /system/menu
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{} "success"
// @Router /system/menus [get]
func systemctlMenu(ctx *gin.Context) {
	middleware.ResponseSuccess(ctx, public.SystemMenu)
}

// @Summary permissions 列表
// @Description permissions 列表
// @Tags system
// @ID /system/permissions
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{} "success"
// @Router /system/permissions [get]
func systemctlPermissions(ctx *gin.Context) {
	middleware.ResponseSuccess(ctx, public.SystemPermissions)
}
