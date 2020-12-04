package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)
func main() {
	r := gin.Default()
	r.Use(state())

	//登录
	r.POST("/login", func(context *gin.Context) {
		var person Person
		context.ShouldBind(&person)
		fmt.Println(person)
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

	//上传
	r.POST("/upload", func(context *gin.Context) {
		cookie,err:=context.Request.Cookie("user_cookie")
		if err==nil{
			file,err := context.FormFile("f1")

			var works Works
			m:=make(map[string]Works)
			fp, _ := os.OpenFile("work.json",os.O_RDWR|os.O_APPEND,0644)
			defer fp.Close()
			bytes, _ := ioutil.ReadAll(fp)
			json.Unmarshal(bytes, &m)

			context.ShouldBind(&works)
			works.user_name=cookie.Value
			if err != nil {
				context.JSON(http.StatusInternalServerError,gin.H{
					"msg":err.Error(),
				})
			}


			log.Println(file.Filename)
			//dst := fmt.Sprintf("D:\\GoProjects\\src\\homework5\\lv3\\%s\\%s",cookie.Value,file.Filename)

			context.SaveUploadedFile(file,file.Filename)

			//将作品名字和简介存入
			m[file.Filename]=works
			os.Truncate("work.json",0)
			fp.Seek(0,0)
			h,_:=json.Marshal(m)
			fp.WriteString(string(h))
			context.JSON(http.StatusOK,gin.H{
				"msg":fmt.Sprintf(`%s uploaded!`,file.Filename),
			})
		}else {
			context.Writer.Write([]byte("请先登录"))
			context.Abort()
		}

	})

	//抽奖
	r.POST("/draw", func(context *gin.Context) {
		m:=make(map[string]Works)
		fp, _ := os.OpenFile("work.json",os.O_RDWR|os.O_APPEND,0644)
		defer fp.Close()
		bytes, _ := ioutil.ReadAll(fp)
		json.Unmarshal(bytes, &m)

		y:=randMapKey(m)
		h,_:=json.Marshal(m[y])
		context.Writer.Write(h)
	})
	r.Run()
}

type Person struct {
	Name string	`form:"name"`
	Password string	`form:"password"`
}
type Works struct {
	user_name string `form:"user_name"`
	Name string	`form:"name"`
	Text string	`form:"text"`
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
//随机从map中取值
func randMapKey(m map[string]Works) string {
	r := rand.Intn(len(m))
	for k := range m {
		if r == 0 {
			return k
		}
		r--
	}
	panic("unreachable")
}