package dto

type AdminLoginRequest struct {
	Username string `json:"username" form:"username" rule:"notnull" label:"用户名" regexp:"account" example:"admin"`
	Password string `json:"password" form:"password" rule:"notnull" label:"密码" example:"admin"`
}

type AdminLoginResponse struct {
	Token string `json:"token" form:"token"`
}
