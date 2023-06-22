# Gin_WoW

from wildcat's test

<div style="color: deepskyblue">
文件目录

|   name    | function | <div style= "color : deepskyblue">备注<div> |
|:---------:|:--------:|:-----------------------------------------:|
|  **app**  |   公共模块   |                   请求，相应                   |
| bootstrap | 项目启动初始化  |                服务器，数据库初始化                 |
|  config   |    配置    |                    没有                     |
|  global   |   全局变量   |                    没有                     |
|  routers  |   路由定义   |                    没有                     |
|  static   |   静态资源   |                    没有                     |
|  storage  | 系统日志，文件  |                    没有                     |
|   utils   |   工具函数   |                    没有                     |
|  main.go  |   主函数    |                    没有                     |

<div style="color: deeppink">
 wildcat's test
</div>

</div>

<div style="color: #14ff33">
 
## 5-26 update
```util/validator.go```
```go
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
```
</div>

然后修改文件```bootstrap/validator.go```
```go
_ = v.RegisterValidation("email", utils.ValidateEmail)
_ = v.RegisterValidation("password", utils.ValidatePassword)
```
需要修改文件 ```app/common.request/user.go```
```go
"email.required":    "邮箱不能为空",
"email.email":       "邮箱格式不正确",
"password.password": "密码必须包含数字和字母",

Email    string `json:"email" binding:"required,email" form:"email"` 
```

完成修改

<div style="color: #42b983">

## 6-7 update
`config.yaml`

<div style="display: flex;">
  <div style="flex: 1; text-align: left;">

    database:
    host:     127.0.0.1
  </div>
  <div style="flex: 1; text-align: left;">

    database:
    host:     172.17.0.3
  </div>
</div>


</div>


## 2023年6月23日Update
> version 1.0.3 ： 更新主要内容 增加黑名单
> windows端 使用 `docker`的mysql和redis一样可以运行成功，十分的nice