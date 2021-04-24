package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	r := gin.Default()
	r.GET("ks", func(context *gin.Context) {
		xh := context.Query("xh")
		info := getInfo("http://jwzx.cqupt.edu.cn/ksap/showKsap.php?type=stu&id=" + xh)
		context.JSON(200,info)
	})

	r.Run()

}

func getInfo(url string) []string {
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

	compile:= regexp.MustCompile(`<td>(.{0,20})</td>`)

	var s []string
	for _, match := range compile.FindAllString(string(body), -1) {
		s=append(s,match)
		fmt.Println(match)
	}
	return s
}