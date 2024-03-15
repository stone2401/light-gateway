package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/stone2401/light-gateway/config"
)

type ResponseErr struct {
	ErrorCode int         `json:"errno"`
	ErrorMsg  string      `json:"errmsg"`
	Data      interface{} `json:"data"`
	Stack     string      `json:"stack"`
}

func ResponseError(ctx *gin.Context, code int, err error) {
	var resp *ResponseErr
	if config.Mode {
		resp = &ResponseErr{ErrorCode: code, ErrorMsg: err.Error(), Stack: string(debug.Stack())}
	} else {
		resp = &ResponseErr{ErrorCode: code, ErrorMsg: err.Error()}
	}
	log.Println(resp)
	ctx.JSON(http.StatusOK, resp)
}

func ResponseSuccess(ctx *gin.Context, data any) {
	resp := &ResponseErr{ErrorCode: 0, Data: data}
	ctx.JSON(http.StatusOK, resp)
}
