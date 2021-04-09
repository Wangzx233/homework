package socket

import "log"

func (c *connection) writer() {
	for message := range c.message {
		err := c.ws.WriteJSON(message)
		if err != nil {
			log.Println(err)
			return
		}
		err = c.ws.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

