package main

import (
	"fmt"
	"time"
)

func main() {
	ch:=make(chan string)

		go p1(ch)
		go p2(ch)
	time.Sleep(time.Second*3)
}
func p1(ch chan string){
	for i:=1;i<=10;i++ {
		ch<-"123"
		if i%2==1  {
			fmt.Println("p1:", i)
		}
	}
}
func p2(ch chan string){
	for i:=1;i<=10;i++ {
		<-ch
		if i%2==0 {
			fmt.Println("p2:", i)
		}

	}

}