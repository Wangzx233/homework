package main

import (
	"fmt"
	"github.com/MashiroC/begonia"
	"github.com/MashiroC/begonia/app/option"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	SqlInit()

	s := begonia.NewServer(option.Addr("localhost:12306"), option.P2P())

	s.Register("user_center", &UserCenter{})

	s.Wait()
}

type UserCenter struct{}

type User struct {
	gorm.Model
	UserName string `form:"username" json:"username"` //binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password"` //binding:"required,min=8,max=40"`
}
func (*UserCenter) SayHello(name string) string {
	fmt.Println("sayHello")
	return "😀Hello " + name
}

func (*UserCenter) Register(username string, password string) (bool, string) {
	err := DB.Where("user_name = ?", username).First(&User{}).Error
	if err == nil {
		return false, "用户已存在"
	}
	err = DB.Create(&User{UserName: username, Password: password}).Error
	if err != nil {
		return false, "创建失败"
	}
	return true, "创建成功"
}
func (*UserCenter) Login(username string, password string) (bool) {
	var user User
	err := DB.Where("user_name = ? and password = ?", username, password).First(&user).Error
	if err == nil {
		return true
	}
	return false
}

func (*UserCenter) ChangePassword(username string, newPassword string) error {
	var user User
	update := DB.Model(user).Where("user_name = ?", username).Update("password", newPassword).Error
	if update != nil {
		return update
	}
	return nil
}

var DB *gorm.DB

const (
	dns = "root:_aA2664190827@tcp(localhost:3306)/five_data?charset=utf8mb4&parseTime=True&loc=Local"
)
func SqlInit() {
	db, err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	//db = db.Debug()
	if err != nil {
		log.Panicln(err)
		return
	}
	if db.AutoMigrate(&User{}) != nil {
		log.Println(err)
		return
	}
	DB = db
}
