package router

import (
	"github.com/gin-gonic/gin"
	socket "work12/websocket"
)

func Router()  {
	r :=gin.Default()

	r.GET("/ws",socket.Go)
	r.Run("127.0.0.1:8080").Error()
}