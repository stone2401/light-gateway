package dao

import (
	"errors"

	"github.com/stone2401/light-gateway/app/public"
)

// tcp网关路由匹配表，所有操作通过 ServiceId 字段进行，必须初始化
type ServiceTcpRule struct {
	Id        uint64 `json:"id" xorm:"bigint pk notnull autoincr 'id' comment('自增主键')"`
	ServiceId uint64 `json:"service_id" xorm:"bigint unique index notnull default(0) 'service_id'"`
	Port      int    `json:"port" xorm:"int 'port' notnull comment('端口')"`
}

func (s *ServiceTcpRule) GetId() uint64 {
	return s.Id
}

func (m *ServiceTcpRule) Find() (err error) {
	b, err := DBEngine.Get(m)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("数据不存在，导致查询错误")
	}
	return nil
}

// 以自身为条件，判断是否存在
func (m *ServiceTcpRule) Exist(tag string) error {
	ok, err := DBEngine.Exist(m)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}
