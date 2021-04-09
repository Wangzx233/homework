package router

import (
	"demo/control"
	"demo/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router()  {
	r := gin.Default()
	r.Use(middleware.Cors())
	store := cookie.NewStore([]byte("secret11111"))
	r.Use(sessions.Sessions("uid", store))


	r.GET("/ws",middleware.Auth(),control.Socket)
	r.GET("/user/login",control.Login)
	r.POST("/user/register",control.Register)
	r.GET("user/isLogin",control.IsLogin)
	r.Run("127.0.0.1:8080")
}