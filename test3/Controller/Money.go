package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"test3/Model"
)

func Recharge(context *gin.Context)  {
	cookie,err:=context.Request.Cookie("uid")
	if err!=nil{
		fmt.Println("获取cookie失败")
		return
	}
	uid:=cookie.Value
	money:=context.PostForm("money")
	m,err:=strconv.ParseFloat(money,64)
	if err!=nil{
		fmt.Println("转换字符串失败")
	}
	Model.Recharge(uid,m)
}

func Transfer(context *gin.Context)  {
	cookie,err:=context.Request.Cookie("uid")
	if err!=nil{
		fmt.Println("获取cookie失败")
		return
	}
	uid:=cookie.Value
	money:=context.PostForm("money")
	toUid:=context.PostForm("toUid")
	remarks:=context.PostForm("remarks")

	m,err:=strconv.ParseFloat(money,64)
	if err!=nil{
		fmt.Println("转换字符串失败")
	}
	Model.Transfer(uid,toUid,m,remarks)
}