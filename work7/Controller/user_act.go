package Controller

import (
	"1/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context)  {
	uid:=context.PostForm("uid")
	password:=context.PostForm("password")
	name:=context.PostForm("name")
	res:=Model.Register(uid,name,password)
	if res {
		context.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "创建用户成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "id已被使用",
		})
	}
}
func Login(context *gin.Context) {
	uid:=context.PostForm("uid")
	password:=context.PostForm("password")
	res:=Model.Login(uid,password)
	if res{
		context.SetCookie("uid",uid, 9999, "/", "localhost", false, true)

		context.JSON(200,gin.H{
			"code":		10000,
			"message":	"你好"+context.PostForm("uid"),
		})
	}else {
		context.JSON(200,gin.H{
			"code":		20000,
			"message":	"用户名或密码错误",
		})
	}
}
func Cancel(context *gin.Context)  {
	context.SetCookie("uid","",-1, "/", "localhost", false, true)
}