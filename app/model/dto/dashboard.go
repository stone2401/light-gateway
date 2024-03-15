package dto

type PanelGroupDataResponse struct {
	ServiceNum      int64 `json:"service_num" form:"service_num"`
	AppNum          int64 `json:"app_num" form:"app_num"`
	CurrentQps      int64 `json:"current_qps" form:"current_qps"`
	TodayRequestNum int64 `json:"today_request_num" form:"today_request_num"`
}

type ServiceStatAllResponse struct {
	Legend []string                     `json:"legend"`
	Data   []ServiceStatAllItemResponse `json:"data"`
}

type ServiceStatAllItemResponse struct {
	Name  string `json:"name"`
	Value uint64 `json:"value"`
}
