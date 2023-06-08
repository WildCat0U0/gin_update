package app

import (
	"Gin_Start/app/common/request"
	"Gin_Start/app/common/response"
	"Gin_Start/app/services"
	"github.com/gin-gonic/gin"
)

func ChangePasswordHandler(c *gin.Context) {
	var form request.ChangePassword
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err := services.UserService.UpdatePassword(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "密码更新成功")
	}
}
