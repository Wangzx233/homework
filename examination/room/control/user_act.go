package control

import (
	"examination/room/stuct"
	"github.com/MashiroC/begonia"
	"github.com/MashiroC/begonia/app/option"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var c = begonia.NewClient(option.Addr(":12306"))
func Register(context *gin.Context)  {
	var user _struct.User
	err:=context.ShouldBind(&user)
	if err!=nil{
		context.JSON(800,"bind_err")
		return
	}


	// 获取一个服务
	s, err := c.Service("user_center")
	if err != nil {
		panic(err)
	}
	// 获取一个远程函数的同步调用
	Register, err := s.FuncSync("Register")
	if err != nil {
		panic(err)
	}

	res,rs:=Register(user.UserName,user.Password)

	if rs!=nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": res,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": res,
		})
	}
}

func Login(context *gin.Context) {
	var user _struct.User
	err:=context.ShouldBind(&user)
	if err!=nil{
		context.JSON(900,err)
		return
	}


	// 获取一个服务
	s, err := c.Service("user_center")
	if err != nil {
		panic(err)
	}
	// 获取一个远程函数的同步调用
	Login, err := s.FuncSync("Login")
	if err != nil {
		panic(err)
	}
	res,err:=Login(user.UserName,user.Password)
	if err != nil {
		log.Println(err)
	}
	if res.(bool){
		context.SetCookie("uid",user.UserName, 9999, "/", "47.106.170.23", false, false)

		context.JSON(200,gin.H{
			"code":200,
			"message":	"你好"+user.UserName,
		})

	}else {
		context.JSON(200,gin.H{
			"code":		20000,
			"message":	"用户名或密码错误",
		})
	}
}
func Cancel(context *gin.Context)  {
	session := sessions.Default(context)
	context.SetCookie("uid","",-1, "/", "/", false, false)
	session.Delete("uid")
}

