package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
var db *sql.DB
var err  error

func initdb()  {
	db,err=sql.Open("mysql","root:root@tcp(localhost:3306)/demo")
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
	//update()
	fqueryRow(2)
	fquery()
	//insert()
}
//单行查询
func fqueryRow(id int){
	initdb()
	defer db.Close()
	var u1 user
	sqlStr:="select id,name,password,num from u_1 where id=?;"
	rowObj:=db.QueryRow(sqlStr,id)
	rowObj.Scan(&u1.id,&u1.name,&u1.password,&u1.num)
	fmt.Printf("%#v",u1)
}
//插入数据
func insert() {
	initdb()
	defer db.Close()
	sqlStr:=`insert into u_1(id,name,password,num) values(3,"王五","123",3);`
	ret,err:=db.Exec(sqlStr)
	if err!=nil{
		log.Println(err)
		return
	}
	id,err:=ret.LastInsertId()
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("id:",id)
	return
}
//跟新数据
func update() {
	initdb()
	defer db.Close()
	sqlStr:=`update u_1 set name="王六" where id="3";`
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
	fmt.Println("修改的行数为:",n)
	return
}
//多行查询
func fquery(){
	initdb()
	defer db.Close()

	sqlStr:="select id,name,password,num from u_1;"
	rowObj,err:=db.Query(sqlStr)
	if err!=nil{
		fmt.Printf("查询出错\n")
		return
	}
	for rowObj.Next(){
		var u1 user
		err:=rowObj.Scan(&u1.id,&u1.name,&u1.password,&u1.num)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("%#v\n",u1)
	}
}
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
type user struct {
	id int
	name string
	password string
	num int
}