package Model

import (
	"database/sql"
	"fmt"
)

func Recharge(uid string,money float64)  {
	row,err:=DB.Query("select balance from user where uid=?",uid)
	defer row.Close()
	if err!=nil {
		fmt.Println("recharge.query",err)
		return
	}
	var m float64
	for row.Next() {
		err = row.Scan(&m)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return
		}
	}
	m=m+money
	_,err=DB.Exec("update user set balance=? where uid=?",m,uid)
	if err!=nil{
		fmt.Printf("Exec failed: %v",err)
		return
	}
}
func Tans(conn *sql.Tx,uid string,money float64)  {
	row,err:=DB.Query("select balance from user where uid=?",uid)
	defer row.Close()
	if err!=nil {
		fmt.Println("recharge.query",err)
		conn.Rollback()
		return
	}
	var m float64
	for row.Next() {
		err = row.Scan(&m)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			conn.Rollback()
			return
		}
	}
	m=m+money
	_,err=DB.Exec("update user set balance=? where uid=?",m,uid)
	if err!=nil{
		fmt.Printf("Exec failed: %v",err)
		conn.Rollback()
		return
	}
}
func Transfer(uid,toUid string,money float64,remarks string)  {
	conn, err := DB.Begin()
	if err != nil {
		fmt.Println("开启事务失败")
		return
	}
	Tans(conn,uid,-money)
	Tans(conn,toUid,money)
	_,err=DB.Exec("insert into records (fromUid,toUid,money,remarks) values (?,?,?,?)",uid,toUid,money,remarks)
	if err!=nil{
		fmt.Printf("Exec failed: %v",err)
		conn.Rollback()
		return
	}
	conn.Commit()
}