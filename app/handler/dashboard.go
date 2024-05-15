package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
	"github.com/stone2401/light-gateway/app/tools/db"
	"github.com/stone2401/light-gateway/app/tools/redis"
)

func RegisterDashBoard(router *gin.RouterGroup) {
	router.GET("/panelGroupData", PanelGroupData)
	router.GET("/flowStat", FlowStat)
	router.GET("/service_stat", ServiceStatAll)
	router.GET("/:id", Dashboard)
}

// @Summary 指标统计
// @Description 指标统计
// @Tags 大盘
// @ID /dashboard/panelGroupData
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.PanelGroupDataResponse} "success"
// @Router /dashboard/panelGroupData [get]
func PanelGroupData(ctx *gin.Context) {
	service := &dao.ServiceInfo{}
	totleService, err := service.GetTatle()
	if err != nil {
		middleware.ResponseError(ctx, 1002, err)
	}
	appinfo := &dao.App{}
	totleApp, err := appinfo.GetTatle()
	if err != nil {
		middleware.ResponseError(ctx, 1002, err)
	}
	response := &dto.PanelGroupDataResponse{ServiceNum: totleService,
		AppNum: totleApp,
	}
	middleware.ResponseSuccess(ctx, response)
}

// @Summary 指标统计
// @Description 指标统计
// @Tags 大盘
// @ID /dashboard/flowStat
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceStatResponse} "success"
// @Router /dashboard/flowStat [get]
func FlowStat(ctx *gin.Context) {
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

// @Summary 指标统计
// @Description 指标统计
// @Tags 大盘
// @ID /dashboard/service_stat
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceStatAllResponse} "success"
// @Router /dashboard/service_stat [get]
func ServiceStatAll(ctx *gin.Context) {
	serviceInfo := &dao.ServiceInfo{}
	ssair, err := serviceInfo.GroupByLoadType()
	if err != nil {
		middleware.ResponseError(ctx, 1002, err)
		return
	}
	// 获取 数据
	serviceStat := &dto.ServiceStatAllResponse{
		Legend: []string{}, Data: ssair}
	// 发送
	middleware.ResponseSuccess(ctx, serviceStat)
}

// @Summary 指标统计
// @Description 指标统计
// @Tags 大盘
// @ID /dashboard/
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.ServiceStatAllResponse} "success"
// @Router /api/v1/dashboard/:id [get]
func Dashboard(ctx *gin.Context) {
	req := &dto.ServiceDashboardRequest{}
	resp := &dto.ServiceDashboardResponse{}
	if err := ctx.ShouldBindUri(req); err != nil {
		middleware.ResponseError(ctx, 1002, err)
		return
	}
	if req.Id == 0 {
		info := &dao.ServiceInfo{}
		serNum, err := db.GetDBDriver().Table(info).Count(info)
		if err != nil {
			middleware.ResponseError(ctx, 1002, err)
			return
		}
		openNUm, err := db.GetDBDriver().Table(info).Where("status = ?", public.StatusUp).Count(info)
		if err != nil {
			middleware.ResponseError(ctx, 1002, err)
			return
		}
		qps := &dao.Qps{}
		_, err = db.GetDBDriver().Table(&dao.Qps{}).
			Where("service_id = ? and time = ?", 0, time.Now().Format("20060102")).
			Get(qps)
		if err != nil {
			middleware.ResponseError(ctx, 1002, err)
			return
		}
		yesterQps := &dao.Qps{}
		_, err = db.GetDBDriver().Table(&dao.Qps{}).
			Where("service_id = ? and time = ?", 0, time.Now().Add(time.Hour*-24).Format("20060102")).
			Get(yesterQps)
		if err != nil {
			middleware.ResponseError(ctx, 1002, err)
			return
		}
		redisCmd := redis.GetRedisConn().Get(ctx, "all#"+time.Now().Format("2006010215"))
		if redisCmd.Err() != nil {
			middleware.ResponseError(ctx, 1002, err)
			return
		}
		current, _ := redisCmd.Int64()
		qps.SetHour(time.Now().Format("15"), current)
		fmt.Println(qps, yesterQps)
		resp.TodayRequestNum = qps.GetQps()
		resp.ServiceNum = serNum
		resp.OpenService = openNUm
		resp.CurrentQps = current
		resp.Datas = qps.GetDay()
		resp.YesterDates = yesterQps.GetDay()
	} else {
		nowQps := &dao.Qps{}
		_, err := db.GetDBDriver().Table(nowQps).
			Where("service_id = ? and time = ?", req.Id, time.Now().Format("20060102")).Get(nowQps)
		if err != nil {
			middleware.ResponseError(ctx, 1002, err)
			return
		}
		yesterQps := &dao.Qps{}
		_, err = db.GetDBDriver().Table(yesterQps).
			Where("service_id = ? and time = ?", req.Id, time.Now().Add(time.Hour*-24).Format("20060102")).Get(yesterQps)
		if err != nil {
			middleware.ResponseError(ctx, 1002, err)
			return
		}
		resp.Datas = nowQps.GetDay()
		resp.YesterDates = yesterQps.GetDay()
	}
	resp.Times = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	middleware.ResponseSuccess(ctx, resp)
}
