package handler

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.GET("/profile", AccountProfile)
	router.GET("/permissions", AccountPermissions)
	router.GET("/menus", AccountMenu)
	router.GET("/logout", AccountLogout)
}

// @Summary 获取账户资料
// @Schemes
// @Description 获取账户资料
// @Tags accountApi
// @ID /account/info
// @Accept application/json
// @Produce application/json
// @Success 200 {object} middleware.ResponseErr{data=dto.AccountInfoResponse} "success"
// @Router /account/info [get]
func AccountProfile(ctx *gin.Context) {
	// 获取token
	value, exists := ctx.Get("token")
	if !exists {
		middleware.ResponseError(ctx, 2000, errors.New("服务器错误，未截取到token, 请检测token状态"))
		return
	}
	token := value.(*public.Claims)
	middleware.ResponseSuccess(ctx, &dto.AdminInfoResponse{
		ID:        token.ID,
		Name:      token.UserName,
		LoginTime: time.Unix(token.LoginTime.Unix(), 0).Format("2006-01-02 15:04:05"),
		Avatar:    "https://thirdqq.qlogo.cn/g?b=qq&s=100&nk=1743369777",
		Remark:    "管理员",
		Roles:     []string{"0"},
		Nickname:  "stone2401",
		Email:     "2919390584@qq.com",
		Phone:     "15931112401",
		Ip:        ctx.ClientIP(),
	})
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
func AccountMenu(ctx *gin.Context) {
	// 读取文件, 先获取根路径
	root, err := os.Getwd()
	if err != nil {
		middleware.ResponseError(ctx, 2000, errors.New("用户菜单读取失败"))
		return
	}
	f, err := os.ReadFile(path.Join(root, "conf", "route_group.json"))
	if err != nil {
		middleware.ResponseError(ctx, 2000, errors.New("用户菜单读取失败"))
		return
	}
	var menu []gin.H
	json.Unmarshal(f, &menu)
	middleware.ResponseSuccess(ctx, (menu))
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
func AccountPermissions(ctx *gin.Context) {
	middleware.ResponseSuccess(ctx, public.SystemPermissions)
}

// @Summary 登出
// @Description 登出
// @Tags accountApi
// @ID /account/logout
// @Accept application/json
// @Produce application/json
// @Success 200 {object} middleware.ResponseErr{} "success"
// @Router /account/logout [get]
func AccountLogout(ctx *gin.Context) {
	middleware.ResponseSuccess(ctx, nil)
}
