package main

import "fmt"
func main() {
	var a = []int{23,34,65,24,64,24,54}
	for i:=0;i<len(a);i++{
		for j:=i;j<len(a);j++{
			if(a[i]<a[j]){
				a[i]=a[i]+a[j]
				a[j]=a[i]-a[j]
				a[i]=a[i]-a[j]
			}
		}
	}
	fmt.Print(a)
}
