package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

var m map[string]bool
var mu sync.Mutex
func main() {
	m = make(map[string]bool)

	quit := make(chan bool)
	urls := make(chan string)
	errURLs := make(chan string, 40)

	go download(urls, errURLs, quit)
	go getURL("http://jwzx.cqupt.edu.cn/", urls)
	//go dealErrorURL(urls, errURLs)
	<-quit
	fmt.Println(m)
}

func download(urls, errURLs chan string, quit chan bool) {

	for {
		select {
		case url := <-urls:
			//fmt.Println("download拿到了")
			client := &http.Client{
				Transport: &http.Transport{
					Dial: func(netw, addr string) (net.Conn, error) {
						conn, err := net.DialTimeout(netw, addr, time.Second*2)    //设置建立连接超时
						if err != nil {
							return nil, err
						}
						conn.SetDeadline(time.Now().Add(time.Second * 2))    //设置发送接受数据超时
						return conn, nil
					},
					ResponseHeaderTimeout: time.Second * 2,
				},
			}
			request, err := http.NewRequest("GET", url, nil)
			request.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
			do, err := client.Do(request)
			if err != nil {
				//fmt.Println("url出错了")
				//url = "http://jwzx.cqupt.edu.cn/" + url
				//errURLs <- url
				log.Println(err)
				break
			}
			//函数结束后关闭相关链接
			defer do.Body.Close()

			body, err := ioutil.ReadAll(do.Body)
			if err != nil {
				fmt.Println(1)
				log.Println(err)

			}
			//fmt.Println(string(body))
			//.{0,20}[0-9]{3,20}.{0,5}人
			compile,err:= regexp.Compile(".{0,20}男生.{0,10}")//.{0,20}男生.{0,10}
			if err != nil {
				fmt.Println(2)
				log.Println(err)
				return
			}
			allString := compile.FindAllString(string(body), -1)
			//allString := compile.FindAllString("学生总数：200", 99)
			if allString!=nil {
				fmt.Println(allString)
			}

			//links := collectlinks.All(do.Body)
			//for _, link := range links {
			//	fmt.Println("parse url", link)
			//}
		}

	}
	fmt.Println("download_end")
	quit <- true
}

func getURL(url string, urls chan string) {

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))

	})

	c.OnRequest(func(r *colly.Request) {
		urls<-r.URL.String()
		//fmt.Println(r.URL.String())
	})

	//c.Visit("http://jwzx.cqupt.edu.cn/")
	c.Visit("http://jwzx.cqupt.edu.cn/")


	////判重
	////mu.Lock()
	//if m[url] {
	//	//mu.Unlock()
	//	return
	//} else {
	//
	//	m[url] = true
	//	//mu.Unlock()
	//}
	//fmt.Println(len(m))
	//if url=="http://jwzx.cqupt.edu.cn/fileShowContent.php?id=7708" {
	//	fmt.Println("11111")
	//}
	//urls <- url
	////get网页
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		Dial: func(netw, addr string) (net.Conn, error) {
	//			conn, err := net.DialTimeout(netw, addr, time.Second*2)    //设置建立连接超时
	//			if err != nil {
	//				return nil, err
	//			}
	//			conn.SetDeadline(time.Now().Add(time.Second * 2))    //设置发送接受数据超时
	//			return conn, nil
	//		},
	//		ResponseHeaderTimeout: time.Second * 2,
	//	},
	//}
	//request, err := http.NewRequest("GET", url, nil)
	//request.Header.Set("Content-Type","application/x-www-form-urlencoded")
	//do, err := client.Do(request)
	//if err != nil {
	//	//url = "http://jwzx.cqupt.edu.cn/" + url
	//	//getURL(url, urls)
	//	log.Println(err)
	//	return
	//}
	////函数结束后关闭相关链接
	//defer do.Body.Close()
	//
	////body, err := ioutil.ReadAll(do.Body)
	////if err != nil {
	////	fmt.Println("read error", err)
	////	return
	////}
	////fmt.Println(string(body))
	//links := collectlinks.All(do.Body)


	//for _, link := range links {
	//	absolutely:=urlJoin(link,"http://jwzx.cqupt.edu.cn/" )
	//
	//	if absolutely != "" {
	//		urls <- absolutely
	//		//fmt.Println("传入urls了")
	//		getURL(absolutely, urls)
	//		//fmt.Println(links)
	//	}
	//}
	//fmt.Println("END")
}

func dealErrorURL(urls, errURLs chan string) {
	for eul := range errURLs {
		//判重
		mu.Lock()
		if m[eul] {
			mu.Unlock()
		} else {

			m[eul] = true
			mu.Unlock()
			urls<-eul
		}
		//fmt.Println("解决了错误url")
	}
}

func urlJoin(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	return baseUrl.ResolveReference(uri).String()
}
