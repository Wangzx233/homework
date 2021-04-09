package main

import (
	"work12/router"
	socket "work12/websocket"
)

func main() {
	router.Router()
	go socket.H.RUN()
}