package Router

import (
	"1/Controller"
	"github.com/gin-gonic/gin"
)
func Entrance() {
	router := gin.Default()
	router.POST("/register",Controller.Register)
	router.GET("/login",Controller.Login)
	router.POST("/Message",Controller.Auth(),Controller.SendMessage)
	router.GET("/Message",Controller.ViewMessage)
	router.PUT("/Message",Controller.Auth(),Controller.Like)
	router.DELETE("/login",Controller.Auth(),Controller.Cancel)
	router.DELETE("/Message",Controller.Auth(),Controller.DeleteMessage)
	router.Run(":8080")
}