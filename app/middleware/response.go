package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/config"
)

type ResponseErr struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Stack   string      `json:"stack"`
	Code    int         `json:"code"`
}

func ResponseError(ctx *gin.Context, code int, err error) {
	var resp *ResponseErr
	if config.Mode {
		resp = &ResponseErr{Code: code, Message: err.Error(), Stack: string(debug.Stack())}
	} else {
		resp = &ResponseErr{Code: code, Message: err.Error()}
	}
	log.Println(resp)
	ctx.JSON(http.StatusOK, resp)
}

func ResponseSuccess(ctx *gin.Context, data any) {
	resp := &ResponseErr{Data: data, Code: 200, Message: "success"}
	ctx.JSON(http.StatusOK, resp)
}
