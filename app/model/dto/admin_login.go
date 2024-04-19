package dto

type AdminLoginRequest struct {
	Username   string `json:"username" form:"username" rule:"notnull" label:"用户名" regexp:"account" example:"admin"`
	Password   string `json:"password" form:"password" rule:"notnull" label:"密码" example:"admin"`
	VerifyCode string `json:"verifyCode" form:"verifyCode" rule:"notnull" label:"验证码" example:"1234"`
	CaptchaaId string `json:"captchaId" form:"captchaId" rule:"notnull" label:"验证码ID" example:"captchaa_id"`
}

type AdminLoginResponse struct {
	Token string `json:"token" form:"token"`
}
