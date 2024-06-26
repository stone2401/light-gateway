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
	Id          uint64    `json:"id" uri:"id" xorm:"bigint pk autoincr 'id' comment('自增主键')" rule:"notnull" label:"ID" form:"id"`
	LoadType    int       `json:"loadType" xorm:"tinyint(4) 'load_type' notnull default(0) comment('负载类型 0=http 1=tcp 2=grpc')"`
	ServiceName string    `json:"serviceName" xorm:"varchar(255) 'service_name' notnull unique comment('服务名称 6-128 数字字母下划线')"`
	ServiceDesc string    `json:"serviceDesc" xorm:"varchar(255) 'service_desc' default('') comment('服务描述')"`
	Status      int       `json:"status" xorm:"tinyint(4) 'status' notnull default(0) comment('开关 0 关 1 开')"`
	CreateAt    time.Time `json:"create_at" xorm:"created 'create_at' comment('创建时间')"`
	UpdateAt    time.Time `json:"update_at" xorm:"updated 'update_at' comment('更新时间')"`
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
func (s *ServiceInfo) PageList(param *dto.ServiceListRequest) (list []ServiceInfo, totle int, err error) {
	list = []ServiceInfo{}
	query := db.GetDBDriver().NewSession()
	defer query.Close()
	// 如果 info 存在则进行模糊匹配
	if param.Info != "" {
		query = query.Where("service_name like ?", "%"+param.Info+"%").Or("service_desc like ?", "%"+param.Info+"%")
	}
	err = query.Desc("id").Limit(param.PageSize, param.Page-1).Find(&list)
	if err != nil {
		return nil, 0, err
	}
	count, err := query.Count(s)
	if err != nil {
		return nil, 0, err
	}
	return list, int(count), err
}

// 删除操作，需要初始化 id
func (s *ServiceInfo) Delete() error {
	ok, err := db.GetDBDriver().Where("id = ?", s.Id).Get(s)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("数据不存在" + public.EndMark)
	}
	var i int64
	switch s.LoadType {
	case public.LoadTypeHttp:
		i, err = db.GetDBDriver().Where("service_id = ?", s.Id).Delete(&ServiceHttpRule{ServiceId: s.Id})
	case public.LoadTypeGrpc:
		i, err = db.GetDBDriver().Where("service_id = ?", s.Id).Delete(&ServiceGrpcRule{ServiceId: s.Id})
	case public.LoadTypeTcp:
		i, err = db.GetDBDriver().Where("service_id = ?", s.Id).Delete(&ServiceTcpRule{ServiceId: s.Id})
	}
	if i == 0 {
		return errors.New("删除失败，请联系管理员" + public.EndMark)
	}
	if err != nil {
		return err
	}
	db.GetDBDriver().Where("id = ?", s.Id).Delete(&ServiceInfo{})
	return nil
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
