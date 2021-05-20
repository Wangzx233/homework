package main

import "github.com/jinzhu/gorm"

//用户登录结构
type UserLogin struct {
	UserName string `form:"username" json:"username"` //binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password"` //binding:"required,min=8,max=40"`
}

//用户注册结构
type UserRegister struct {
	UserName string `form:"username" json:"username"` //binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password"` //binding:"required,min=8,max=40"`
}

//用户更改密码结构
type UserChangePassword struct {
	UserName    string `form:"username" json:"username"` //binding:"required,min=5,max=30"`
	Password    string `form:"password" json:"password"` //binding:"required,min=8,max=40"`
	NewPassword string `form:"new_password" json:"new_password"`
} //用户结构
type User struct {
	//ID			string
	gorm.Model
	UserName   string   `form:"username" json:"username"` //binding:"required,min=5,max=30"`
	Password   string   `form:"password" json:"password"` //binding:"required,min=8,max=40"`
	Token      string   `form:"token" json:"token"`
	UserInfo   UserInfo `gorm:"ForeignKey:UserInfoID"`
	UserInfoID int
}

//关注数、粉丝数、获赞数、播放数
type UserInfo struct {
	gorm.Model
	Uid        string `form:"uid" json:"uid"`
	Stars      int    `form:"stars" json:"stars"`
	Fans       int    `form:"fans" json:"fans"`
	Likes      int    `form:"likes" json:"likes"`
	Plays      int    `form:"plays" json:"plays"`
	PictureURL string `form:"picture_url" json:"picture_url"`
}
