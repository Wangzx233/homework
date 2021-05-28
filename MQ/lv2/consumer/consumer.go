package main

import (
	"fmt"
	"github.com/MashiroC/begonia"
	"github.com/MashiroC/begonia/app/option"
	"log"
)

func main() {
	c := begonia.NewClient(option.Addr(":9033"))

	// 获取一个服务
	s, err := c.Service("MQ")
	if err != nil {
		panic(err)
	}

	// 获取一个远程函数的同步调用
	testFun, err := s.FuncSync("Rec")
	if err != nil {
		panic(err)
	}

	res, err := testFun("hello")
	if err!=nil {
		log.Println(err)
	}
	fmt.Println(res)
}

