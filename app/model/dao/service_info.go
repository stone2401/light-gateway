package dao

import (
	"time"

	"github.com/pkg/errors"

	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/app/tools/db"
)

// 网关基本信息表，所有操作通过 Id  字段进行，必须初始化
type ServiceInfo struct {
	Id          uint64    `json:"id" xorm:"bigint pk autoincr 'id' comment('自增主键')" rule:"notnull" label:"ID" form:"id"`
	LoadType    int       `json:"load_type" xorm:"tinyint(4) 'load_type' notnull default(0) comment('负载类型 0=http 1=tcp 2=grpc')"`
	ServiceName string    `json:"service_name" xorm:"varchar(255) 'service_name' notnull unique comment('服务名称 6-128 数字字母下划线')"`
	ServiceDesc string    `json:"service_desc" xorm:"varchar(255) 'service_desc' default('') comment('服务描述')"`
	CreateAt    time.Time `json:"create_at" xorm:"created 'create_at' comment('创建时间')"`
	UpdateAt    time.Time `json:"update_at" xorm:"updated 'update_at' comment('更新时间')"`
	DeleteAt    time.Time `json:"delete_at" xorm:"deleted 'delete_at' comment('删除时间')"`
	IsDelete    int       `json:"is_delete" xorm:"int 'is_delete' notnull default(0) comment('是否删除 1=删除')"`
}

func (m *ServiceInfo) Find() (err error) {
	b, err := db.GetDBDriver().Get(m)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("数据不存在，导致查询错误")
	}
	return nil
}

// 获取 []ServiceInfo
func (s *ServiceInfo) PageList(param *dto.ServiceListRequest) (list []ServiceInfo, totle uint64, err error) {
	list = []ServiceInfo{}
	query := db.GetDBDriver().NewSession()
	defer query.Close()
	// 如果 info 存在则进行模糊匹配
	if param.Info != "" {
		query = query.Where("service_name like ?", "%"+param.Info+"%").Or("service_desc like ?", "%"+param.Info+"%")
	}
	err = query.Desc("id").Limit(param.PageSize, param.PageNo-1).Find(&list)
	// err = query.Asc("id").Limit(param.PageSize, param.PageNo-1).Find(&list)
	return list, uint64(len(list)), err
}

// 删除操作，需要初始化 id
func (s *ServiceInfo) Delete() error {
	i, err := db.GetDBDriver().ID(s.Id).Delete(s)
	if i < 1 {
		return errors.New("删除失败，无此记录")
	}
	return err
}

// 以自身为条件，判断是否存在
func (s *ServiceInfo) Exist(tag string) error {
	ok, err := db.GetDBDriver().Exist(s)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}

func (s *ServiceInfo) GetId() uint64 {
	return s.Id
}

func (s *ServiceInfo) GetTatle() (int64, error) {
	i, err := db.GetDBDriver().Count(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (s *ServiceInfo) GroupByLoadType() ([]dto.ServiceStatAllItemResponse, error) {
	statList := []dto.ServiceStatAllItemResponse{}
	for index, name := range public.LoadTypeSlice {
		value, err := db.GetDBDriver().Cols("load_type").Where("load_type = ?", index).Count(s)
		if err != nil {
			return nil, err
		}
		stat := dto.ServiceStatAllItemResponse{Name: name, Value: uint64(value)}
		statList = append(statList, stat)
	}

	return statList, nil
}
