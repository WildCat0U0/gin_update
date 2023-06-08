package app

import (
	"Gin_Start/app/common/request"
	"Gin_Start/app/common/response"
	"Gin_Start/app/services"
	"github.com/gin-gonic/gin"
)

// Register :注册结构体 校验入参，调用 UserService 注册逻辑，返回结果
//func Register(c *gin.Context) {
//	// 1.查询是否存在
//	// 2.如果不存在则创建
//	// 3.创建用户
//	var form request.Register                       // 定义一个结构体
//	if err := c.ShouldBindJSON(&form); err != nil { // 绑定数据
//		response.ValidateFail(c, request.GetErrorMsg(form, err))
//		return
//	}
//	if err, user := services.UserService.Register(form); err != nil { // 创建用户
//		response.BusinessFail(c, err.Error()) // 失败响应
//	} else {
//		response.Success(c, user) // 成功响应
//	}
//}

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
