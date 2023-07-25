package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/app/public"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ResponseError(ctx, 2001, errors.New("未传递token"))
			ctx.Abort()
			return
		}
		c, err := public.ParseToken(token)
		ctx.Set("token", c)
		if err != nil {
			ResponseError(ctx, 2002, err)
			ctx.Abort()
			return
		}
	}
}
