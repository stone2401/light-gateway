package dto

type AppListResponse struct {
	List  []*AppListItem `json:"list" form:"list"`
	Total int64          `json:"total" form:"total"`
}

type AppListItem struct {
	Id       uint64 `json:"id"`
	AppId    string `json:"app_id" rule:"notnull" label:"租户id" regexp:"^[a-zA-Z0-9_]{6,128}$"`
	Name     string `json:"name" rule:"notnull" label:"租户名称" regexp:"^.{0,255}$"`
	Secret   string `json:"secret" label:"密钥"`
	WhiteIps string `json:"white_ips" label:"ip白名单，支持前缀匹配"`
	Qpd      uint64 `json:"qpd" label:"日请求量限制"`
	Qps      uint64 `json:"qps" label:"每秒请求量限制"`
	RealQpd  int64  `json:"real_qpd"`
	RealQps  int64  `json:"real_qps"`
}

type AppDetailRequest struct {
	ID uint64 `json:"id" form:"id" rule:"notnull"`
}
