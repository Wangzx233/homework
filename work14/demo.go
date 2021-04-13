package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("t", func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		fmt.Println(header)
		s := strings.SplitN(header, " ", 2)
		//s := header[6:]
		decodeString, err := base64.StdEncoding.DecodeString(s[0])
		fmt.Println(string(decodeString), err)
		c.Header("WWW-Authenticate", `Basic realm=”localhost”`)
		c.String(401, "")
	})
	r.Run()
}
