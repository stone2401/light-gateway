package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/stone2401/light-gateway/app/middleware"
	"github.com/stone2401/light-gateway/app/model/dao"
	"github.com/stone2401/light-gateway/app/model/dto"
	"github.com/stone2401/light-gateway/app/public"
)

func RegisterApp(router *gin.RouterGroup) {
	router.GET("/app_list", AppList)
	router.GET("/app_detail", AppDetail)
	router.GET("/app_stat", AppStat)
	router.GET("/app_delete", AppDelete)

	router.POST("app_add", AppAdd)
	router.POST("app_update", AppUpdate)
}

// @Summary 用户列表
// @Description 用户列表
// @Tags 用户管理
// @ID /app/app_list
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param info query string false "查询词汇"
// @Param page_no query int true "页数"
// @Param page_size query int true "条数"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dto.AppListResponse} "success"
// @Router /app/app_list [get]
func AppList(ctx *gin.Context) {
	// 获取数据
	params := &dto.ServiceListRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
	}
	// 查询信息
	app := &dao.App{}
	list, totle, err := app.PageList(params)
	if err != nil {
		middleware.ResponseError(ctx, 1002, err)
	}
	// 信息组合
	response := &dto.AppListResponse{Total: int64(totle)}
	lists := []*dto.AppListItem{}
	for _, value := range list {
		item := &dto.AppListItem{RealQpd: 0, RealQps: 0}
		copier.Copy(item, value)
		lists = append(lists, item)
	}
	response.List = lists
	middleware.ResponseSuccess(ctx, response)
}

// @Summary 用户详情
// @Description 用户详情
// @Tags 用户管理
// @ID /app/app_detail
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param id query string true "关键词"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=dao.App} "success"
// @Router /app/app_detail [get]
func AppDetail(ctx *gin.Context) {
	params := &dto.AppDetailRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
	}
	app := &dao.App{Id: params.ID}
	if err := app.Find(); err != nil {
		middleware.ResponseError(ctx, 1002, err)
	}
	middleware.ResponseSuccess(ctx, app)
}

// @Summary 用户统计
// @Description 用户统计
// @Tags 用户管理
// @ID /app/app_stat
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param id query string true "关键词"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /app/app_stat [get]
func AppStat(ctx *gin.Context) {
	params := &dto.AppDetailRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
	}
	middleware.ResponseSuccess(ctx, "用户统计!")
}

// @Summary 用户删除
// @Description 用户删除
// @Tags 用户管理
// @ID /app/app_delete
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param id query string true "关键词"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /app/app_delete [get]
func AppDelete(ctx *gin.Context) {
	params := &dto.AppDetailRequest{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
	}
	app := &dao.App{Id: params.ID}
	if err := app.Delete(); err != nil {
		middleware.ResponseError(ctx, 1002, err)
	}
	middleware.ResponseSuccess(ctx, "删除成功，ok!")
}

// @Summary 修改用户
// @Description 修改用户
// @Tags 用户管理
// @ID /app/app_update
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.AppListItem true "add"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /app/app_update [post]
func AppUpdate(ctx *gin.Context) {
	params := &dto.AppListItem{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	if params.Id == 0 {
		middleware.ResponseError(ctx, 1001, errors.New("id 不可为空"))
		return
	}
	app := &dao.App{}
	copier.Copy(app, params)
	if app.Secret == "" {
		app.Secret = uuid.NewString()
	}
	if err := app.Update(); err != nil {
		middleware.ResponseError(ctx, 1002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "修改成功")
}

// @Summary 添加用户
// @Description 添加用户
// @Tags 用户管理
// @ID /app/app_add
// @Param Authorization	header string true "token"
// @Security ApiKeyAuth
// @Param body body dto.AppListItem true "add"
// @Accept application/json
// @Produce json
// @Success 200 {object} middleware.ResponseErr{data=string} "success"
// @Router /app/app_add [post]
func AppAdd(ctx *gin.Context) {
	params := &dto.AppListItem{}
	if err := public.Authenticator(ctx, params); err != nil {
		middleware.ResponseError(ctx, 1001, err)
		return
	}
	{
		temp := &dao.App{AppId: params.AppId}
		err2 := temp.Exist("AppId")
		if err2 != nil {
			middleware.ResponseError(ctx, 1001, err2)
			return
		}
	}
	{
		temp := &dao.App{Name: params.Name}
		err2 := temp.Exist("Name")
		if err2 != nil {
			middleware.ResponseError(ctx, 1001, err2)
			return
		}
	}
	app := &dao.App{}
	copier.Copy(app, params)
	if app.Secret == "" {
		app.Secret = uuid.NewString()
	}
	if err := app.Save(); err != nil {
		middleware.ResponseError(ctx, 1002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "添加成功")
}
