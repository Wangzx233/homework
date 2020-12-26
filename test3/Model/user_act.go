package Model

import "fmt"

func Login(uid string,password string) bool {
	sqlStr:="select uid,password from user where uid=?"
	row,err:=DB.Query(sqlStr,uid)
	if err!=nil{
		fmt.Println("false")
		return false
	}
	defer row.Close()
	var u User
	for row.Next() {
		err = row.Scan(&u.Uid, &u.Password)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	if u.Password == password {
		return true
	}
	return false
}
func Register(uid string,password string) bool {
	row,err:=DB.Query("select uid from user where uid=?",uid)
	defer row.Close()
	var u User
	for row.Next() {
		err = row.Scan(&u.Uid)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	if u.Uid==uid {
		return false
	}
	R_update(uid,password)
	return true
}
//新建用户
func R_update(uid,password string)  {

	_,err:=DB.Exec("insert user (uid,password) values (?,?)",uid,password)
	if err!=nil{
		fmt.Printf("Exec failed: %v",err)
	}
}
