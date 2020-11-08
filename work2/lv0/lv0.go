package main

import "fmt"

func main() {
	var a struct{
		b int
	}
	a.b=10
	fmt.Print(a.b)
}
