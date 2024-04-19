package dao

import (
	"errors"
	"strings"

	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/app/tools/db"
)

// 网关负载表，ServiceId 字段进行，必须初始化
type ServiceLoadBalance struct {
	Id                     uint64 `json:"id" xorm:"bigint pk notnull autoincr 'id' comment('自增主键')"`
	ServiceId              uint64 `json:"service_id" xorm:"bigint 'service_id' unique notnull index comment('服务id')"`
	CheckMethod            int    `json:"check_method" xorm:"int 'check_method' notnull default(0) comment('检查方法 0=tcpchk,检测端口是否握手成功')"`
	CheckTimeout           int    `json:"check_timeout" xorm:"int 'check_timeout' notnull default(0) comment('check超时时间,单位s')"`
	CheckInterval          int    `json:"check_interval" xorm:"int 'check_interval' notnull default(0) comment('检查间隔, 单位s')"`
	RoundType              int    `json:"round_type" xorm:"int 'round_type' notnull default(0) comment('轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash')"`
	IpList                 string `json:"ip_list" xorm:"varchar(2000) 'ip_list' notnull default('') comment('ip列表')"`
	WeightList             string `json:"weight_list" xorm:"varchar(2000) 'weight_list' notnull default('') comment('权重列表')"`
	ForbidList             string `json:"forbid_list" xorm:"varchar(2000) 'forbid_list' notnull default('') comment('禁用ip列表')"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" xorm:"int 'upstream_connect_timeout' notnull default(0) comment('建立连接超时, 单位s')"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" xorm:"int 'upstream_header_timeout' notnull default(0) comment('获取header超时, 单位s')"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" xorm:"int 'upstream_idle_timeout' notnull default(0) comment('链接最大空闲时间, 单位s')"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" xorm:"int 'upstream_max_idle' notnull default(0) comment('最大空闲链接数')"`
}

func (s *ServiceLoadBalance) FindIpList() []string {
	return strings.Split(s.IpList, ",")
}

func (s *ServiceLoadBalance) GetId() uint64 {
	return s.Id
}

func (m *ServiceLoadBalance) Find() (err error) {
	b, err := db.GetDBDriver().Get(m)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("数据不存在，导致查询错误")
	}
	return nil
}

// 以自身为条件，判断是否存在
func (m *ServiceLoadBalance) Exist(tag string) error {
	ok, err := db.GetDBDriver().Exist(m)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}
