package Controller

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, err := context.Request.Cookie("uid")
		if err == nil {
			context.Next()
		} else {
			context.Abort()
			context.Writer.Write([]byte("请先登录"))
		}
	}
}