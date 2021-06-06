package router

import (
	"examination/room/control"
	"examination/room/middleware"
	"github.com/gin-gonic/gin"
)

func Router()  {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.GET("/ws",control.Socket)
	r.POST("/cheek_create",control.CreateCheek)
	r.POST("/cheek_inter",control.InterCheek)
	r.GET("/user/login",control.Login)
	r.POST("/user/register",control.Register)
	r.Run()
}