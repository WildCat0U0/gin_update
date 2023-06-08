package models

import "strconv"

//定义 User 模型

type User struct {
	ID
	Name     string `gorm:"not null;comment:用户名称" json:"name"`
	Password string `gorm:"not null;comment:用户密码" json:"password"`
	Mobile   string `gorm:"not null;index;comment:用户手机号" json:"mobile"`
	Email    string `gorm:"not null;index;comment:用户邮箱" json:"email"` // index 表示 Email 字段在数据库中创建索引
	Timestamps
	SoftDeletes
}

// GetUid 获取用户id 通过实现Jwtuser接口来调用CreateToken方法
func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
