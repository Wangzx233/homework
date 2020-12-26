package main

import (
	"fmt"
	"math"
)

func main() {
	for i:=1;i<=123456;i++{
		go func(i int) {
			if prime(i){
				fmt.Println(i)
			}
		}(i)
	}
}

func prime(i int) bool {
	//优化：判断素数只需要除至i的方，节省了时间
	x:=math.Sqrt(float64(i))
	for j:=2;float64(j)<=x;j++{
		if i%j==0{
			return false
		}
	}
	return true
}

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(z*z-x) > 0.000001{
		z -= (z*z -x) / (2*z)
	}
	return z
}