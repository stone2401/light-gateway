package handler

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
)

func RegisterAdmin(router *gin.RouterGroup) {
	router.GET("/admin_info", AdminInfo)
	router.POST("/change_pwd", AdminChange)
}

// @Summary admin基本信息
// @Schemes
// @Description 获取admin信息
// @Security ApiKeyAuth
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @PersistAuthorization true
// @Tags adminApi
// @ID /admin/admin_info
// @Accept application/json
// @Produce application/json
// @Success 200 {object} middleware.ResponseErr{data=dto.AdminInfoResponse} "success"
// @Router /admin/admin_info [get]
func AdminInfo(ctx *gin.Context) {
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
	})
}

// @Summary admin密码修改
// @Schemes
// @Description 修改密码
// @Security ApiKeyAuth
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body dto.AdminChangeRequest true "body"
// @Tags adminApi
// @ID /admin/change_pwd
// @Accept application/json
// @Produce application/json
// @Success 200 {string} middleware.ResponseErr{data=dto.AdminChangeRequest} "success"
// @Router /admin/change_pwd [post]
func AdminChange(ctx *gin.Context) {
	// 获取参数
	params := &dto.AdminChangeRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 获取tokne
	value, exists := ctx.Get("token")
	if !exists {
		middleware.ResponseError(ctx, 2000, errors.New("服务器错误，未截取到token, 请检测token状态"))
		return
	}
	token := value.(*public.Claims)
	// 获取 dao admin
	adminInfo := &dao.Admin{UserName: token.UserName}
	err := adminInfo.Find()
	if err != nil {
		middleware.ResponseError(ctx, 1002, errors.New("用户不存在"))
		return
	}
	oldSaltPassword := public.GenSaltPassword(adminInfo.Salt, params.OldPassword)
	if oldSaltPassword != adminInfo.Password {
		middleware.ResponseError(ctx, 1002, errors.New("原始密码错误"))
		return
	}
	newSaltPassword := public.GenSaltPassword(adminInfo.Salt, params.NewPassword)
	err = adminInfo.Update(&dao.Admin{Password: newSaltPassword})
	if err != nil {
		middleware.ResponseError(ctx, 1003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "密码修改成功")

}
