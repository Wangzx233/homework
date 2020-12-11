package Controller

import (
	"1/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ViewMessage(context *gin.Context)  {
	uid:=context.PostForm("uid")
	//清空MP
	var mp=make(map[string]Model.Reply)
	Model.Mp=mp

	cookie,_:=context.Request.Cookie("uid")


	Model.ViewMessage(0,uid,cookie.Value)
	context.JSON(200,Model.Mp)
}
func SendMessage(context *gin.Context)  {
	var r Model.Reply
	//被回复的用户form为to_uid,reply_id为回复“评论”,如果直接回复主页reply_id为0,
	err:=context.ShouldBind(&r)
	if err!=nil{
		fmt.Println("绑定错误")
	}
	cookie,cerr:=context.Request.Cookie("uid")
	fromUid:=cookie.Value
	if cerr!=nil{
		fmt.Println("获取cookie失败")
	}
	
	anonymous:=context.PostForm("anonymous")
	if anonymous=="true"{
		fromUid="匿名用户"
	}
	Model.SendMessage(r.Reply_id,r.Content,fromUid,r.To_uid,r.Power)
}
func DeleteMessage(context *gin.Context)  {
	mid:=context.PostForm("id")
	id,err:=strconv.Atoi(mid)
	if err!=nil{
		fmt.Println("转换字符串失败",err)
	}
	cookie,err:=context.Request.Cookie("uid")
	if err!=nil{
		fmt.Println("获取cookie失败")
	}
	bo:=Model.DeleteMessage(id,cookie.Value)
	if !bo {
		context.Writer.Write([]byte("删除失败，权限不足"))
	}else {
		context.Writer.Write([]byte("删除成功"))
	}
}
func Like(context *gin.Context)  {
	id:=context.PostForm("reply_id")

	replyId,err:=strconv.Atoi(id)
	if err!=nil{
		fmt.Println("转换字符串失败:",err)
	}
	Model.Like(replyId)
}