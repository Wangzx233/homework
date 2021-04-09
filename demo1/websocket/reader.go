package socket

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *connection) reader() {
	for {
		var data = Data{}
		err := c.ws.ReadJSON(&data)
		if err != nil {
			H.register <- c
			log.Println(err)
			return
		}
		//err = json.Unmarshal(message, &c.data)
		//if err != nil {
		//	H.register <- c
		//	log.Println(err)
		//	return
		//}

		fmt.Println("收到了")

		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			userlist = append(userlist, c.data.User)
			c.data.UserList = userlist
			marshal, err := json.Marshal(c.data)
			if err != nil {
				log.Println(err)
				return
			}



			H.broad <- marshal
			fmt.Println("连接了")
		case "normal":
			c.data.Type = "normal"
			marshal, err := json.Marshal(c.data)
			if err != nil {
				log.Println(err)
				return
			}
			H.broad <- marshal
		case "logout":
			c.data.Type = "logout"
			userlist = remove(userlist, c.data.User)
			c.data.UserList = userlist
			c.data.Content = c.data.User

			marshal, err := json.Marshal(c.data)
			if err != nil {
				log.Println(err)
				return
			}
			H.broad <- marshal
			H.unregister <- c
		default:
			fmt.Println("else")
		}
	}
}
func remove(slice []string , user string) []string {
	if len(slice) == 0 {
		return slice
	}
	if len(slice) == 1 {
		return []string{}
	}
	newSlice := []string{}
	for i := range slice {
		if slice[i] == user && i == len(slice) {
			return slice[:len(slice)]
		} else if slice[i] == user {
			newSlice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return newSlice
}
