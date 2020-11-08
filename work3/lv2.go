package main

import (
	"fmt"
	"os"
)

func main() {
	//filObj,err:=os.OpenFile("./proverb.txt",os.O_CREATE|os.O_TRUNC,0)
	filObj,err:=os.Create("proverb.txt")
	if err!=nil{
		fmt.Printf("打开或创建文件错误%v",err)
		return
	}
	filObj.Write([]byte("don't communicate by sharing memory share memory by communicating"))

	var b = make([]byte,128)
	n,err1:=filObj.Read(b[:])
	if err!=nil{
		fmt.Printf("读取错误%v",err1)
		return
	}
	fmt.Printf("读了%d个",n)
	fmt.Println(string(b[:n]))

	filObj.Close()
}
