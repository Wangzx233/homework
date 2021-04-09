package socket

import (
	"encoding/json"
	"log"
)

type Hub struct {
	connections map[string]map[*connection]bool
	broad       chan []byte
	register    chan *connection
	unregister  chan *connection
}

var H = Hub{
	connections: make(map[string]map[*connection]bool),
	broad:       make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
}

func (h *Hub) RUN() {
	for {
		select {
		case c := <-h.register:
			h.connections[c.data.RoomID][c] = true
			//组装数据
			c.data.UserList = userlist
			//marshal, err := json.Marshal(c.data)
			//if err != nil {
			//	log.Println(err)
			//	return
			//}
			c.message <- []byte("上线了")
		case c := <-h.unregister:
			if h.connections[c.data.RoomID][c] {
				delete(h.connections[c.data.RoomID], c)
				close(c.message)
			}
		case data := <-h.broad:
			var con connection
			err := json.Unmarshal(data,&con)
			if err != nil {
				log.Println(err)
				return
			}
			for c := range h.connections[con.data.RoomID] {
				select {
				case c.message <- []byte(con.data.Content):
				default:
					//防止死循环
					delete(h.connections[con.data.RoomID],c)
					close(c.message)
					if len(h.connections[con.data.RoomID]) == 0 {
						delete(h.connections, con.data.RoomID)
					}
				}
			}
		}
	}
}
