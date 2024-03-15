package dao

import (
	"errors"

	"github.com/stone2401/light-gateway/app/public"
)

// 网关权限控制表，所有操作通过 ServiceId 字段进行，必须初始化
type ServiceAccessControl struct {
	Id                uint64 `json:"id" xorm:"bigint pk autoincr 'id' comment('自增主键')"`
	ServiceId         uint64 `json:"service_id" xorm:"bigint 'service_id' unique notnull default(0) comment('服务id')"`
	OpenAuth          bool   `json:"open_auth" xorm:"tinyint 'open_auth' notnull default(0) comment('是否开启权限 1=开启')"`
	BlackList         string `json:"black_list" xorm:"varchar(1000) 'black_list' notnull default('') comment('黑名单')"`
	WhiteList         string `json:"white_list" xorm:"varchar(1000) 'white_list' notnull default('') comment('白名单')"`
	WhiteHostName     string `json:"white_host_name" xorm:"varchar(1000) 'white_host_name' notnull default('') comment('白名单主机')"`
	ClientipFlowLimit int    `json:"clientip_flow_limit" xorm:"int 'clientip_flow_limit' notnull default(0) comment('客户端ip限流')"`
	ServiceFlowLimit  int    `json:"service_flow_limit" xorm:"int 'service_flow_limit' notnull default(0) comment('服务的限流')"`
}

func (s *ServiceAccessControl) GetId() uint64 {
	return s.Id
}

func (m *ServiceAccessControl) Find() (err error) {
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
func (m *ServiceAccessControl) Exist(tag string) error {
	ok, err := GetDBDriver().Exist(m)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}
