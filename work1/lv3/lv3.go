package main

import "fmt"

func main() {
	a:=[]float32{-88,-25,-12,-8,-0.5,1,5,6,77}
	fmt.Print(check(a))
}

func check(a []float32)(jg float32){		//找出有序数组中绝对值最小的数
	var mid int

	mid=(len(a)-1)/2
	for {
		if(jdz(a[mid])<jdz(a[mid+1])&&jdz(a[mid])<jdz(a[mid-1])){
			jg=a[mid]

			break
		}else {
			if(jdz(a[mid]-1)<jdz(a[mid+1])){
				jg=a[mid-1]

				break
			}else {
					jg=a[mid+1]
					break
			}
		}
	}
	return jg
}
func jdz(s float32)(jg float32)  {		//判断绝对值
	if(s<0){
		s=-s;
	}
	jg=s
	return jg
}