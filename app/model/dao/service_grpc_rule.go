package dao

import (
	"errors"

	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/app/tools/db"
)

// grpc网关路由匹配表，所有操作通过 ServiceId 字段进行，必须初始化
type ServiceGrpcRule struct {
	Id             uint64 `json:"id" xorm:"bigint pk notnull autoincr 'id' comment('自增主键')"`
	ServiceId      uint64 `json:"service_id" xorm:"bigint unique index notnull default(0) 'service_id'"`
	Port           int    `json:"port" xorm:"int 'port' notnull comment('端口')"`
	HeaderTransfor string `json:"headerTransfor" xorm:"varchar(5000) notnull default('') 'header_transfor' comment('header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔')"`
}

func (s *ServiceGrpcRule) GetId() uint64 {
	return s.Id
}

func (m *ServiceGrpcRule) Find() (err error) {
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
func (m *ServiceGrpcRule) Exist(tag string) error {
	ok, err := db.GetDBDriver().Exist(m)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}
