package main

import (
	"gay/gay"
)

func main() {
	r := gay.New()
	r.GET("/123", func(ctx *gay.Context) {
		ctx.JSON(200,gay.H{
			"asd":"asd",
		})
	})
	r.Run()
}
