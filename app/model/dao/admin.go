package dao

import (
	"errors"
	"log"
	"time"

	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
)

// 管理员表
type Admin struct {
	Id       uint64    `json:"id" xorm:"bigint pk notnull autoincr 'id' comment('自增id')"`
	UserName string    `json:"user_name" xorm:"varchar(255) notnull 'user_name' unique comment('用户名')"`
	Salt     string    `json:"salt" xorm:"varchar(255) notnull 'salt' comment('盐')"`
	Password string    `json:"password" xorm:"varchar(255) notnull 'password' comment('密码')"`
	CreateAt time.Time `json:"create_at" xorm:"created notnull  'create_at' comment('创建时间')"`
	UpdateAt time.Time `json:"update_at" xorm:"updated notnull 'update_at' comment('更新时间')"`
	DeleteAt time.Time `json:"delete_at" xorm:"deleted 'delete_at' comment('删除时间')"`
	IsDelete int       `json:"is_delete" gorm:"int 'is_delete' notnull default(0) comment('是否删除')"`
}

// 密码对比，以传入的param.Username为查询字段，也会以自身为查询字段
func (a *Admin) LoginCheck(param *dto.AdminLoginRequest) error {
	a.UserName = param.Username
	// 查询
	err := a.Find()
	if err != nil {
		log.Panicln(err)
	}
	// 获取加密 password，之后与数据库中密码对比
	saltPassword := public.GenSaltPassword(a.Salt, param.Password)
	if a.Password != saltPassword {
		return errors.New("密码错误，请重新输入")
	}
	return nil
}

// 以自身为条件，以参数为更新
func (a *Admin) Update(newAdmin *Admin) error {
	_, err := DBEngine.ID(a.Id).Update(newAdmin)
	return err
}

func (m *Admin) Find() (err error) {
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
func (m *Admin) Exist(tag string) error {
	ok, err := DBEngine.Exist(m)
	if ok {
		return errors.New(tag + "已存在" + public.EndMark)
	}
	return err
}
