package main

import (
	"fmt"
	"rpc/server/call"
)

func main() {
	hello, _ := call.SayHello("wzx")
	fmt.Println(hello)
}
