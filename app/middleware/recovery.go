package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				ResponseError(ctx, 500, errors.Errorf("服务器内部错误: %s", err))
			}
		}()
		ctx.Next()
	}
}
