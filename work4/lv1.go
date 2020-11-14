package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func init()  {
	println("***警告:缝合怪出没！***")
}
func main() {
	var start=time.Now()
	var h string
	var f string
	f="3.1415926"
	var idx string
	idx="hello_world"
	h="2333"
	//map
	m:= make(map[int]string)
	for i,value:=range m{
		fmt.Println(i,value)
	}
	//panic&&recover
	err()
	fmt.Println("以下为字符串的转换：")
	fmt.Printf("字符串“%s”转化为十进制数是：",h)
	fmt.Println(strconv.ParseInt(h, 10, 0))
	fmt.Printf("字符串“%s”转化为十六进制数用十进制表示为是：",h)
	fmt.Println(strconv.ParseInt(h, 16, 0))
	fmt.Printf("字符串“%s”转化为八进制数用十进制表示为是：",h)
	fmt.Println(strconv.ParseInt(h, 8, 0))
	fmt.Println("字符串被转化成了十进制")
	a,_:=strconv.ParseInt(h, 10, 0)
	fmt.Println("十进制被转化成了字符串")
	fmt.Println(strconv.FormatInt(a,10))
	fmt.Printf("字符串“%s”转化为float是：",f)
	fmt.Println(strconv.ParseFloat(f,32))
	if strings.Contains(idx,"_"){
		fmt.Print("“hello_world“中的”_“在第")
		fmt.Print(strings.Index(idx,"_"))
		fmt.Println("个")
	}
	if strings.EqualFold(idx,h){
		fmt.Println("h和idx相等")
	}else {
		fmt.Println("h和idx不相等")
	}
	fmt.Print(idx)
	fmt.Println("由",strings.Split(idx,"_"),"组成")
	fmt.Println(strings.Replace(idx, "h", "H", -1))
	fmt.Println(idx)
	if strings.HasPrefix(idx,"hallo"){
		fmt.Println(idx)
	}
	if strings.HasSuffix(idx,"world"){
		fmt.Println(idx)
	}
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now().Format(time.Kitchen))
	fmt.Println(time.Now().Unix())

	fmt.Println(time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51"))
	fmt.Println(time.Hour)
	fmt.Println(time.Second)
	fmt.Println(time.Hour*25)
	fmt.Println(time.Wednesday*2)
	fmt.Println(time.Minute*61)
	fmt.Println(time.Nanosecond*10)
	fmt.Println(time.Nanosecond*1000)
	t, _ := time.ParseDuration("1.5s")
	fmt.Println(t.Seconds(), t.Nanoseconds())
	fmt.Println("程序耗时:",time.Now().Sub(start))
	ticker := time.NewTicker(time.Second * 1) // 运行时长
	ch := make(chan int)
	go func() {
		var x int
		for x < 5 {
			select {
			case <-ticker.C:
				x++
				fmt.Printf("%d\n", x)
			}
		}
		ticker.Stop()
		ch <- 0
	}()
	fmt.Println(<-ch)
}

//函数作为类型
func f(f1 func())func(){
	defer println(1)
	return func() {
	}
}

func err(){
	defer func() {
		err:=recover()
		fmt.Println(err)
		fmt.Println("错误解决了，继续执行")
	}()
	panic("出现错误啦")
}

