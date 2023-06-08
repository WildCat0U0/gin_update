package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateMobile 验证手机号
// validator.FieldLevel 为验证器的上下文，可以获取到当前验证的字段值
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()                                                                                       //获取当前验证的字段值
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile) //正则验证手机号
	//^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$ 正则表达式
	if ok {
		return ok
	}
	return false
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9]{6,16}$`, password) //正则验证密码
	if ok {
		return ok
	}
	return false
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	ok, _ := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, email)
	if ok {
		return ok
	}
	return false
}
