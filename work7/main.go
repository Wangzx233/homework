package main

import (
	"1/Model"
	"1/Router"
)

func main() {
	Model.MysqlInit()
	Router.Entrance()
}
