package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/proxy"
	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/config"
)

func RegisterService(router *gin.RouterGroup) {
	router.GET("/", ServiceList)
	router.GET("/service_stat", ServiceStat)
	router.GET("/:id", ServiceDetail)
	router.DELETE("/:id", ServiceDelete)

	router.POST("/http", ServiceAddHttp)
	router.POST("/tcp", ServiceAddTcp)
	router.POST("/grpc", ServiceAddGrpc)
	router.PUT("/http/:id", ServiceUpdateHttp)
	router.PUT("/tcp/:id", ServiceUpdateTcp)
	router.PUT("/grpc/:id", ServiceUpdateGrpc)
}

// @Summary 服务列表
// @Schemes
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param info query string false "关键词"
// @Param page query int true "页数"
// @Param pageSize query int true "个数"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceListResponse} "success"
// @Router /api/v1/service/list [get]
func ServiceList(ctx *gin.Context) {
	// 获取 ServiceListRequest 并绑定值
	parmas := &dto.ServiceListRequest{}
	if err := public.Authenticator(ctx, parmas); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 查询操作
	serviceList := &dao.ServiceInfo{}
	list, total, err := serviceList.PageList(parmas)
	if err != nil {
		middleware.ResponseError(ctx, 1002, err)
		return
	}
	// 遍历获取需要的值
	responseList := []dto.ServiceListItem{}
	for _, item := range list {
		// 使用 detail 进行全部字段查询
		detail := &dao.ServiceDetail{Info: &item}
		// 得到全部 rule 字段
		err2 := detail.FindRule()
		if err2 != nil {
			middleware.ResponseError(ctx, 1002, err2)
		}
		// 得到全部 LoadBalance 字段
		err2 = detail.FindLoadBalance()
		if err2 != nil {
			middleware.ResponseError(ctx, 1002, err2)
		}
		responseItem := dto.ServiceListItem{
			ID:          item.Id,
			ServiceName: item.ServiceName,
			ServiceDesc: item.ServiceDesc,
			LoadType:    item.LoadType,
			ServiceAddr: detail.GetServiceAddr(ctx),
			QPS:         0,
			QPD:         0,
			Status:      item.Status,
			TotalNode:   len(detail.LoadBalance.FindIpList()),
		}
		responseList = append(responseList, responseItem)
	}
	middleware.ResponseSuccess(ctx, &dto.ServiceListResponse{Items: responseList, Total: uint64(total), Meta: dto.BaseMeta{
		TotalItems:   total,
		ItemCount:    len(list),
		ItemsPerPage: parmas.PageSize,
		TotalPages:   (0 + total) / parmas.PageSize,
		CurrentPage:  parmas.Page,
	}})
}

// @Summary 服务删除
// @Description 服务删除
// @Tags 服务管理
// @ID /service/service_delete
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param id query string true "关键词"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /service/:id [get]
func ServiceDelete(ctx *gin.Context) {
	serviceDelete := &dto.ServiceDeleteRequest{}
	// 绑定 ServiceInfo，之后就直接用这个对象进行查询
	if err := ctx.ShouldBindUri(serviceDelete); err != nil {
		middleware.ResponseError(ctx, 1003, err)
		return
	}
	// 进行删除
	serviceInfo := &dao.ServiceInfo{Id: serviceDelete.Id}
	serviceDetail := &dao.ServiceDetail{Info: &dao.ServiceInfo{Id: serviceDelete.Id}}
	err := serviceDetail.FindAll()
	if err != nil {
		middleware.ResponseError(ctx, 1004, err)
		return
	}
	proxy.GetHttpProxy().Remove(serviceDetail)
	err = serviceInfo.Delete()
	if err != nil {
		middleware.ResponseError(ctx, 1004, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功，ok!")
}

// @Summary 服务详情
// @Description 服务详情
// @Tags 服务管理
// @ID /service/service_detail
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param id query string true "关键词"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceUpdateHttpRequest} "success"
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceUpdateTcpRequest} "success"
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceUpdateGrpcRequest} "success"
// @Router /service/service_detail [get]
func ServiceDetail(ctx *gin.Context) {
	// 获取参数 id
	serviceInfo := &dao.ServiceInfo{}
	if err := ctx.BindUri(serviceInfo); err != nil {
		middleware.ResponseError(ctx, 1003, err)
		return
	}
	if err := public.Authenticator(ctx, serviceInfo); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 根据 id 获取 detail，调用 DrawServiecResponse 提取出 需要的字段
	detail := &dao.ServiceDetail{Info: serviceInfo}
	suhr, err := detail.DrawServiecResponse()
	if err != nil {
		middleware.ResponseError(ctx, 1002, err)
		return
	}
	middleware.ResponseSuccess(ctx, suhr)
}

// @Summary 服务统计
// @Description 服务统计
// @Tags 服务管理
// @ID /service/service_stat
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param id query string true "关键词"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceStatResponse} "success"
// @Router /service/service_stat [get]
func ServiceStat(ctx *gin.Context) {
	// 获取参数 id
	serviceInfo := &dao.ServiceInfo{}
	if err := public.Authenticator(ctx, serviceInfo); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 获取 数据
	today := []uint64{}
	for i := 0; i < time.Now().Hour(); i++ {
		today = append(today, 0)
	}
	yesterday := []uint64{}
	for i := 0; i < 24; i++ {
		yesterday = append(yesterday, 0)
	}
	serviceStat := &dto.ServiceStatResponse{Today: today, Yesterday: yesterday}
	// 发送
	middleware.ResponseSuccess(ctx, serviceStat)
}

// @Summary 添加http服务
// @Description 添加http服务
// @Tags 服务管理
// @ID /service/service_add_http
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.ServiceAddHttpRequest true "add"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /service/service_add_http [post]
func ServiceAddHttp(ctx *gin.Context) {
	// 获取参数并校验
	params := &dto.ServiceAddHttpRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 查询服务是否存在
	serviceName := &dao.ServiceInfo{ServiceName: params.ServiceName}
	err := serviceName.Exist("服务")
	if err != nil {
		middleware.ResponseError(ctx, 1005, err)
		return
	}
	// 查询域名或前后缀是否存在
	httpRule := &dao.ServiceHttpRule{RuleType: params.RuleType, Rule: params.Rule}
	err = httpRule.Exist("服务接入前缀或域名")
	if err != nil {
		middleware.ResponseError(ctx, 1005, err)
		return
	}
	// 检查端口是否被占用
	if httpRule.NeedHttps {
		params.Port = config.Config.Cluster.SSLPort
	} else {
		params.Port = config.Config.Cluster.Port
	}
	proxy.GetHttpProxy().Register(params)
	// 保存
	err = dao.TransactionSaveServiceAll(params, public.LoadTypeHttp)
	if err != nil {
		middleware.ResponseError(ctx, 1003, err)
	}
	middleware.ResponseSuccess(ctx, "添加成功")
}

// @Summary 修改http服务
// @Description 修改http服务
// @Tags 服务管理
// @ID /service/service_update_http
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.ServiceUpdateHttpRequest true "update"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /service/service_update_http [post]
func ServiceUpdateHttp(ctx *gin.Context) {
	// 获取参数并校验
	params := &dto.ServiceUpdateHttpRequest{}
	if err := ctx.BindUri(params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}

	// 检查端口是否被占用
	if params.NeedHttps {
		params.Port = config.Config.Cluster.SSLPort
	} else {
		params.Port = config.Config.Cluster.Port
	}
	serviceDetail := &dao.ServiceDetail{Info: &dao.ServiceInfo{Id: params.ID}}
	err := serviceDetail.FindAll()
	if err != nil {
		middleware.ResponseError(ctx, 1004, err)
		return
	}
	proxy.GetHttpProxy().Update(serviceDetail, &params.ServiceAddHttpRequest)
	// 根据id更新数据
	if err := dao.TransactionUpdateAll(params.ID, params); err != nil {
		middleware.ResponseError(ctx, 1003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "修改成功")
}

// @Summary 添加tcp服务
// @Description 添加tcp服务
// @Tags 服务管理
// @ID /service/service_add_tcp
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.ServiceAddTcpRequest true "add"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /service/service_add_tcp [post]
func ServiceAddTcp(ctx *gin.Context) {
	// 获取参数并校验
	params := &dto.ServiceAddTcpRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 查询服务是否存在
	serviceName := &dao.ServiceInfo{ServiceName: params.ServiceName}
	err := serviceName.Exist("服务")
	if err != nil {
		middleware.ResponseError(ctx, 1005, err)
		return
	}
	// 查询tcp端口是否被占用
	httpRule := &dao.ServiceTcpRule{Port: params.Port}
	err = httpRule.Exist("端口")
	if err != nil {
		middleware.ResponseError(ctx, 1005, err)
		return
	}
	// 保存
	err = dao.TransactionSaveServiceAll(params, public.LoadTypeTcp)
	if err != nil {
		middleware.ResponseError(ctx, 1003, err)
	}
	middleware.ResponseSuccess(ctx, "添加成功")
}

// @Summary 修改tcp服务
// @Description 修改tcp服务
// @Tags 服务管理
// @ID /service/service_update_tcp
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.ServiceUpdateTcpRequest true "update"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /service/service_update_tcp [post]
func ServiceUpdateTcp(ctx *gin.Context) {
	// ServiceUpdateTcpRequest 保存着需要更新的字段
	// 获取参数并校验
	params := &dto.ServiceUpdateTcpRequest{}
	if err := ctx.BindUri(params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 根据 id 的值进行 update
	if err := dao.TransactionUpdateAll(params.ID, params); err != nil {
		middleware.ResponseError(ctx, 1003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "修改成功")
}

// @Summary 添加hGrpc服务
// @Description 添加hGrpc服务
// @Tags 服务管理
// @ID /service/service_add_grpc
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.ServiceAddGrpcRequest true "add"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /service/service_add_grpc [post]
func ServiceAddGrpc(ctx *gin.Context) {
	// 获取参数 并校验
	params := &dto.ServiceAddGrpcRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 查询服务是否存在
	serviceName := &dao.ServiceInfo{ServiceName: params.ServiceName}
	err := serviceName.Exist("服务")
	if err != nil {
		middleware.ResponseError(ctx, 1005, err)
		return
	}
	// 单独查询 grpc 端口是否被占用
	httpRule := &dao.ServiceGrpcRule{Port: params.Port}
	err = httpRule.Exist("端口")
	if err != nil {
		middleware.ResponseError(ctx, 1005, err)
		return
	}
	// 保存
	err = dao.TransactionSaveServiceAll(params, public.LoadTypeGrpc)
	if err != nil {
		middleware.ResponseError(ctx, 1003, err)
	}
	middleware.ResponseSuccess(ctx, "添加成功")
}

// @Summary 修改grpc服务
// @Description 修改tcp服务
// @Tags 服务管理
// @ID /service/service_update_grpc
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.ServiceUpdateGrpcRequest true "update"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /service/service_update_grpc [post]
func ServiceUpdateGrpc(ctx *gin.Context) {
	// ServiceUpdateGrpcRequest 保存着需要更新的字段
	// 获取参数并校验
	params := &dto.ServiceUpdateGrpcRequest{}
	if err := ctx.BindUri(params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	// 根据 id 的值 进行 update操作
	if err := dao.TransactionUpdateAll(params.ID, params); err != nil {
		middleware.ResponseError(ctx, 1003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "修改成功")
}
