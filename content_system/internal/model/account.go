package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserID   string `gorm:"column:user_id;"`
	Password string `gorm:"column:password;"`
	Nickname string `gorm:"column:nickname;"`
}

// 指定表名
func (Account) TableName() string {
	table := "cms_content.user"
	return table
}
