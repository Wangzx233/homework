package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myres = make(map[int]int, 20)
	mu    sync.Mutex
)

func factorial(n int,ch chan string) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	myres[n] = res
	ch<-"❤峰峰❤"
}

func main() {
	ch:=make(chan string)
	for i := 1; i <= 20; i++ {
		go factorial(i,ch)
	}

	for i, v := range myres {
		<-ch
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
	time.Sleep(time.Second*3)
}