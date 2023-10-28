package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

const (
	ERROR   = false
	SUCCESS = true
)

func Result(code int, success bool, msg string, data any, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		success,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(200, SUCCESS, "操作成功", nil, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(200, SUCCESS, message, nil, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(200, SUCCESS, "查询成功", data, c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(200, SUCCESS, message, data, c)
}

func Fail(c *gin.Context) {
	Result(200, ERROR, "操作失败", nil, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(200, ERROR, message, nil, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(200, ERROR, message, data, c)
}
