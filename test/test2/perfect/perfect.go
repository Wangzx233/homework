package main

import (
	"fmt"
)

func main() {
	for i:=1;i<=123456;i++{
		go func(i int) {
			if perfect(i){
				fmt.Println(i)
			}
		}(i)
	}
}
func perfect(i int) bool {
	s:=0
	for j:=1;j<=i/2;j++{
		if i%j==0{
			s+=j
		}
	}
	if s==i {
		return true
	}
	return false
}