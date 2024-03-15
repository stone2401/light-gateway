package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
)

func RegisterDashBoard(router *gin.RouterGroup) {
	router.GET("/panelGroupData", PanelGroupData)
	router.GET("/flowStat", FlowStat)
	router.GET("/service_stat", ServiceStatAll)
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
