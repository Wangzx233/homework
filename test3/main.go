package main

import (
	"test3/Model"
	"test3/Router"
)

func main() {
	Model.MysqlInit()
	Router.Entrance()
}
