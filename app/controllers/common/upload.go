package common

import (
	"Gin_Start/app/common/request"
	"Gin_Start/app/common/response"
	"Gin_Start/app/services"
	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
