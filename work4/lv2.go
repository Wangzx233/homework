package main

import (
	"fmt"
	"time"
)
var now = time.Now()
func main() {
	t3 := time.NewTicker(time.Hour * 1)
	go func() {
		for {
			select {
			case <-t3.C:
				fmt.Println("芜湖！起飞！")
			}
		}
	}()


		go func(){
		 for {

			 next1 := now.Add(time.Hour * 24)
			 next1 = time.Date(next1.Year(), next1.Month(), next1.Day(), 2, 0, 0, 0, next1.Location())
			 t1 := time.NewTimer(next1.Sub(now))
			 <-t1.C
			 fmt.Printf("%v 晚安，玛卡巴卡\n", time.Now())
			 t1.Reset(time.Hour*24)

		 }
	 }()
	go func() {
		for {

			next2 := now.Add(time.Hour * 24)
			next2 = time.Date(next2.Year(), next2.Month(), next2.Day(), 8, 0, 0, 0, next2.Location())
			t2 := time.NewTimer(next2.Sub(now))
			<-t2.C
			fmt.Printf("%v 早安，打工人！\n", time.Now())
			t2.Reset(time.Hour*24)
		}
	}()
	time.Sleep(time.Hour*9999)
}
