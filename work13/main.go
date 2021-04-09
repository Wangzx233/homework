package main

import (
	Model "demo/model"
	"demo/router"
)

func main() {
	Model.SqlInit()


	router.Router()

}