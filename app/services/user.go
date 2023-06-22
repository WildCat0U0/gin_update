package services

import (
	"Gin_Start/app/common/request"
	"Gin_Start/app/models"
	"Gin_Start/global"
	"Gin_Start/utils"
	"errors"
	"strconv"
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
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password)), Email: params.Email}
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

func (userService *userService) CheckPassword(mobile string, oldPassword string) error {
	var user models.User
	var err error
	var result = global.App.DB.Model(&models.User{}).Where("mobile = ?", mobile).First(&user)
	if result.RowsAffected == 0 {
		err = errors.New("此电话号码并没有注册过用户")
		return err
	}
	if !utils.BcryptMakeCheck([]byte(oldPassword), user.Password) {
		err = errors.New("旧密码错误，无法更改密码，请确认")
		return err
	}
	return nil
}

func (userService *userService) UpdatePassword(params request.ChangePassword) error {
	//var user models.User
	err := userService.CheckPassword(params.Mobile, params.OldPassword)
	if err != nil {
		return err
	}
	result := global.App.DB.Model(&models.User{}).Where("mobile = ?", params.Mobile).Update("password", utils.BcryptMake([]byte(params.NewPassword)))
	if result.RowsAffected == 0 {
		return errors.New("修改密码失败，请联系管理员")
	}
	return nil
}

func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, _ := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		global.App.Log.Info("用户数据不存在")
	}
	return
}
