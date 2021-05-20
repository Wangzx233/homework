package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	//r := gin.Default()
	//r.GET("keb", func(context *gin.Context) {
	//	xh := context.Query("xh")
	//	info := getKeb("http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + xh)
	//
	//	context.JSON(200, info)
	//})
	//
	//r.Run()
	keb := getKeb("http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + "2020215058")
	getKec(keb)
}

//获取课表
func getKeb(url string) string {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	do, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	defer do.Body.Close()

	body, err := ioutil.ReadAll(do.Body)
	if err != nil {
		fmt.Println(1)
		log.Println(err)
	}

	compile := regexp.MustCompile(`<div\s*id='stuPanel'>\s*((?:\w|\W)*)</tbody></table>\s*</div>`)

	keb := compile.FindString(string(body))

	return keb
}

//从课表中提取课程
func getKec(keb string) (string) {
	compile := regexp.MustCompile(`(?mU)<tr style='text-align:center'><td style='font-weight:bold;'>\s*((?:\w|\W)*)</td></tr>`)
	kec := compile.FindAllStringSubmatch(keb,-1)
	//fmt.Println(kec[0][1])
	time:=kec[0][1][0:8]
	getKcm(kec[0][1])
	fmt.Println(time)
	return time
}

func getKcm(kec string) ([]string) {
	var s []string
	compile := regexp.MustCompile(`(?mU)<td>\s*((?:\w|\W))*</td>`)
	submatch := compile.FindAllStringSubmatch(kec,-1)
	for _,k:=range submatch{
		s=append(s,k[1])
	}
	fmt.Println(s)
	return s
}