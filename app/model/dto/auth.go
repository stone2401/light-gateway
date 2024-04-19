package dto

type AuthcaptchaReq struct {
	Width  int `json:"width,omitempty" form:"width" label:"验证码宽度"`
	Height int `json:"height,omitempty" form:"height" label:"验证码高度"`
}

type AuthcaptchaResp struct {
	Img string `json:"img,omitempty" form:"img" label:"验证码数据"`
	Id  string `json:"id,omitempty" form:"id" label:"id"`
}
