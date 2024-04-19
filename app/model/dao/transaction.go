package dao

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/app/tools/db"
)

type ServiceInterface interface {
	GetId() uint64
}

// 事务的方式存入多个数据库表
func TransactionSaveServiceAll(serviceAddHttpRequest any, loadType int) error {
	// 注：要先调session.Begin()，最后要调session.Commit(),Rollback()一般可以不用调，调用session.Close的时候，如果没调过session.Commit()，则Rollback()会被自动调。
	session := db.GetDBDriver().NewSession()
	defer session.Close()
	//	添加一个开始
	if err := session.Begin(); err != nil {
		return err
	}
	// service_info 插入
	serviceInfo := &ServiceInfo{LoadType: loadType}
	copier.Copy(serviceInfo, serviceAddHttpRequest)
	if _, err1 := session.Insert(serviceInfo); err1 != nil {
		return errors.New(err1.Error())
	}
	// Rule 插入
	switch loadType {
	case public.LoadTypeHttp:
		// service_http_rule
		serviceHttpRule := &ServiceHttpRule{ServiceId: serviceInfo.Id}
		copier.Copy(serviceHttpRule, serviceAddHttpRequest)
		if _, err1 := session.Insert(serviceHttpRule); err1 != nil {
			return errors.New(err1.Error())
		}
	case public.LoadTypeTcp:
		serviceTcpRule := &ServiceTcpRule{ServiceId: serviceInfo.Id}
		copier.Copy(serviceTcpRule, serviceAddHttpRequest)
		if _, err1 := session.Insert(serviceTcpRule); err1 != nil {
			return errors.New(err1.Error())
		}
	case public.LoadTypeGrpc:
		serviceGrpcRule := &ServiceGrpcRule{ServiceId: serviceInfo.Id}
		copier.Copy(serviceGrpcRule, serviceAddHttpRequest)
		if _, err1 := session.Insert(serviceGrpcRule); err1 != nil {
			return errors.New(err1.Error())
		}
	default:
		return errors.New("load type 错误")
	}
	// service_access_control
	serviceAccessControl := &ServiceAccessControl{ServiceId: serviceInfo.Id}
	copier.Copy(serviceAccessControl, serviceAddHttpRequest)
	if _, err1 := session.Insert(serviceAccessControl); err1 != nil {
		return errors.New(err1.Error())
	}
	// service_load_balance
	serviceLoadBalance := &ServiceLoadBalance{ServiceId: serviceInfo.Id}
	copier.Copy(serviceLoadBalance, serviceAddHttpRequest)
	if _, err1 := session.Insert(serviceLoadBalance); err1 != nil {
		return errors.New(err1.Error())
	}
	session.Commit()
	return nil
}

// 更新http数据
// func TransactionUpdateHttp(serviceUpdateHttp *dto.ServiceUpdateHttpRequest) error {
// 	// 注：要先调session.Begin()，最后要调session.Commit(),Rollback()一般可以不用调，调用session.Close的时候，如果没调过session.Commit()，则Rollback()会被自动调。
// 	session := db.GetDBDriver().NewSession()
// 	defer session.Close()
// 	//	添加一个开始
// 	if err := session.Begin(); err != nil {
// 		return err
// 	}
// 	// service_info 插入
// 	serviceInfo := &ServiceInfo{Id: serviceUpdateHttp.ID}
// 	copier.Copy(serviceInfo, serviceUpdateHttp)
// 	if _, err1 := session.ID(serviceUpdateHttp.ID).AllCols().Update(serviceInfo); err1 != nil {
// 		return errors.New(err1.Error())
// 	}
// 	// service_http_rule
// 	serviceHttpRule := &ServiceHttpRule{ServiceId: serviceUpdateHttp.ID}
// 	copier.Copy(serviceHttpRule, serviceUpdateHttp)
// 	if _, err1 := session.Where("service_id = ?", serviceUpdateHttp.ID).AllCols().Update(serviceHttpRule); err1 != nil {
// 		return errors.New(err1.Error())
// 	}
// 	// service_access_control
// 	serviceAccessControl := &ServiceAccessControl{ServiceId: serviceUpdateHttp.ID}
// 	copier.Copy(serviceHttpRule, serviceUpdateHttp)
// 	if _, err1 := session.Where("service_id = ?", serviceUpdateHttp.ID).AllCols().Update(serviceAccessControl); err1 != nil {
// 		return errors.New(err1.Error())
// 	}
// 	// service_load_balance
// 	serviceLoadBalance := &ServiceLoadBalance{ServiceId: serviceUpdateHttp.ID}
// 	copier.Copy(serviceLoadBalance, serviceUpdateHttp)
// 	if _, err1 := session.Where("`service_id` = ?", serviceUpdateHttp.ID).AllCols().Update(serviceLoadBalance); err1 != nil {
// 		return errors.New(err1.Error())
// 	}
// 	session.Commit()
// 	return nil
// }

// id 将作为主键 和 service_id 用于查询
//
//	方法流程：通过 id 查询出 ServiceDetail，根据 ServiceDetail.Info.LoadType 获取需要更新的 rule
//	根据 传入第二个参数 与 ServiceDetail 中个字段匹配，匹配完成之后逐个执行 全写入
func TransactionUpdateAll(id uint64, serviceUpdate any) error {
	serviceDetail := &ServiceDetail{Info: &ServiceInfo{Id: id}}
	if err := serviceDetail.FindAll(); err != nil {
		return errors.New(err.Error())
	}
	listService := []ServiceInterface{serviceDetail.Info, serviceDetail.AccessControl, serviceDetail.LoadBalance}
	switch serviceDetail.Info.LoadType {
	case public.LoadTypeHttp:
		listService = append(listService, serviceDetail.HttpRule)
	case public.LoadTypeTcp:
		listService = append(listService, serviceDetail.TcpRule)
	case public.LoadTypeGrpc:
		listService = append(listService, serviceDetail.GrpcRule)
	}
	session := db.GetDBDriver().NewSession()
	defer session.Close()
	//	添加一个开始
	if err := session.Begin(); err != nil {
		return err
	}
	// 循环 更新
	for _, value := range listService {
		copier.Copy(value, serviceUpdate)
		if _, err1 := session.ID(value.GetId()).AllCols().Update(value); err1 != nil {
			return errors.New(err1.Error())
		}
	}
	session.Commit()
	return nil
}
