package dao

import (
	"errors"
	"time"

	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
)

// 网关租户表
type App struct {
	Id       uint64    `json:"id" xorm:"bigint pk autoincr notnull 'id' comment('自增id')"`
	AppId    string    `json:"app_id" xorm:"varchar(255) notnull default('') unique 'app_id' comment('租户id')"`
	Name     string    `json:"name" xorm:"varchar(255) 'name' notnull unique default('租户名称')"`
	Secret   string    `json:"secret" xorm:"varchar(255) 'secret' notnull default('') comment('密钥')"`
	WhiteIps string    `json:"white_ips" xorm:"varchar(1000) 'white_ips' notnull default('') comment('ip白名单，支持前缀匹配')"`
	Qpd      uint64    `json:"qpd" xorm:"bigint 'qpd' notnull default(0) comment('日请求量限制')"`
	Qps      uint64    `json:"qps" xorm:"bigint 'qps' notnull default(0) comment('每秒请求量限制')"`
	CreateAt time.Time `json:"create_at" xorm:"created 'create_at' comment('添加时间')"`
	UpdateAt time.Time `json:"update_at" xorm:"updated 'update_at' comment('更新时间')"`
	DeleteAt time.Time `json:"delete_at" xorm:"deleted 'delete_at' comment('删除时间')"`
	IsDelete int       `json:"is_delete" xorm:"int 'is_delete' notnull default(0) comment('是否删除 1=删除')"`
}

func (m *App) Find() (err error) {
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
func (m *App) Exist(tag string) error {
	ok, err := GetDBDriver().Exist(m)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}

func (a *App) PageList(param *dto.ServiceListRequest) (list []App, totle uint64, err error) {
	list = []App{}
	query := GetDBDriver().NewSession()
	defer query.Close()
	if param.Info != "" {
		query = query.Where("app_id like ?", "%"+param.Info+"%").Or("name like ?", "%"+param.Info+"%")
	}
	err = query.Desc("id").Limit(param.PageSize, param.PageNo-1).Find(&list)
	return list, uint64(len(list)), err
}

func (a *App) Delete() error {
	i, err := GetDBDriver().ID(a.Id).Delete(a)
	if i < 1 {
		return errors.New("删除失败，无此记录")
	}
	return err
}

func (a *App) Save() error {
	i, err := GetDBDriver().Insert(a)
	if i < 1 {
		return errors.New("数据存入失败")
	}
	return err
}

func (a *App) Update() error {
	i, err := GetDBDriver().ID(a.Id).AllCols().Update(a)
	if i < 1 {
		return errors.New("数据写入失败")
	}
	return err
}

func (a *App) GetTatle() (int64, error) {
	i, err := GetDBDriver().Count(a)
	if err != nil {
		return 0, err
	}
	return i, nil
}
