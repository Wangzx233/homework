package control

import (
	"demo/model"
	_struct "demo/struct"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context)  {
	var user _struct.User
	err:=context.ShouldBind(&user)
	if err!=nil{
		context.JSON(800,"bind_err")
		return
	}
	res,rs:=model.Register(user.UserName,user.Password)
	if res {
		context.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": rs,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": rs,
		})
	}
}

func Login(context *gin.Context) {
	session := sessions.Default(context)
	var user _struct.User
	err:=context.ShouldBind(&user)
	if err!=nil{
		context.JSON(900,err)
		return
	}
	res:=model.Login(user.UserName,user.Password)
	if res{
		context.SetCookie("uid",user.UserName, 9999, "/", "127.0.0.1", false, false)
		session.Set("uid",user.UserName)

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

func IsLogin(c *gin.Context)  {
	session := sessions.Default(c)
	if session.Get("uid")==""{
		_, err := c.Request.Cookie("uid");if err!=nil{
			//c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/")
			c.JSON(300,gin.H{
				"message":"请先登录",
			})
		}
	}else {
		c.JSON(300,gin.H{
			"message":"连接成功",
		})
	}


}