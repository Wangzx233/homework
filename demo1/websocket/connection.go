package socket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type connection struct {
	ws      *websocket.Conn
	message chan []byte
	data    *Data
}


var userlist []string



//ws升级器

//ws回调函数
func Go(c *gin.Context) {
	upgrade := websocket.Upgrader{
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"message": "failed"})
		return
	}
	connection := &connection{
		ws:      conn,
		message: make(chan []byte),
		data:    &Data{},
	}
	go connection.writer()
	go connection.reader()
	//H.register <- connection


	defer func() {
		connection.data.Type = "logout"
		userlist = remove(userlist, connection.data.User)
		connection.data.UserList = userlist
		connection.data.Content = connection.data.User
		marshal, err := json.Marshal(connection.data)
		if err != nil {
			log.Println(err)
			return
		}
		H.broad <- marshal
		H.unregister <- connection
	}()
}
