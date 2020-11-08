package main

import "fmt"
//(x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y==0
func main() {
	var x,y float64
	for y=1.5;y>=-1;y=y-0.2{
		for x=-1.5;x<1.5;x=x+0.08{
			if((x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y<0){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

