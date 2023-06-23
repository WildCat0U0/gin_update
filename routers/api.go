package routers

import (
	"Gin_Start/app/controllers/app"
	"Gin_Start/app/middleware"
	"Gin_Start/app/services"
	"Gin_Start/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetApiGroupRouters
// @Summary Gin_Start API
// @title Gin_Start API
// @description Gin_Start API
// @version
// @host localhost:8080
// @basePath /api
func SetApiGroupRouters(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/test_shutdown", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(200, "test_shutdown_success")
	})
	// 用户注册
	//router.POST("/user/register", func(c *gin.Context) {
	//	var form request.Register
	//	if err := c.ShouldBindJSON(&form); err != nil { // 绑定数据
	//		c.JSON(http.StatusOK, gin.H{ // 返回错误信息
	//			"error": request.GetErrorMsg(form, err), // 获取错误信息
	//		})
	//		return // 终止程序
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "register_success",
	//	})
	//})
	//router.POST("/user/register", func(c *gin.Context) {
	//	var form request.Register
	//	if err := c.ShouldBindJSON(&form); err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"error": request.GetErrorMsg(form, err),
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "success",
	//	})
	//})
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)
	router.POST("/auth/updatePassword", app.ChangePasswordHandler)
	router.POST("/cookie/test", func(c *gin.Context) {
		cookie, err := c.Cookie("mycookie")
		if err != nil {
			global.App.Log.Info("cookie获取错误，请联系管理员")
		}
		if cookie == "1234567" {
			c.JSON(200, gin.H{
				"message": "hello",
			})
		} else {
			cookie1 := &http.Cookie{
				Name:     "mycookie",
				Value:    "1234567",
				Expires:  time.Now(),
				HttpOnly: true,
			}
			c.SetCookie(cookie1.Name, cookie1.Value, cookie1.MaxAge, cookie1.Path, cookie1.Domain, cookie1.Secure, cookie1.HttpOnly)
			c.Redirect(301, "/ping")
		}
	})

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/auth/is_login", app.IsLogin)
	}

}
