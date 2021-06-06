package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var roomReady = make(map[string]int)
var NowPlayer = make(map[string]string)
var	Checkerboard = make(map[string][size][size]string)
var	Prepared=make(map[string]bool)
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
		//创建房间
		case "create":

			m:=make(map[string]chan Message)
			m[msg.User] = c.MsgChan
			Mng.Chans[msg.RoomID]=mapAppend(m,Mng.Chans[msg.RoomID])
			Mng.Broad(Message{
				Type:    "message",
				User:    "",
				Content: msg.User+"上线了",
				RoomID:  msg.RoomID,
			})

		//加入房间
		case "login":

			m:=make(map[string]chan Message)
			m[msg.User] = c.MsgChan
			Mng.Chans[msg.RoomID]=mapAppend(m,Mng.Chans[msg.RoomID])
			Mng.Broad(Message{
				Type:    "message",
				User:    "",
				Content: msg.User+"上线了",
				RoomID:  msg.RoomID,
			})
		case "broad":
			Mng.Broad(Message{
				Type:    "message",
				User:    msg.User,
				Content: msg.Content,
				RoomID:  msg.RoomID,
			})
		case "logout":
			Mng.Broad(Message{
				Type:    "message",
				User:    "",
				Content: msg.User+"下线了",
				RoomID:  msg.RoomID,
			})
			Mng.Logout(c.User,msg.RoomID)
		case "ready":
			roomReady[msg.RoomID]++
			Prepared[msg.User]=true
			Mng.Broad(Message{
				Type:    "message",
				User:    "",
				Content: msg.User+"已准备",
				RoomID:  msg.RoomID,
			})
			if roomReady[msg.RoomID]>=2 {
				Mng.GameStart(Message{
					Type:    "gameStart",
					User:    "",
					Content: "",
					RoomID:  msg.RoomID,
				})

				NowPlayer[msg.RoomID]="1"
			}
			fmt.Println(roomReady[msg.RoomID],NowPlayer)
		case "unready":
			roomReady[msg.RoomID]--
			Prepared[msg.User]=true
			Mng.Broad(Message{
				Type:    "message",
				User:    "",
				Content: msg.User+"已取消准备",
				RoomID:  msg.RoomID,
			})
		case "move":

			if NowPlayer[msg.RoomID]==msg.User {

				split := strings.Split(msg.Content, ",")

				first,err:=strconv.Atoi(split[0])
				if err!=nil {
					log.Println(err)
				}
				second,err:=strconv.Atoi(split[1])
				if err!=nil {
					log.Println(err)
				}
				if b:=cp(first,second,NowPlayer[msg.RoomID],msg.RoomID);b!="0" {
					NowPlayer[msg.RoomID]=b
					Mng.Broad(Message{
						Type:    "next_move",
						User:    b,
						Content: msg.Content,
						RoomID:  msg.RoomID,
					})
					Mng.Broad(Message{
						Type:    "moved",
						User:    msg.User,
						Content: msg.Content,
						RoomID:  msg.RoomID,
					})
				}else {
					Mng.Broad(Message{
						Type:    "game_over",
						User:    msg.User,
						Content: "",
						RoomID:  msg.RoomID,
					})
				}
			}

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