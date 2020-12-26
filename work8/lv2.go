package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r:=gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.Writer.Write([]byte("Hello,world"))
	})
	r.Run()
}
