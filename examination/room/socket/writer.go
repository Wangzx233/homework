package socket

import "log"

func (c *Client) Write()  {

	for {
		select {
		case m := <-c.MsgChan:
			err := c.Conn.WriteJSON(m)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

