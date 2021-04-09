package _struct

import "github.com/jinzhu/gorm"


//用户结构
type User struct {
	//ID			string
	gorm.Model
	UserName   string   `form:"username" json:"username"` //binding:"required,min=5,max=30"`
	Password   string   `form:"password" json:"password"` //binding:"required,min=8,max=40"`
}
