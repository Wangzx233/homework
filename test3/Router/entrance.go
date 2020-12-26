package Router

import (
	"test3/Controller"
	"github.com/gin-gonic/gin"
)
func Entrance() {
	router := gin.Default()
	router.POST("/register",Controller.Register)
	router.GET("/login",Controller.Login)
	router.PUT("/Recharge",Controller.Auth(),Controller.Recharge)
	router.PUT("/transfer",Controller.Auth(),Controller.Transfer)
	router.GET("/queryremarks",Controller.QueryRemarks)
	router.GET("/queryUid",Controller.Auth(),Controller.QueryUid)

	router.Run(":8080")
}