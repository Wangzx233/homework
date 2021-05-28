package main

import (
	"fmt"
	"github.com/MashiroC/begonia"
	"github.com/MashiroC/begonia/app/option"
)

type Broker struct {
}

type message struct {
	topic   string
	content string
}

func main() {

	s := begonia.NewServer(option.Addr("localhost:9033"), option.P2P())

	s.Register("MQ", &Broker{})
	go center()
	s.Wait()
}

var Messages = make(chan message, 20)

func (*Broker) SendMessage(topic string, content string) {
	var m message
	m.topic = topic
	m.content = content
	Messages <- m
	fmt.Println("成功发送消息")
}

var Msg = make(map[string]string)

func center() {

	for {
		select {
		case m := <-Messages:
			Msg[m.topic] = m.content
			fmt.Println(Msg)
		}
	}
}
func (*Broker) Rec(topic string) (msg []string) {
	for k, m := range Msg {
		if k == topic {
			msg = append(msg, m)
		}
	}
	fmt.Println(msg)
	return msg
}
