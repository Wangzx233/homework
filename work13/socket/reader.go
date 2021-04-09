package socket

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *Client) Reader() {

	for {

		_, byte, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		msg := Message{}
		err = json.Unmarshal(byte, &msg)
		if err != nil {
			log.Println(err)

		}
		//fmt.Printf("%v",msg)
		switch msg.Type {
		case "login":
			m:=make(map[string]chan Message)
			m[msg.User] = c.MsgChan
			Mng.Chans[msg.RoomID]=mapAppend(m,Mng.Chans[msg.RoomID])
			Mng.Broad(Message{
				Type:    "",
				User:    "",
				Content: msg.User+"上线了",
				RoomID:  msg.RoomID,
			})
		case "broad":
			Mng.Broad(msg)
		case "logout":
			Mng.Broad(Message{
				Type:    "",
				User:    "",
				Content: msg.User+"下线了",
				RoomID:  msg.RoomID,
			})
			Mng.Logout(c.User,msg.RoomID)
		default:
			fmt.Println("else")
		}
	}
}

//合并map
func mapAppend(m1,m2 map[string]chan Message) map[string]chan Message {
	n := make(map[string]chan Message)
	for i, v := range m1 {
		n[i]=v
	}
	for j, w := range m2 {
		n[j]=w
	}
	return n
}