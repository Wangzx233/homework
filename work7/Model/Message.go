package Model

import (
	"fmt"
	"strings"
)

var Mp=make(map[string]Reply)
func ViewMessage(reply_id int,to_uid string,nowUid string) (int,string,string) {
	row,err:=DB.Query("select id,reply_id,content,from_uid,to_uid,power from reply where reply_id=? and to_uid=?",reply_id,to_uid)
	defer row.Close()
	if err!=nil{
		fmt.Println("查询出错")
		return 0,"",""
	}


	var reply Reply
	for row.Next(){
		err=row.Scan(&reply.Id,&reply.Reply_id,&reply.Content,&reply.From_uid,&reply.To_uid,&reply.Power)
		if err!=nil{
			fmt.Println("写入出错")
		}

		if reply.Power!="all"{
			if strings.Index(reply.Power,"&"+nowUid+"&")==-1{
				reply.Content="对方已设置为私密留言，无权限访问"
			}
		}

		Mp[reply.From_uid+"回复"+to_uid]=reply
		if len(reply.Content)!=0 {
			fmt.Println(reply.Id)
			ViewMessage(reply.Id,reply.From_uid,nowUid)
		}
	}

	return -1, to_uid,""
}

func SendMessage(reply_id int,content,from_uid,to_uid,power string)  {

	res,err:=DB.Exec("insert reply (reply_id,content,from_uid,to_uid,power) values (?,?,?,?,?)",reply_id,content,from_uid,to_uid,power)
	if err!=nil{
		fmt.Printf("回复信息出错: %v",err)
	}
	id,_:=res.LastInsertId()
	_,err=DB.Exec("insert likes (reply_id,likes) values (?,0)",id)
	if err!=nil{
		fmt.Printf("创建点赞出错: %v",err)
	}
}

func DeleteMessage(id int,uid string) bool {
	row:=DB.QueryRow("select id,from_uid from reply where id=?",id)
	var muid string
	err:=row.Scan(&id,&muid)
	if err!=nil{
		fmt.Println("插入数据出错",err)
	}
	if uid!=muid{
		return false
	}
	DeleteCycle(id)
	DB.Exec("delete from reply where id=?",id)
	return true
}
func DeleteCycle(id int)  {
	fmt.Println(id)
	row:=DB.QueryRow("select id from reply where reply_id=?",id)
	_,err:=DB.Exec("delete from reply where reply_id=?",id)
	if err!=nil{
		fmt.Println("删除信息出错")
		return
	}
	var r int
	er:=row.Scan(&r)
	if er!=nil{
		fmt.Println("获取最后一条id失败")
		return
	}
	DeleteCycle(r)
}


func Like(reply_id int)  {
	_, err := DB.Exec("update likes set likes = likes+1 where reply_id = ?",reply_id)
	if err!=nil{
		fmt.Println("点赞插入出错")
	}
}