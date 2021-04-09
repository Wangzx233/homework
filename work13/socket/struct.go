package socket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	User    string
	Conn    *websocket.Conn
	MsgChan chan Message
}

type Message struct {
	Type     string   `json:"type"`
	User     string   `json:"user"`
	Content  string   `json:"content"`
	RoomID   string   `json:"room_id"`
	//pb_users []string `json:"pb"`
}
