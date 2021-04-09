package socket

type Data struct {
	//对象地址
	Type     string `json:"type"`
	Content  string `json:"content"`
	User     string `json:"user"`
	RoomID   string `json:"room_id"`
	UserList []string
}
