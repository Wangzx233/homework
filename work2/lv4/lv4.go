package main

import "fmt"

func main() {

	paint(cycle{})
}
type Painter interface {
	Exp(x,y float64) (bool)
}
type heart struct {
}
type cycle struct {
}

func (h heart) Exp(x,y float64) (bool)  {
	return (x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y<0
}
func (c cycle) Exp(x,y float64) bool {
	return x*x+y*y<0.5
}
func paint(painter Painter){					//画心型的函数
	var x,y float64
	for y=1.5;y>=-1;y=y-0.2{
		for x=-1.5;x<1.5;x=x+0.08{
			b:=painter.Exp(x,y)
			if  b {
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
