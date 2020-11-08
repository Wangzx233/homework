package main

import "fmt"
func main() {
	var a,b,s float32
	var fh byte
	for true{
		fmt.Println("putinto:")
		fmt.Scanf("%f%c%f",&a,&fh,&b)
		if(a<0||a>9){
			break
		}
		s=count(a,fh,b)
		fmt.Println("putout:")
		fmt.Println(s)
	}
}
func count(a1 float32,fh byte,b1 float32) float32{ //计算两个数的加减乘除
	var s float32
	switch fh {
		case '*':
			s=a1*b1
		case '+':
			s=a1+b1
		case '-':
			s=a1-b1
		case '/':
			s=a1/b1
	}
	return s
}