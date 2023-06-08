package request

import (
	"github.com/go-playground/validator/v10"
)

//在此文件中可以定义自己的验证器，比如手机号验证器，邮箱验证器等等，这里以手机号验证器为例

//Gin  自带验证器返回的错误信息格式不太友好，本篇将进行调整，实现自定义错误信息，
//并规范接口返回的数据格式，分别为每种类型的错误定义错误码，前端可以根据对应的错误码实现后续不同的逻辑操作，
//篇末会使用自定义的 Validator 和 Response 实现第一个接口

type Validator interface {
	// GetMessages :Validate 验证器
	GetMessages() validatorMessages
}

type validatorMessages map[string]string //定义一个map类型，key为string类型，value为string类型

//func GetErrorMsg(request interface{}, err error) string {
//	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors { //判断是否是validator.ValidationErrors类型
//		_, isValidator := request.(Validator) //判断是否实现了Validator接口
//
//		for _, v := range err.(validator.ValidationErrors) {
//			//如果request结构体实现了validator接口，就调用GetMessages方法获取自定义的错误信息
//			if isValidator {
//				if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
//					return message
//				}
//			}
//			//如果没有实现validator接口，就使用默认的错误信息
//			return v.Error()
//		}
//	}
//	// 如果不是validator.ValidationErrors类型就返回默认的错误信息
//	return "parameter error"
//}

// request.(Validator)是类型断言，判断request是否实现了Validator接口，如果实现了就返回true，否则返回false
// GetErrorMsg 获取错误信息

func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)

		for _, v := range err.(validator.ValidationErrors) {
			// 若 request 结构体实现 Validator 接口即可实现自定义错误信息
			if isValidator {
				if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			}
			return v.Error()
		}
	}

	return "Parameter error"
}
