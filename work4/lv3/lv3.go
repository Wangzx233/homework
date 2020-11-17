package main

import (
	"encoding/json"
	"fmt"
	//"encoding/json"
	"io/ioutil"
	"os"
)
type Person struct {
	Id string
	Mm string
}
func main() {
	var qr int
	var zh,mm string
	Allperson := make(map[string]Person)
	fp, _ := os.OpenFile("user.json",os.O_RDWR|os.O_APPEND,0644)
	defer fp.Close()
	bytes, _ := ioutil.ReadAll(fp)
	json.Unmarshal(bytes, &Allperson)
	for {
		fmt.Println("登录输入1，注册输入2，退出输入3")
		fmt.Scan(&qr)
		if qr == 1 {
			dl(Allperson)
		}
		if qr == 2 {
			zhbo:=true
			for zhbo{
				zhbo=false
				zh, mm = sr()
				for _,value:= range Allperson{
					if value.Id==zh{
						zhbo=true
						fmt.Println("账号已存在，请重新输入：")
					}
				}
			}

			Allperson[zh]=Person{zh,mm}
			fmt.Println(Allperson)
			fmt.Println("数据已保存，退出可实现注册")
		}
		if qr == 3 {
			if (len(zh) > 0 && len(mm) > 0) {
				os.Truncate("user.json",0)
				fp.Seek(0,0)
				h,_:=json.Marshal(Allperson)
				fp.WriteString(string(h))
				fmt.Println("成功注册")
				break
			}
		}
	}

}

func sr()(string,string)  {
	var zh,mm string
	for {
		fmt.Println("请输入账号密码：")
		fmt.Scan(&zh)
		fmt.Scan(&mm)
		if len(mm)>=6{break}else {fmt.Println("密码长度不能小于6")}
	}
	return zh,mm
}
func dl(mp map[string]Person)int{
	var zh,mm string
	zhbo:=false
	mmbo:=false
	for !(zhbo&&mmbo) {
		zh, mm = sr()
		for _, value := range mp {
			if value.Id == zh {
				zhbo = true
				if value.Mm==mm{
					mmbo = true
				}
			}
		}
		if !zhbo{
			fmt.Println("账号错误")
		}else if !mmbo{
			fmt.Println("密码错误")
		}
	}
	fmt.Println("登录成功")
	return 2
}