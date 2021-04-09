package model

import "demo/struct"

func Register(username string, password string) (bool, string) {
	err := DB.Where("user_name = ?", username).First(&_struct.User{}).Error
	if err == nil {
		return false, "用户已存在"
	}
	err = DB.Create(&_struct.User{UserName: username, Password: password}).Error
	if err!=nil {
		return false, "创建失败"
	}else {
		return true,"创建成功"
	}

}
func Login(username string, password string) bool {
	err := DB.Where("user_name = ? and password = ?", username, password).First(&_struct.User{}).Error
	if err == nil {
		return true
	}
	return false
}

