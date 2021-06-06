package socket

import (
	"strconv"
)

type Manager struct {
	Chans map[string]map[string]chan Message
}

//房间内广播
func (m *Manager) Broad(msg Message) {
	for _, ch := range m.Chans[msg.RoomID] {
			ch <- msg
	}
}

func (m *Manager) GameStart(msg Message) {
	i:=1
	for _, ch := range m.Chans[msg.RoomID] {
		user:=strconv.Itoa(i)
		msg.User=user
		ch <- msg
		i++
	}
}

func (m *Manager) GameOver(msg Message) {
	for _, ch := range m.Chans[msg.RoomID] {
		ch <- msg
	}
	delete(Checkerboard,msg.RoomID)
}
//房间已满
func RoomSpill(c chan Message,msg Message) {
	c<-msg
}

//退出登录
func (m *Manager) Logout(u,roomId string)  {

	if Prepared[u]==true {
		Prepared[u]=false
		roomReady[roomId]--
	}
	delete(Checkerboard,roomId)
	delete(Mng.Chans[roomId],u)
	i:=0
	for range Mng.Chans[roomId]{
		i++
	}
	if i==0 {
		delete(Mng.Chans,roomId)
	}
	if m.Chans[roomId][u]!=nil {
		close(m.Chans[roomId][u])
	}

}

//准备就绪
func (m *Manager) Ready(u,roomId string)  {
	delete(Mng.Chans[roomId],u)
	close(m.Chans[roomId][u])
}



var Mng =Manager{
Chans: make(map[string]map[string]chan Message),
}

