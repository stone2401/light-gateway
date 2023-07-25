package handler

import "github.com/gin-gonic/gin"

// @description 检测接口
// @Tags api-v1
// @success 200 {string} string "pong"
// @router /ping [GET]
func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200, "pong")
	}
}
