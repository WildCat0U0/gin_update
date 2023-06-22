package response

import (
	"Gin_Start/global"
	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrorCode int         `json:"error_code"` // 自定义错误码
	Message   string      `json:"message"`    // 信息
	Data      interface{} `json:"data"`       // 数据
}

// Success ---成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		ErrorCode: 0,
		Message:   "ok",
		Data:      data,
	})
}
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// Fail ---失败响应
func Fail(c *gin.Context, errorCode int, message string) {
	c.JSON(200, Response{
		ErrorCode: errorCode, // 自定义错误码
		Message:   message,   // 信息
		Data:      nil,       // 数据
	})
}

// FailWithDetailed ---失败响应
func FailWithDetailed(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail ---失败响应
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail ---失败响应
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.TokenError)
}
