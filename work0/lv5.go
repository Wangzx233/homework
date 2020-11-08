package main

import "fmt"
func main() {
var a string
var b[30] string
var tou1,wei1 int
j:=0
fmt.Scan(&a)
s:=a[0:len(a)]
for tou:=0;tou<len(a);tou++{
	for wei:=len(a)-1;wei>tou;wei--{
		if(s[tou:tou+1]==s[wei:wei+1]){
			tou1=tou+1
			wei1=wei-1
			for ;tou1<wei1;{
				if(s[tou1:tou1+1]!=s[wei1:wei1+1]){
					break
				}
				tou1++
				wei1--
			}
			if(tou1>=wei1){
				b[j]=s[tou:wei+1]
				j++
			}
		}
	}
}
fmt.Println("其中所有的回文字符串：",b)
var zj string
for i:=0;i<len(b);i++{
	if(len(b[0])<len(b[i])){
		zj=b[0]
		b[0]=b[i]
		b[i]=zj
	}
}
fmt.Print("其中最长的回文字符串：",b[0])
}
