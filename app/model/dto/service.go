package dto

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/stone2401/light-gateway/app/public"
)

type ServiceListRequest struct {
	Info     string `json:"info" form:"info" lable:"关键词"`
	Page     int    `json:"page" form:"page" query:"page" label:"页数"  example:"1"`
	PageSize int    `json:"pageSize" form:"pageSize" query:"pageSize" label:"条数" example:"20"`
}

type ServiceListResponse struct {
	Total uint64            `json:"total" lable:"总数"`
	Items []ServiceListItem `json:"items"`
	Meta  BaseMeta          `json:"meta"`
}

type ServiceListItem struct {
	ID          uint64 `json:"id"`
	ServiceName string `json:"serviceName"`
	ServiceDesc string `json:"serviceDesc"`
	LoadType    int    `json:"loadType"`
	ServiceAddr string `json:"serviceAddr"`
	QPS         uint64 `json:"qps"`
	QPD         uint64 `json:"qpd"`
	TotalNode   int    `json:"totalNode"`
}
type ServiceInfo struct {
	ID       uint64 `json:"id" form:"id" rule:"notnull"`
	LoadType int    `json:"load_type" form:"load_type" rule:"notnull"`
}
type ServiceLoadBalance struct {
	// service_load_balance 表字段
	RoundType              int    `json:"round_type" form:"round_type" example:"0" label:"轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash"` //轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash
	IpList                 string `json:"ip_list" form:"ip_list" rule:"notnull" label:"ip列表" example:"47.113.203.197:2401"`                            // ip列表
	WeightList             string `json:"weight_list" form:"weight_list" label:"权重列表" example:"10"`                                                    //权重列表
	ForbidList             string `json:"forbid_list" form:"forbid_list" example:"" label:"禁用ip列表"`                                                    // 禁用ip列表
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" example:"0" label:"建立连接超时, 单位s"`                    // 建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" example:"0" label:"获取header超时, 单位s"`                  // 获取header超时, 单位s
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" example:"0" label:"链接最大空闲时间, 单位s"`                        // 链接最大空闲时间, 单位s
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" example:"0" label:"最大空闲链接数"`                                      // 最大空闲链接数
}

func (s *ServiceLoadBalance) VerifyIpWeight(err *strings.Builder) ([]string, []string) {
	ipList := strings.Split(s.IpList, "\n")
	if s.IpList != "" {
		regIp := regexp.MustCompile(`^\S+:\d+$`)
		for i, value := range ipList {
			matched := regIp.MatchString(value)
			if !matched {
				err.WriteString("ip列表格式错误，错误在第" + strconv.Itoa(i) + "行" + public.EndMark)
			}
		}
	}
	weightList := strings.Split(s.WeightList, "\n")
	if s.WeightList != "" {
		regWe := regexp.MustCompile(`^\d+$`)
		for i, value := range weightList {
			matched := regWe.MatchString(value)
			if !matched {
				err.WriteString("权重列表格式错误，错误在第" + strconv.Itoa(i) + "行" + public.EndMark)
			}
		}
	}
	return ipList, weightList
}

type ServiceAccessControl struct {
	// service_access_control 表字段
	OpenAuth          bool   `json:"open_auth" form:"open_auth" example:"false" label:"是否开启权限 1=开启"`             // 是否开启权限 1=开启
	BlackList         string `json:"black_list" form:"black_list" example:"" label:"黑名单"`                        // 黑名单
	WhiteList         string `json:"white_list" form:"white_list" example:"" label:"白名单"`                        // 白名单
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" example:"0" label:"客户端ip限流"` // 客户端ip限流
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" example:"0" label:"服务的限流"`     // 服务的限流
}

type ServiceHttpRule struct {
	RuleType       int    `json:"ruleType" form:"ruleType" label:"匹配类型 0=url前缀url_prefix 1=域名domain" example:"0"`                                                 // 匹配类型 0=url前缀url_prefix 1=域名domain
	Rule           string `json:"rule" form:"rule" rule:"notnull" label:"type=domain表示域名，type=url_prefix时表示url前缀" example:"必填 域名或后缀"`                             // type=domain表示域名，type=url_prefix时表示url前缀
	NeedHttps      bool   `json:"need_https" form:"need_https" label:"支持https 1=支持" example:"false"`                                                              // 支持https 1=支持
	NeedStripUrl   bool   `json:"need_strip_url" form:"need_strip_url" lable:"启用strip_uri 1=启用" example:"false"`                                                  // 启用strip_uri 1=启用
	NeedWebsocket  bool   `json:"need_websocket" form:"need_websocket" label:"是否支持websocket 1=支持" example:"false"`                                                // 是否支持websocket 1=支持
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" label:"url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔" example:""`                           // url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" label:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔" example:""` //header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

type ServiceAddHttpRequest struct {
	// service_info 表字段
	ServiceName string `json:"serviceName" form:"serviceName" rule:"notnull" label:"服务名称 6-128 数字字母下划线" example:"必填 服务名称" regexp:"^[a-zA-Z0-9_]{6,128}$"` // 服务名称
	ServiceDesc string `json:"serviceDesc" form:"serviceDesc" rule:"notnull" label:"服务描述" example:"必填 服务描述" regexp:"^.{5,255}$"`                          // 服务描述
	// service_hrrp_rule 表字段
	ServiceHttpRule
	// service_access_control 表字段
	ServiceAccessControl
	// service_load_balance 表字段
	ServiceLoadBalance
}

func (s *ServiceAddHttpRequest) Check(err *strings.Builder) {
	if s.RoundType > 3 || 0 > s.RoundType {
		err.WriteString("轮询方式错误" + public.EndMark)
	}
	if s.RuleType > 1 || s.RuleType < 0 {
		err.WriteString("匹配类型错误" + public.EndMark)
	}
	if s.UrlRewrite != "" {
		for i, value := range strings.Split(s.UrlRewrite, "\n") {
			if len(strings.Split(value, " ")) != 2 {
				err.WriteString("url重写列表错误, 错误在第" + strconv.Itoa(i) + "行" + public.EndMark)
			}
		}
	}
	if s.HeaderTransfor != "" {
		for i, value := range strings.Split(s.HeaderTransfor, "\n") {
			if len(strings.Split(value, " ")) != 3 {
				err.WriteString("header转换列表错误，错误在第" + strconv.Itoa(i) + "行" + public.EndMark)
			}
		}
	}
	ipList, weightList := s.VerifyIpWeight(err)
	if s.RoundType == 2 && len(ipList) != len(weightList) {
		err.WriteString("ip列表与权重列表数量不一致" + public.EndMark)
	}
}

type ServiceUpdateHttpRequest struct {
	ID uint64 `json:"id" form:"id" rule:"notnull" label:"修改唯一识别"`
	ServiceAddHttpRequest
}

type ServiceAddTcpRequest struct {
	// service_info 表字段
	ServiceName string `json:"serviceName" form:"serviceName" rule:"notnull" label:"服务名称 6-128 数字字母下划线" example:"必填 服务名称" regexp:"^[a-zA-Z0-9_]{6,128}$"` // 服务名称
	ServiceDesc string `json:"serviceDesc" form:"serviceDesc" rule:"notnull" label:"服务描述" example:"必填 服务描述" regexp:"^.{5,255}$"`                          // 服务描述
	// service_hrrp_rule 表字段
	Port int `json:"port" form:"port" xorm:"int 'port' notnull comment('端口')" rule:"notnull" label:"端口"`
	// service_access_control 表字段
	ServiceAccessControl
	// service_load_balance 表字段
	ServiceLoadBalance
}

func (s *ServiceAddTcpRequest) Check(err *strings.Builder) {
	if s.RoundType > 3 || 0 > s.RoundType {
		err.WriteString("轮询方式错误" + public.EndMark)
	}
	if s.Port < 8001 || s.Port > 8999 {
		err.WriteString("端口超出范围" + public.EndMark)
	}
	s.VerifyIpWeight(err)

}

type ServiceUpdateTcpRequest struct {
	ID uint64 `json:"id" form:"id" rule:"notnull" label:"修改唯一识别"`
	ServiceAddTcpRequest
}

type ServiceAddGrpcRequest struct {
	// service_info 表字段
	ServiceName string `json:"serviceName" form:"serviceName" rule:"notnull" label:"服务名称 6-128 数字字母下划线" example:"必填 服务名称" regexp:"^[a-zA-Z0-9_]{6,128}$"` // 服务名称
	ServiceDesc string `json:"serviceDesc" form:"serviceDesc" rule:"notnull" label:"服务描述" example:"必填 服务描述" regexp:"^.{5,255}$"`                          // 服务描述
	// service_hrrp_rule 表字段
	Port           int    `json:"port" form:"port" xorm:"int 'port' notnull comment('端口')" rule:"notnull" label:"端口"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" label:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔" example:""` //header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
	// service_access_control 表字段
	ServiceAccessControl
	// service_load_balance 表字段
	ServiceLoadBalance
}

func (s *ServiceAddGrpcRequest) Check(err *strings.Builder) {
	if s.RoundType > 3 || 0 > s.RoundType {
		err.WriteString("轮询方式错误" + public.EndMark)
	}
	if s.HeaderTransfor != "" {
		for i, value := range strings.Split(s.HeaderTransfor, "\n") {
			if len(strings.Split(value, " ")) != 3 {
				err.WriteString("header转换列表错误，错误在第" + strconv.Itoa(i) + "行" + public.EndMark)
			}
		}
	}
	if s.Port < 8001 || s.Port > 8999 {
		err.WriteString("端口超出范围" + public.EndMark)
	}
	s.VerifyIpWeight(err)

}

type ServiceUpdateGrpcRequest struct {
	ID uint64 `json:"id" form:"id" rule:"notnull" label:"修改唯一识别"`
	ServiceAddGrpcRequest
}

type ServiceStatResponse struct {
	Today     []uint64 `json:"today" form:"today" label:"今日流量"`
	Yesterday []uint64 `json:"yesterday" form:"yesterday" label:"昨日流量"`
}

type ServiceDeleteRequest struct {
	Id uint64 `json:"id" form:"id" rule:"notnull" uri:"id" label:"修改唯一识别"`
}
