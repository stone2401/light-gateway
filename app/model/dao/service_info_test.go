package dao

import (
	"testing"
)

func TestServiceInfo_PageList(t *testing.T) {
	// param := &dto.ServiceListRequest{Info: "GRPC", PageNo: 1, PageSize: 5}
	// service := ServiceInfo{}
	// t.Log(service.PageList(param))
	info := &ServiceInfo{Id: 58}
	info.Find()
	t.Logf("%+v", info)

}

func TestUpdate(t *testing.T) {
	param := &ServiceInfo{Id: 61}
	i, err := GetDBDriver().ID(param.Id).Unscoped().Cols("delete_at").Update(param)
	t.Logf("恢复条数:%d 错误日志:%+v, params: %+v", i, err, param)
}
