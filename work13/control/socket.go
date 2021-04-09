package control

import (
	"demo/socket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Socket(c *gin.Context)  {
	//session := sessions.Default(c)
	//uid := session.Get("uid")


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


	client := socket.Client{
		Conn:    conn,
		MsgChan: make(chan socket.Message),
		User:    "",
	}

	go client.Reader()
	go client.Write()
	//defer func() {
	//	socket.Mng.Logout(client.User)
	//}()
}