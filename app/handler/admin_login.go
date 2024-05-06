package handler

import (
	"bytes"
	"encoding/base64"
	"errors"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/app/tools/redis"
)

func RegisterAdminLogin(router *gin.RouterGroup) {
}

func RegisterAuth(router *gin.RouterGroup) {
	// 配置自定义存储
	captcha.SetCustomStore(redis.NewStore())
	// 用户登录
	router.POST("/login", loginAdmin)
	// auth 验证码
	router.GET("/captcha", authcaptcha)
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
	// 校验验证码
	ok := captcha.VerifyString(admin.CaptchaaId, admin.VerifyCode)
	if !ok {
		middleware.ResponseError(ctx, 1002, errors.New("验证码错误"))
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

// @Summary captcha 验证码
// @Schemes
// @Description 验证码接口
// @Tags adminApi
// @ID /auth/captcha
// @Param width query string false "宽度"
// @Param height query string false "高度"
// @Success 200 {object} middleware.ResponseErr{data=dto.AuthcaptchaResp} "success"
// @Router /auth/captcha [get]
func authcaptcha(ctx *gin.Context) {
	var req = &dto.AuthcaptchaReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	if req.Width == 0 {
		req.Width = 100
	}
	if req.Height == 0 {
		req.Height = 50
	}
	id := captcha.NewLen(public.CAPTCHALEN)
	// 取出数据
	buf := new(bytes.Buffer)
	err := captcha.WriteImage(buf, id, req.Width, req.Height)
	if err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 转base64
	data := "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
	middleware.ResponseSuccess(ctx, &dto.AuthcaptchaResp{Id: id, Img: data})
}
