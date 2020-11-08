package main

import "fmt"

func main() {
	var a string
	fmt.Println("请输入cycle或heart来输出圆或心：")
	fmt.Scan(&a)
	paint(a)
}
//type Painter interface {
//	ht()
//}
func ht(){					//画心型的函数
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
func cycle()  {							//画圆
	var x,y float64
	for y=1.5;y>=-1;y=y-0.2{
		for x=-1.5;x<1.5;x=x+0.08{
			if(x*x+y*y<1){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func paint(painter string)  {
	switch painter {
	case "cycle":
		cycle()
	case "heart":
		ht()
	}
}