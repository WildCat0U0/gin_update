package services

import (
	"Gin_Start/app/common/request"
	"Gin_Start/app/models"
	"Gin_Start/global"
	"Gin_Start/utils"
	"errors"
)

type userService struct {
}

var UserService = userService{}

//// Register :注册结构体
//func (userService *userService) Register(params request.Register) (err error, user models.User) {
//	// 1.查询是否存在
//	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{}) // 查询是否存在
//	// 2.如果不存在则创建
//	if result.RowsAffected != 0 {
//		err = errors.New("手机号码已经存在")
//	}
//	user = models.User{
//		Name:     params.Name,                               // 用户名
//		Password: utils.BcryptMake([]byte(params.Password)), // 加密处理
//		Mobile:   params.Mobile,                             // 手机号
//	}
//	// 3.创建用户
//	global.App.DB.Create(&user)
//	return
//}

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error //在这里创建了一个
	return
}

func (userService *userService) Login(params request.Login) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).First(&user)
	if result.RowsAffected == 0 {
		err = errors.New("用户不存在")
		return
	}
	if !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("密码错误")
		return
	}
	return
}
