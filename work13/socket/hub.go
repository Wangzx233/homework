package socket

import "strings"

type Manager struct {
	Chans map[string]map[string]chan Message
}

func (m *Manager) Broad(msg Message) {
	for _, ch := range m.Chans[msg.RoomID] {
		if mgc(msg.Content) {
			ch <- msg
		}
	}
}
func (m *Manager) Logout(u,roomId string)  {
	delete(Mng.Chans[roomId],u)
	close(m.Chans[roomId][u])
}



var Mng =Manager{
Chans: make(map[string]map[string]chan Message),
}

//敏感词过滤
func mgc(s string) bool {
	if find := strings.Contains(s, "敏感词");!find {
		return true
	}
	return false
}