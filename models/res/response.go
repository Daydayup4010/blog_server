package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 1
	ERROR   = -1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg})
}

func Ok(data map[string]any, msg string, ctx *gin.Context) {
	Result(SUCCESS, data, msg, ctx)
}

func OkWithData(data any, ctx *gin.Context) {
	Result(SUCCESS, data, "success", ctx)
}

func OkWithMsg(msg string, ctx *gin.Context) {
	Result(SUCCESS, map[string]any{}, "success", ctx)
}

func Fail(data, msg string, ctx *gin.Context) {
	Result(ERROR, data, msg, ctx)
}

func FailWithMsg(msg, ctx *gin.Context) {
	Result(ERROR, map[string]any{}, "error", ctx)
}

func FailWithCode(code int, ctx *gin.Context) {
	msg, ok := ErrorMap[ErrorCode(code)]
	if ok {
		Result(code, map[string]any{}, msg, ctx)
		return
	}
	Result(ERROR, map[string]any{}, "unknown error", ctx)
}
