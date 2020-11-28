package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)
func main() {
	r := gin.Default()
	r.Use(state())

	//登录
	r.POST("/login", func(context *gin.Context) {
		var person Person
		context.ShouldBind(&person)

		//读取存储账号密码的文件
		Allperson := make(map[string]Person)
		fp, _ := os.OpenFile("person.json",os.O_RDWR|os.O_APPEND,0644)
		defer fp.Close()
		bytes, _ := ioutil.ReadAll(fp)
		json.Unmarshal(bytes, &Allperson)

		fmt.Println(person)
		//判断输入的账号密码是否正确
		t:=Dl(person.Name,person.Password,Allperson)
		switch t {
		case 2:
			context.SetCookie("user_cookie", person.Name, 1000, "/", "localhost", false, true)
			context.Writer.Write([]byte("hello,"+person.Name))
		case 1:
			context.Writer.Write([]byte("账号错误"))
		case 0:
			context.Writer.Write([]byte("密码错误"))
		}
	})
	//注册
	r.POST("/register", func(context *gin.Context) {
		var person Person
		context.ShouldBind(&person)

		//读取存储账号密码的文件
		Allperson := make(map[string]Person)
		fp, _ := os.OpenFile("person.json",os.O_RDWR|os.O_APPEND,0644)
		defer fp.Close()
		bytes, _ := ioutil.ReadAll(fp)
		json.Unmarshal(bytes, &Allperson)

		//判断账号是否存在
		zhbo:=true
		for _,value:= range Allperson{
			if value.Name==person.Name{
				context.Writer.Write([]byte("账号已存在，请重新输入"))
				zhbo=false
			}
		}
		//如果账号不存在就将用户写入文件
		Allperson[person.Name]=person
		if zhbo{
			if (len(person.Name) > 0 && len(person.Password) > 0) {
				os.Truncate("person.json",0)
				fp.Seek(0,0)
				h,_:=json.Marshal(Allperson)
				fp.WriteString(string(h))
				fmt.Println("成功注册")
				context.Writer.Write([]byte("注册成功"))
			}
		}
	})

	//登录注册以外的界面
	r.GET("/other", func(context *gin.Context) {
	})
	r.Run()
}

type Person struct {
	Name string	`form:"name"`
	Password string	`form:"password"`
}
//登录函数
func Dl(zh string,mm string,mp map[string]Person)int{
	zhbo:=false
	mmbo:=false
	for _, value := range mp {
		if value.Name == zh {
			zhbo = true
			fmt.Println("账号正确")
			if value.Password==mm{
				mmbo = true
				fmt.Println("密码正确")
			}
		}
	}
	if !zhbo{
		return 1
	}else if !mmbo{
		return 0
	}
	return 2
}
func state() gin.HandlerFunc{
	 return func(context *gin.Context) {
		 context.Next()
		cookie,err:=context.Request.Cookie("user_cookie")
		if err==nil{
			context.JSON(200,gin.H{
				"code":200,
				"message":cookie.Value+"你好！",
			})
		}else {
			context.JSON(200,gin.H{
				"code":200,
				"message":"游客你好！",
			})
		}
	}
}