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
