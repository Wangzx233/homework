package _struct

import "gorm.io/gorm"

//用户注册结构
type User struct {
	gorm.Model
	UserName string `form:"username" json:"username" binding:"required,min=1,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=1,max=40"`
}