package control

import (
	"examination/room/socket"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Socket(c *gin.Context)  {
	//session := sessions.Default(c)
	//uid := session.Get("uid")


	//方便测试暂时注释
	//cookie, err := c.Request.Cookie("uid")
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(200, gin.H{
	//		"status": 10021,
	//		"info": "failed",
	//	})
	//	return
	//}

	roomID := c.Query("room_id")
	uid := c.Query("uid")

	upgrader := websocket.Upgrader{
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"status": 10022,
			"info": "failed",
		})
		return
	}
	fmt.Println(uid,roomID)
	conn.SetCloseHandler(func(code int, msg string) error {
		socket.Mng.Logout(uid,roomID)
		fmt.Println("关闭了")
		return nil

	})


	client := socket.Client{
		Conn:    conn,
		MsgChan: make(chan socket.Message),
		User:    "",
	}

	go client.Reader()
	go client.Write()

}