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
	//电脑上没装sql就先注释了
	//SqlInit()

	s := begonia.NewServer(option.Addr("localhost:12306"), option.P2P())

	s.Register("Echo", &UserCenter{})

	s.Wait()
}

type UserCenter struct{}

func (*UserCenter) SayHello(name string) string {
	fmt.Println("sayHello")
	return "😀Hello " + name
}

func (*UserCenter) Register(username string, password string) (bool, string) {
	err := DB.Where("user_name = ?", username).First(&User{}).Error
	if err == nil {
		return false, "用户已存在"
	}
	err = DB.Create(&User{UserName: username, Password: password, UserInfo: UserInfo{Uid: username}}).Error
	if err != nil {
		return true, "创建成功"
	}
	return false, "创建失败"
}
func (*UserCenter) Login(username string, password string) (bool, int) {
	var user User
	err := DB.Where("user_name = ? and password = ?", username, password).First(&user).Error
	if err == nil {
		return true, int(user.ID)
	}
	return false, int(user.ID)
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
	dns = "root:root@tcp(127.0.0.1:3306)/data1?charset=utf8mb4&parseTime=True&loc=Local"
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

