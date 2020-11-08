package main

import "fmt"

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			if i == 9 {
				over <- true
			}
		}(i)						//错误1：协程与循环未绑定，导致i可能还在1，而协程已经跑了几个来拿1了。:)
		//if i == 9 {				//错误2：无缓存的channel两边都是在一个协程之中了，无法传递。👌
		//	over <- true
		//}
	}
	<-over
	fmt.Println("over!!!")
}
