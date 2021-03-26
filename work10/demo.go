package main

import (
	"github.com/go-redis/redis"
)
var redisDb *redis.Client
func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:               "",
		Password:           "",
		DB:                 0,
	})
	_, err = redisDb.Ping().Result()
	if err != nil{
		return err
	}
	return err
}
