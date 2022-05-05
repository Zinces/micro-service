package model

import "github.com/Zinces/micro-service/common/pkg/db"

type User struct {
	db.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null;unique;default:''" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255) not null;unique;default:''" valid:"email"`
	RealName string `gorm:"column:real_name;type:varchar(255);not null;default:''" valid:"realName"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null;default:''" valid:"avatar"`
	Status   int    `gorm:"column:status;type:tinyint(1);not null;default:0" `
	Password string `gorm:"column:password;type:varchar(255) not null;;default:''" valid:"password"`
}


