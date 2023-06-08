package bootstrap

//import (
//	"github.com/gin-gonic/gin/binding"
//	"github.com/go-playground/validator/v10"
//	"reflect"
//	"strings"
//)

//定制gin中的validator验证器
//
//func InitializeValidator() {
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		////注册自定义验证器
//		//_ = v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
//		//	mobile := fl.Field().String()
//		//	if mobile == "" {
//		//		return false
//		//	}
//		//	//正则表达式验证手机号
//		//	mobileReg := regexp.MustCompile(`^1[3-9]\d{9}$`)
//		//	return mobileReg.MatchString(mobile)
//		//})
//		_ = v.RegisterValidation("mobile", utils.ValidatorMobile) //注册自定义验证器
//		//注册自定义 json tag 函数 用于自定义错误信息
//		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
//			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0] //获取json tag
//			if name == "-" {
//				return "" //如果tag为-，则返回空
//			}
//			fmt.Println(name)
//			return name
//		})
//
//	}
//}

import (
	"Gin_Start/utils"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)
		_ = v.RegisterValidation("email", utils.ValidateEmail)
		_ = v.RegisterValidation("password", utils.ValidatePassword)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			fmt.Println(name)
			return name
		})
	}
}
