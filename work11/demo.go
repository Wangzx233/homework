package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Password string
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/data1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db = db.Debug()
	if err != nil {
		log.Panicln(err)
		return
	}
	DB = db
}

func main() {
	err := DB.AutoMigrate(&User{})
	if err!=nil{
		log.Panicln(err)
		return
	}
	user := []User{}
	//DB.Create(&User{Name: "12",Password: "12345"})
	find := DB.Where("name=? and password=?",12,1234).Find(&user)
	err = find.Error
	if err!=nil{
		log.Panicln(err)
	}

	fmt.Println(user)
}
