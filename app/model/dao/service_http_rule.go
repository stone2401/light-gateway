package dao

import (
	"github.com/pkg/errors"

	"github.com/stone2401/light-gateway/app/public"
)

// http网关路由匹配表，所有操作通过 ServiceId 字段进行，必须初始化
type ServiceHttpRule struct {
	Id             uint64 `json:"id" xorm:"bigint pk notnull autoincr 'id' comment('自增主键')"`
	ServiceId      uint64 `json:"service_id" xorm:"bigint 'service_id' unique notnull index comment('服务id')"`
	RuleType       int    `json:"rule_type" xorm:"int 'rule_type' notnull default(0) comment('匹配类型 0=url前缀url_prefix 1=域名domain')"`
	Rule           string `json:"rule" xorm:"varchar(255) 'rule' notnull default('') comment('type=domain表示域名，type=url_prefix时表示url前缀')"`
	NeedHttps      bool   `json:"need_https" xorm:"int 'need_https' notnull default(0) comment('支持https 1=支持')"`
	NeedStripUrl   bool   `json:"need_strip_url" xorm:"tinyint 'need_strip_url' notnull default(0) comment('启用strip_uri 1=启用')"`
	NeedWebsocket  bool   `json:"need_websocket" xorm:"tinyint 'need_websocket' notnull default(0) comment('是否支持websocket 1=支持')"`
	UrlRewrite     string `json:"url_rewrite" xorm:"varchar(5000) 'url_rewrite' notnull default('') comment('url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔')"`
	HeaderTransfor string `json:"header_transfor" xorm:"varchar(5000) notnull default('') 'header_transfor' comment('header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔')"`
}

func (m *ServiceHttpRule) Find() (err error) {
	b, err := GetDBDriver().Get(m)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("数据不存在，导致查询错误")
	}
	return nil
}

// 以自身为条件，判断是否存在
func (m *ServiceHttpRule) Exist(tag string) error {
	ok, err := GetDBDriver().Exist(m)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}

func (s *ServiceHttpRule) GetId() uint64 {
	return s.Id
}
