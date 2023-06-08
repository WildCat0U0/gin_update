package request

// Register :注册结构体
// @Summary 注册结构体
// @Description 注册结构体
// @Tags 用户
// @Accept json
// form 表单提交
// blinding 表单验证
type Register struct {
	Name     string `json:"name" binding:"required,min=2,max=20" form:"name"`
	Password string `json:"password" binding:"required,min=6,max=20" form:"password"`
	Mobile   string `json:"mobile" binding:"required,mobile,min=11,max=11" form:"mobile"` // binding:"required" 表示必填字段
	Email    string `json:"email" binding:"required,email" form:"email"`                  // binding:"required" 表示必填字段
}

func (register Register) GetMessages() validatorMessages {
	return validatorMessages{
		"name.required":     "用户名不能为空",
		"name.min":          "用户名长度不能小于2",
		"name.max":          "用户名长度不能大于20",
		"password.required": "密码不能为空",
		"password.min":      "密码长度不能小于6",
		"password.max":      "密码长度不能大于20",
		"mobile.required":   "手机号不能为空",
		"mobile.min":        "手机号长度不能小于11",
		"mobile.max":        "手机号长度不能大于11",
		"mobile.mobile":     "手机号格式不正确",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式不正确",
		"password.password": "密码必须包含数字和字母",
	}
}

type Login struct {
	Mobile   string `json:"mobile" binding:"required,mobile,min=11,max=11" form:"mobile"` // binding:"required" 表示必填字段
	Password string `json:"password" binding:"required,min=6,max=20" form:"password"`
}

func (Login Login) GetMessages() validatorMessages {
	return validatorMessages{
		"mobile.required":   "手机号不能为空",
		"mobile.mobile":     "手机号格式不正确",
		"password.required": "密码不能为空",
	}
}
