package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)
var db *sql.DB
var err  error

func initdb()  {
	db,err=sql.Open("mysql","root:root@tcp(localhost:3306)/data")
	if err!=nil{
		fmt.Println("dsn格式不正确")
		return
	}
	err=db.Ping()
	if err!=nil{
		fmt.Println("数据库链接失败！")
		return
	}
	db.SetMaxOpenConns(10)
	fmt.Println("数据库链接成功！")
	return
}
func main() {
	r := gin.Default()
	r.Use(state())

	//登录
	r.POST("/login", func(context *gin.Context) {
		//检测用户是否已登录
		_,err:=context.Cookie("user_cookie")
		if err==nil{
			fmt.Println("用户已登录")
			context.Request.URL.Path="/detail"
			return
		}

		var user User
		context.ShouldBind(&user)

		//判断输入的账号密码是否正确
		if !Dl(user.Id,user.Password){
			context.JSON(200,gin.H{
				"erro":"账号或密码错误",
			})
			return
		}
		context.SetCookie("user_cookie", user.Id, 1000, "/", "localhost", false, true)
		context.JSON(200,gin.H{
			"成功":"账号密码正确",
		})

		context.Request.URL.Path="/detail"
	})
	//注册
	r.POST("/register", func(context *gin.Context) {
		var user User
		err:=context.ShouldBind(&user)
		if err!=nil{
			log.Println(err)
			return
		}
		//判断账号是否存在
		if fqueryRow(user.Id)==false{
			U_insert(user)
		}else {
			fmt.Println("用户已存在")
		}
	})
	//个人详细页
	r.POST("/personal", func(context *gin.Context) {
		//检测用户是否已登录
		cookie,err:=context.Cookie("user_cookie")
		if err!=nil{
			fmt.Println("用户未登录")
			context.Request.URL.Path="/login"
			return
		}
		request:=context.PostForm("request")

		//退出登录
		if request=="quit"{
			context.SetCookie("user_cookie",cookie, -1, "/", "localhost", false, true)
			context.JSON(200,gin.H{
				"message":"已退出登录！",
			})
		}
		//如果点击修改密码跳转到修改密码页
		if request=="change"{
			context.Request.URL.Path="/change"
		}
	})
	//他人主页
	r.POST("/detail", func(context *gin.Context) {
		to_id:=context.Query("id")
		//检测用户是否已登录
		cookie,err:=context.Cookie("user_cookie")
		if err!=nil{
			fmt.Println("用户未登录")
			context.Request.URL.Path="/login"
			return
		}
		request:=context.PostForm("request")
		content:=context.PostForm("content")

		//当用户点击的留言(这里的留言是动词)时
		if request=="reply"{
			reply_id:=context.PostForm("reply_id")
			r_id,_:=strconv.Atoi(reply_id)
			R_insert(r_id,content,cookie,to_id)
		}
		//单击显示留言时，（当reply_id是否为0即可判断单击的是主页的显示留言，不为0即可判断单击的是留言的留言）
		if request=="dpy_rp"{
			//接收用户单击的哪个留言下的显示留言,为0即主页下的留言
			rp:=context.PostForm("reply")
			reply,_:=strconv.Atoi(rp)
			mp:=Mquery(reply)
			context.JSON(200,mp)
		}

	})
	//修改名字密码页面
	r.POST("/change", func(context *gin.Context) {
		//检测用户是否已登录
		cookie,err:=context.Cookie("user_cookie")
		if err!=nil{
			context.Writer.Write([]byte("用户未登录"))
			context.Request.URL.Path="/detail"
			return
		}
		ty:=context.PostForm("type")
		//修改
		if ty=="name"{
			name:=context.PostForm("name")
			update(cookie,name)
		}
		if ty=="password"{
			password:=context.PostForm("password")
			updatep(cookie,password)
		}
	})
	r.Run()
}
//状态
func state() gin.HandlerFunc{
	return func(context *gin.Context) {
		context.Request.Cookie("user_cookie")
	}
}
//登录函数
func Dl(id string,password string)bool{
	initdb()
	defer db.Close()
	sqlStr:="select id from user where id=?,password=?;"
	err:=db.QueryRow(sqlStr,id,password).Scan()
	if err==sql.ErrNoRows{
		return false
	}
	return true
}

//单行查询
func fqueryRow(id string)bool{
	initdb()
	defer db.Close()
	sqlStr:="select id from user where id=?;"
	err:=db.QueryRow(sqlStr,id).Scan()
	if err==sql.ErrNoRows{
		return false
	}
	return true
}
//插入用户表数据
func U_insert(user User) {
	initdb()
	defer db.Close()
	sqlStr:=`insert into user (id,name,password,sign) values(?,?,?,?);`
	_,err:=db.Exec(sqlStr,user.Id,user.Name,user.Password,user.Sign)
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("插入数据成功！")
	return
}

//留言板插入数据
func R_insert(reply_id int,content string,from_uid string,to_uid string) {
	initdb()
	defer db.Close()
	sqlStr:=`insert into reply (reply_id,content,from_uid,to_uid) values(?,?,?,?);`
	_,err:=db.Exec(sqlStr,reply_id,content,from_uid,to_uid)
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("插入数据成功！")
	return
}
//跟新数据
func updatep(id string,password string) {
	initdb()
	defer db.Close()
	sqlStr:=`update user set password=? where id=?;`
	db.Exec(sqlStr,password,id)
	return
}
func update(id string,name string) {
	initdb()
	defer db.Close()
	sqlStr:=`update user set name=? where id=?;`
	db.Exec(sqlStr,name,id)
	return
}

//查询留言
func Mquery(reply_id int)map[int]reply{
	initdb()
	defer db.Close()
	mp:=make(map[int]reply)
	sqlStr:="select * from reply where reply_id=?;"
	rowObj,err:=db.Query(sqlStr,reply_id)
	if err!=nil{
		fmt.Printf("查询出错\n")
		return mp
	}
	for rowObj.Next(){

		var rp reply
		err:=rowObj.Scan(&rp.Id,&rp.Update_time,&rp.Reply_id,&rp.Content,&rp.From_uid,&rp.To_uid)
		if err!=nil{
			log.Fatal(err)
		}

		mp[rp.Id]=rp
	}
	return mp
}
//删除
func delete()  {
	initdb()
	defer db.Close()
	sqlStr:=`delete from u_1 where id="3";`
	ret,err:=db.Exec(sqlStr)
	if err!=nil{
		log.Println(err)
		return
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("删除的行数为:",n)
	return
}
//留言结构
type reply struct {
	Id int	`form:"id"`
	Update_time string
	Reply_id int
	Content string
	From_uid string
	To_uid string
}
//用户结构
type User struct {
	Id string `form:"id"`
	Name string	`form:"name"`
	Password string	`form:"password"`
	Sign string	`form:"sign"`
}