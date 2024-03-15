package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
)

func RegisterAdminLogin(router *gin.RouterGroup) {
	router.POST("/login", loginAdmin)
}

// @Summary admin登录接口
// @Schemes
// @Description 描述
// @Tags adminApi
// @ID /admin_login/login
// @Param body body dto.AdminLoginRequest true "body"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} middleware.ResponseErr{data=dto.AdminLoginResponse} "success"
// @Router /admin_login/login [post]
func loginAdmin(ctx *gin.Context) {
	var admin = &dto.AdminLoginRequest{}
	if err := public.Authenticator(ctx, admin); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	adminInfo := &dao.Admin{}
	err := adminInfo.LoginCheck(admin)
	if err != nil {
		middleware.ResponseError(ctx, 1002, err)
		return
	}
	token, _ := public.GenerateToken(int(adminInfo.Id), admin.Username)
	out := &dto.AdminLoginResponse{Token: token}
	ctx.Header("Authorization", token)
	middleware.ResponseSuccess(ctx, out)
}
