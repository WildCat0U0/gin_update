package global

// CustomError 自定义错误
type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

// CustomErrors 自定义错误
type CustomErrors struct {
	BusinessError   CustomError // 业务错误
	ValidateError   CustomError // 验证错误
	TokenError      CustomError // Token错误
	TokenEmptyError CustomError //Token为空
}

// Errors 自定义错误
var Errors = CustomErrors{
	BusinessError: CustomError{40000, "业务错误"},
	ValidateError: CustomError{42200, "验证错误"},
	TokenError: CustomError{
		ErrorCode: 40100,
		ErrorMsg:  "授权登录失败",
	},
	TokenEmptyError: CustomError{
		ErrorCode: 40200,
		ErrorMsg:  "当前没有用户登入，请登陆",
	},
}
