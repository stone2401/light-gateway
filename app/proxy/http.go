package proxy

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/stone2401/light-gateway-kernel/pcore"
	"github.com/stone2401/light-gateway-kernel/pkg/zlog"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/app/tools/db"
	"github.com/stone2401/light-gateway/app/tools/etcd"
	"github.com/stone2401/light-gateway/config"
	"go.uber.org/zap/zapcore"
)

var onecHttpProxy *HttpProxy
var onecHttpProxyOnce sync.Once

func GetHttpProxy() *HttpProxy {
	onecHttpProxyOnce.Do(func() {
		onecHttpProxy = NewHttpProxy()
	})
	return onecHttpProxy
}

type (
	Counter *pcore.Counter
	Limiter *pcore.Limiter
	Fuse    *pcore.FuseEntry
)

type HttpProxy struct {
	proxyMap      map[string]*pcore.Engine
	middlewareMap map[string]map[string]any
	deleteMap     map[string]struct{}
}

func NewHttpProxy() *HttpProxy {
	balanc := etcd.GetMonitor().Register(strconv.Itoa(config.Config.Cluster.Port), pcore.LoadBalanceRandom)
	httpEngin := pcore.NewEngine(balanc)

	balanc2 := etcd.GetMonitor().Register(strconv.Itoa(config.Config.Cluster.SSLPort), pcore.LoadBalanceRandom)
	httpsEngin := pcore.NewEngine(balanc2)
	go func() {
		err := httpEngin.Start(":" + strconv.Itoa(config.Config.Cluster.Port))
		if err != nil {
			zlog.Zlog().Error("start http proxy error", zapcore.Field{Key: "err", Type: zapcore.StringType, String: err.Error()})
			return
		}
		zlog.Zlog().Info("start http proxy", zapcore.Field{Key: "port", Type: zapcore.StringType, String: strconv.Itoa(config.Config.Cluster.Port)})
	}()
	go func() {
		// pwd
		pwd, _ := os.Getwd()
		err := httpsEngin.StartTls(":"+strconv.Itoa(config.Config.Cluster.SSLPort), pwd+config.Config.Cluster.SSLCertFile, pwd+config.Config.Cluster.SSLKeyFile)
		if err != nil {
			zlog.Zlog().Error("start https proxy error", zapcore.Field{Key: "err", Type: zapcore.StringType, String: err.Error()})
			return
		}
		zlog.Zlog().Info("start https proxy", zapcore.Field{Key: "port", Type: zapcore.StringType, String: strconv.Itoa(config.Config.Cluster.SSLPort)})
	}()
	return &HttpProxy{
		proxyMap: map[string]*pcore.Engine{
			strconv.Itoa(config.Config.Cluster.Port):    httpEngin,
			strconv.Itoa(config.Config.Cluster.SSLPort): httpsEngin,
		},
		middlewareMap: map[string]map[string]any{
			strconv.Itoa(config.Config.Cluster.Port): {
				"Balance": balanc,
			},
			strconv.Itoa(config.Config.Cluster.SSLPort): {
				"Balance": balanc2,
			},
		},
		deleteMap: map[string]struct{}{},
	}
}

func (h *HttpProxy) Register(info *dto.ServiceAddHttpRequest) error {
	var err error
	err = nil
	defer func() {
		if pia := recover(); pia != nil {
			err = pia.(error)
		}
	}()

	// 初始化中间件
	// 1. 负载均衡器
	rule := ""
	if info.ServiceHttpRule.RuleType != public.HTTPRuleTypeDomainURL {
		rule = info.ServiceHttpRule.Rule
	} else {
		// 端口占用
		if _, ok := h.proxyMap[strconv.Itoa(info.ServiceHttpRule.Port)]; ok {
			return errors.New("端口占用")
		}
	}
	// 判断是否存在
	if _, ok := h.deleteMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]; ok {
		h.Update(nil, info)
		return nil
	}
	// 1. 开关
	status := NewStatus()
	headler := []pcore.Handler{status.StatusHeadler}
	if info.Status == 1 {
		status.Open()
	}
	balance := etcd.GetMonitor().Register(strconv.Itoa(info.ServiceHttpRule.Port)+rule, pcore.GetLoadBalance(info.ServiceLoadBalance.RoundType))
	if info.ServiceLoadBalance.IpList != "" {
		ips := strings.Split(info.ServiceLoadBalance.IpList, "\n")
		weights := strings.Split(info.ServiceLoadBalance.WeightList, "\n")
		for i := 0; i < len(ips); i++ {
			weight, err := strconv.Atoi(weights[i])
			if err != nil {
				return err
			}
			balance.AddNode(ips[i], weight)
		}
	}
	// 2. 计数器
	counter := pcore.NewCounter(info.ServiceName, 30)
	// 3. 限流
	// limiter := pcore.NewLimiter(1000)
	// 4. 熔断
	// fuse := pcore.NewFuseEntry(strconv.Itoa(info.ServiceHttpRule.Port)+rule, 0, 100, 0.5)

	// 6. 黑白名单
	if info.ServiceAccessControl.BlackList != "" || info.ServiceAccessControl.WhiteList != "" {
		access := NewAccessControl()
		access.AddBlackList(info.ServiceAccessControl.BlackList)
		access.AddWhiteList(info.ServiceAccessControl.WhiteList)
		headler = append(headler, access.AccessControlHeadler)
		h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["AccessControl"] = access
	}
	// headler = append(headler, counter.CounterHandler, limiter.ProxyHandler, fuse.FuseHandler)
	headler = append(headler, counter.CounterHandler)
	// 7. header 重写
	if info.ServiceHttpRule.HeaderTransfor != "" {
		header := NewResetHeader()
		header.Set(info.ServiceHttpRule.HeaderTransfor)
		headler = append(headler, header.ResetHeader)
		h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["Header"] = header
	}

	// 8. uri 重写
	if info.ServiceHttpRule.UrlRewrite != "" {
		rewrite := NewUrlRewrite()
		rewrite.Add(info.ServiceHttpRule.UrlRewrite)
		headler = append(headler, rewrite.RewriteHandler)
		h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["UrlRewrite"] = rewrite
	}

	// 初始化中间件
	h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule] = map[string]any{
		"Counter": counter,
		// "Limiter": limiter,
		// "Fuse":    fuse,
		"Balance": balance,
		"Status":  status,
	}
	// end. 代理
	if info.ServiceHttpRule.RuleType == public.HTTPRuleTypeDomainURL {
		h.proxyMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule] = pcore.NewEngine(balance)
		h.proxyMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule].Use(headler...)
		h.proxyMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule].Start(":" + strconv.Itoa(info.ServiceHttpRule.Port))
	} else {
		h.proxyMap[strconv.Itoa(info.ServiceHttpRule.Port)].Register(info.ServiceHttpRule.Rule, balance, headler...)
	}
	GetSurveillant().Register(info.ServiceName, counter)
	return err
}

func (h *HttpProxy) Update(oldInfo *dao.ServiceDetail, info *dto.ServiceAddHttpRequest) error {
	rule := ""
	if info.ServiceHttpRule.RuleType != public.HTTPRuleTypeDomainURL {
		rule = info.ServiceHttpRule.Rule
	}
	// 规则发生改变
	if oldInfo != nil && info.ServiceHttpRule.Rule != oldInfo.HttpRule.Rule {
		if status, ok := h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["Status"]; ok {
			status.(*Status).Close()
		}
		h.Register(info)
		return nil
	}
	// 是否需要关闭
	if info.Status == 0 {
		h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["Status"].(*Status).Close()
	} else {
		h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["Status"].(*Status).Open()
		h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["Counter"].(*pcore.Counter).Reset()
		GetSurveillant().Register(info.ServiceName, h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["Counter"].(*pcore.Counter))
	}

	// 2. 计数器
	if oldInfo == nil || oldInfo.Info.ServiceName != info.ServiceName {
		GetSurveillant().Rename(oldInfo.Info.ServiceName, info.ServiceName)
	}
	// 6. 黑白名单
	if oldInfo == nil || info.ServiceAccessControl.BlackList != oldInfo.AccessControl.BlackList || info.ServiceAccessControl.WhiteList != oldInfo.AccessControl.WhiteList {
		access := h.middlewareMap[strconv.Itoa(info.Port)+rule]["AccessControl"].(*AccessControl)
		access.Reset()
		access.AddBlackList(info.ServiceAccessControl.BlackList)
		access.AddWhiteList(info.ServiceAccessControl.WhiteList)
	}

	// 7. header 重写
	if oldInfo == nil || info.ServiceHttpRule.HeaderTransfor != oldInfo.HttpRule.HeaderTransfor {
		header := h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["Header"].(*ResetHeader)
		header.Set(info.ServiceHttpRule.HeaderTransfor)
	}

	// 8. uri 重写
	if oldInfo == nil || info.ServiceHttpRule.UrlRewrite != oldInfo.HttpRule.UrlRewrite {
		rewrite := h.middlewareMap[strconv.Itoa(info.ServiceHttpRule.Port)+rule]["UrlRewrite"].(*UrlRewrite)
		rewrite.Reset()
		rewrite.Add(info.ServiceHttpRule.UrlRewrite)
	}
	return nil
}

func (h *HttpProxy) Remove(info *dao.ServiceDetail) error {
	rule := ""
	if info.HttpRule.RuleType != public.HTTPRuleTypeDomainURL {
		rule = info.HttpRule.Rule
	}
	if h.middlewareMap[strconv.Itoa(info.HttpRule.Port)+rule]["Status"] != nil {
		h.middlewareMap[strconv.Itoa(info.HttpRule.Port)+rule]["Status"].(*Status).Close()
		h.deleteMap[strconv.Itoa(info.HttpRule.Port)+rule] = struct{}{}
	}
	return nil
}

func (h *HttpProxy) Init() {
	// 获取所有服务
	services := []*dao.ServiceInfo{}
	err := db.GetDBDriver().Table(&dao.ServiceInfo{}).
		Where("load_type = ? and status = ?", public.LoadTypeHttp, public.StatusUp).Find(&services)
	if err != nil {
		return
	}
	for _, info := range services {
		detail := dao.ServiceDetail{Info: info}
		err := detail.FindAll()
		if err != nil {
			continue
		}
		if detail.HttpRule.Port == 0 {
			continue
		}
		h.Register(&dto.ServiceAddHttpRequest{
			ServiceName: detail.Info.ServiceName,
			ServiceDesc: detail.Info.ServiceDesc,
			Status:      detail.Info.Status,
			ServiceHttpRule: dto.ServiceHttpRule{
				Port:           detail.HttpRule.Port,
				RuleType:       detail.HttpRule.RuleType,
				Rule:           detail.HttpRule.Rule,
				NeedHttps:      detail.HttpRule.NeedHttps,
				NeedStripUrl:   detail.HttpRule.NeedStripUrl,
				NeedWebsocket:  detail.HttpRule.NeedWebsocket,
				UrlRewrite:     detail.HttpRule.UrlRewrite,
				HeaderTransfor: detail.HttpRule.HeaderTransfor,
			},
			ServiceAccessControl: dto.ServiceAccessControl{
				OpenAuth:          detail.AccessControl.OpenAuth,
				BlackList:         detail.AccessControl.BlackList,
				WhiteList:         detail.AccessControl.WhiteList,
				ClientipFlowLimit: detail.AccessControl.ClientipFlowLimit,
				ServiceFlowLimit:  detail.AccessControl.ServiceFlowLimit,
			},
			ServiceLoadBalance: dto.ServiceLoadBalance{
				RoundType:              detail.LoadBalance.RoundType,
				IpList:                 detail.LoadBalance.IpList,
				WeightList:             detail.LoadBalance.WeightList,
				ForbidList:             detail.LoadBalance.ForbidList,
				UpstreamConnectTimeout: detail.LoadBalance.UpstreamConnectTimeout,
				UpstreamHeaderTimeout:  detail.LoadBalance.UpstreamHeaderTimeout,
				UpstreamIdleTimeout:    detail.LoadBalance.UpstreamIdleTimeout,
				UpstreamMaxIdle:        detail.LoadBalance.UpstreamMaxIdle,
			},
		})
	}
}
