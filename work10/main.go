package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	redisInit "github.com/my/repo/redis"
)

func main() {
	c := redisInit.NewClient(redisInit.NewFuncPoolExt(func(pool *redisInit.OptionPool) {
		pool = &redisInit.OptionPool{}
	}))
	err := c.HSet("hash", "123","abc")
	if err!=nil {
		err.Error()
	}
	value, err := redis.String(c.HGet("hash","123"))
	if err!=nil {
		err.Error()
	}

	fmt.Println(value)
}
