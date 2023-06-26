package models

import (
	"gorm.io/gorm"
	"time"
)

//定义公用模型字段进行数据库迁移
//
//type ID struct {
//	ID uint `gorm:"primarykey" json:"id"`
//}
//
//// Timestamps 定义公共的时间戳字段
//// gorm:"column:create_at" 为该字段在数据库中的名称
//// json:"create_at" 为该字段在json中的名称
//type Timestamps struct {
//	CreateAt time.Time `json:"create_at"` // 创建时间 gorm:"column:create_at" 为该字段在数据库中的名称
//	UpdateAt time.Time `json:"update_at"` // 更新时间 gorm:"column:update_at" 为该字段在数据库中的名称
//}
//
//// SoftDeletes 定义软删除字段
//// gorm :"index" 为该字段创建索引
//// json:"delete_at" 为该字段在json中的名称
//type SoftDeletes struct {
//	DeleteAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
//}

// ID 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建、更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
