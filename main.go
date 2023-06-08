package main

import (
	"Gin_Start/bootstrap"
	"Gin_Start/global"
)

func main() {
	bootstrap.InitializeConfig()

	//初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("日志初始化成功")

	//初始化数据库
	var err error
	global.App.DB, err = bootstrap.InitializeDB()
	if err != nil {
		global.App.Log.Info("数据库初始化失败")
	} else {
		global.App.Log.Info("数据库初始化成功")
	}

	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	bootstrap.InitializeValidator() //初始化验证器
	bootstrap.RunServer()
}
