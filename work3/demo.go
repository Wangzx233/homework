package main

import (
	"fmt"
)
func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int )
	quit := make(chan int)
	go func() {
		for i := 0; i < 10 ; i++ {
			// 接受通道c传来的值，并打印到控制台
			fmt.Println(<-c)
		}
		// 当协程执行完上述操作后，向quit发送数据
		quit <- 0
	}()
	fibonacci(c, quit)
}