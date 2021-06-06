package control

import (
	"examination/room/socket"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCheek(c *gin.Context)  {
	roomID := c.PostForm("room_id")

	//i:=false
	//for range {
	//	i=true
	//}
	fmt.Println(roomID)
	if socket.Mng.Chans[roomID]==nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "创建成功"})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    10023,
			"message": "房间已存在"})
	}
}

func InterCheek(c *gin.Context) {
	roomID := c.PostForm("room_id")
	i:=0
	for range socket.Mng.Chans[roomID]{
		i++
	}
	if i==0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    10024,
			"message": "房间不存在"})
	}else if i>=2 {
		c.JSON(http.StatusOK, gin.H{
			"code":    10025,
			"message": "房间已满"})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "加入成功"})
	}
}