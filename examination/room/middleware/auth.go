package middleware

import (
	"github.com/gin-gonic/gin"
)
//需要登录
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {

		_, err := context.Request.Cookie("uid")
		if err == nil {
			context.Next()
		} else {
			context.Abort()
			//context.JSON(500,gin.H{
			//	"message":"请先登录",
			//})
		}
	}
}

