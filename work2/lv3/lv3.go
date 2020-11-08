package main

import "fmt"

func main() {
	var n1,t1 string
	v1:=new(Videoinfo)
	fmt.Scanf("%s",&t1)
	fmt.Scanf("%s",&n1)
	v1=release(n1,t1)
	fmt.Println(v1)

}


func (V Video)up(){			//点赞
	V.Good++
}
func (v Video)coin()  {			//投币
	v.Coins++
}
func (v Video) collect()  {			//收藏
	v.Collection++
}
func (v Video) triple()  {		//三连
	v.Collection++
	v.Coins++
	v.Good++
}
type Head_title struct{			//视频标题部分
	Title string
	Bullet int
	Plays int
	Date_time string
}
type Author struct {
	Name string             //名字
	VIP bool                //是否是高贵的带会员
	Icon string             //头像
	Signature string        //签名
	Focus int               //关注人数
}
type Video struct {				//视频下方部分
	Good int
	Coins int
	Collection int
}
type Videoinfo struct {			//视频详细页
	Head Head_title
	Author Author
	Video Video
}
func release(name string,video_title string)*Videoinfo  {			//接收者函数
	v1:= Videoinfo{
		Head:   Head_title{Title:video_title},
		Author: Author{Name: name},
		Video:  Video{Good: 1},
	}
	return &v1
}