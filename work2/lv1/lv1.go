package lv1

import "fmt"
func main() {


	var video1 Videoinfo
	video1=Videoinfo{
		 Author: Author{
			Name:      "天天卡牌",
			VIP:       true,
			Icon:      "123.png",
			Signature: "微信公众号:ttkapai 商务合作邮箱:ttkapai_hz@1",
			Focus:     225,
		},
		Head: Head_title{
			Title:     "炉石传说：【天天素材库】 第218期",
			Bullet:    2385,
			Plays:     68,
			Date_time: "2020-10-26 11:28:39",
		},
	}
	fmt.Println(video1)


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
