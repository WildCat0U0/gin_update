package app

import (
	"Gin_Start/app/common/request"
	"Gin_Start/app/common/response"
	"Gin_Start/app/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		cookie1 := &http.Cookie{Name: "tokenoflogin", Value: tokenData.AccessToken, Expires: time.Time{}.Add(3600 * time.Second), HttpOnly: true}
		c.SetCookie(cookie1.Name, cookie1.Value, cookie1.MaxAge, cookie1.Path, cookie1.Domain, cookie1.Secure, cookie1.HttpOnly)
		response.Success(c, tokenData)
	}
}
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	//response.Success(c, "成功退出")
	cookie1 := &http.Cookie{Name: "tokenoflogin", Value: "", Expires: time.Time{}.Add(3600 * time.Second), HttpOnly: true}
	c.SetCookie(cookie1.Name, cookie1.Value, cookie1.MaxAge, cookie1.Path, cookie1.Domain, cookie1.Secure, cookie1.HttpOnly)
	c.Redirect(301, "/")
}

func IsLogin(c *gin.Context) {

}
