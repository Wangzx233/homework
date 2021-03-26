package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test3/Model"
)

func QueryRemarks(context *gin.Context)  {
	remarks:=context.PostForm("remarks")
	records:=Model.QueryRemarks(remarks)
	context.JSON(http.StatusOK, gin.H{
		"id":    records.Id,
		"fromUid": records.FromUid,
		"toUid": records.ToUid,
		"time": records.Createtime,
		"remarks": records.Remarks,
	})
}
func QueryUid(context *gin.Context)  {
	cookie,err:=context.Request.Cookie("uid")
	if err!=nil{
		fmt.Println("获取cookie失败")
		return
	}
	uid:=cookie.Value
	re:=Model.QueryUid(uid)
	context.JSON(200,re)
}

func Help(context *gin.Context)  {
	context.JSON(200,gin.H{
		"注册账号":"POST /register    用户名：uid   密码：password",
		"登录账号":"GET /login        用户名：uid   密码：password",
		"给自己充值":"PUT /Recharge   金额：money   备注：remarks",
		"转账":"PUT	/transfer       金额：money   对象：toUid	备注：remarks",
		"根据备注查询转账记录":"GET /queryremarks    查询关键字:remarks",
		"查询自己转账记录":"GET /queryUid",
	})
}