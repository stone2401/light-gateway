package dao

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
)

// ServiceDetail 任何操作都是通过 Info 字段进行的，所以 Info 字段必须初始化
type ServiceDetail struct {
	Info          *ServiceInfo          `json:"info"`
	HttpRule      *ServiceHttpRule      `json:"http_rule"`
	TcpRule       *ServiceTcpRule       `json:"tcp_rule"`
	GrpcRule      *ServiceGrpcRule      `json:"grpc_rule"`
	LoadBalance   *ServiceLoadBalance   `json:"load_balance"`
	AccessControl *ServiceAccessControl `json:"access_control"`
}

type Rule interface {
	Find() error
	Exist(string) error
}

// 以自身为查询条件查询，查询Rule部分字段
func (a *ServiceDetail) FindRule() error {
	// 负载类型 0=http 1=tcp 2=grpc
	switch a.Info.LoadType {
	case public.LoadTypeHttp:
		httpRule := &ServiceHttpRule{ServiceId: a.Info.Id}
		err := httpRule.Find()
		if err != nil {
			return err
		}
		a.HttpRule = httpRule
	case public.LoadTypeTcp:
		tcpRule := &ServiceTcpRule{ServiceId: a.Info.Id}
		err := tcpRule.Find()
		if err != nil {
			return err
		}
		a.TcpRule = tcpRule
	case public.LoadTypeGrpc:
		grpcRule := &ServiceGrpcRule{ServiceId: a.Info.Id}
		err := grpcRule.Find()
		if err != nil {
			return err
		}
		a.GrpcRule = grpcRule
	}
	return nil
}

// 以自身为条件，查询此条信息等等全部字段
func (s *ServiceDetail) FindAll() (err error) {
	// 查询 info 字段
	err = s.Info.Find()
	if err != nil {
		return err
	}
	// 查询 rule 部分
	err = s.FindRule()
	if err != nil {
		return err
	}
	// 查询 ServiceAccessControl 部分
	err = s.FindAccessControl()
	if err != nil {
		return err
	}
	// 查询 ServiceLoadBalance 部分
	err = s.FindLoadBalance()
	if err != nil {
		return err
	}
	return nil
}

// 在 detail 中抽离出 response 数据
func (s *ServiceDetail) DrawServiecResponse() (any, error) {
	// 因为 findall 方法只需要 id 字段所有随便操作
	if err := s.FindAll(); err != nil {
		return nil, errors.New(err.Error())
	}
	listService := []ServiceInterface{s.Info, s.AccessControl, s.LoadBalance}
	// 获取 rule 字段
	var serviceDetail any
	switch s.Info.LoadType {
	case public.LoadTypeHttp:
		serviceDetail = &dto.ServiceUpdateHttpRequest{}
		listService = append(listService, s.HttpRule)
	case public.LoadTypeTcp:
		serviceDetail = &dto.ServiceUpdateTcpRequest{}
		listService = append(listService, s.TcpRule)
	case public.LoadTypeGrpc:
		serviceDetail = &dto.ServiceUpdateGrpcRequest{}
		listService = append(listService, s.GrpcRule)
	}
	for _, value := range listService {
		copier.Copy(serviceDetail, value)
	}
	return serviceDetail, nil
}

// 以自身为条件，查询此条信息等等全部字段
func (s *ServiceDetail) FindLoadBalance() (err error) {
	// 查询 ServiceLoadBalance 部分
	loadBalance := &ServiceLoadBalance{ServiceId: s.Info.Id}
	err = loadBalance.Find()
	if err != nil {
		return err
	}
	s.LoadBalance = loadBalance
	return nil
}

// 以自身为条件，查询此条信息等等全部字段
func (s *ServiceDetail) FindAccessControl() (err error) {
	// 查询 ServiceAccessControl 部分
	access := &ServiceAccessControl{ServiceId: s.Info.Id}
	err = access.Find()
	if err != nil {
		return err
	}
	s.AccessControl = access
	return nil
}

func (s *ServiceDetail) GetServiceAddr(ctx *gin.Context) string {
	var builder strings.Builder
	addr := ctx.Request.Host
	builder.WriteString(strings.Split(addr, ":")[0] + ":")
	switch s.Info.LoadType {
	case public.LoadTypeHttp:
		if s.HttpRule.RuleType == public.HTTPRuleTypeDomainURL {
			builder.Reset()
			builder.WriteString(s.HttpRule.Rule)
		} else if s.HttpRule.RuleType == public.HTTPRuleTypePrefixURL {
			builder.WriteString(strconv.Itoa(s.HttpRule.Port) + s.HttpRule.Rule)
		}
	case public.LoadTypeTcp:
		builder.WriteString(strconv.Itoa(s.TcpRule.Port))
	case public.LoadTypeGrpc:
		builder.WriteString(strconv.Itoa(s.GrpcRule.Port))
	}
	// 在 ctx 中获取到的地址是 0.0.0.0:8080，需要进行替换
	return builder.String()
}
