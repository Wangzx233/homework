package main

import "fmt"

func main() {
	var v int

	receiver(v)
}

func receiver(v interface{})  {
	switch v.(type) {
	case string:
		fmt.Println("这是一个string类型")
	case int:
		fmt.Println("这是一个int类型")
	case float32:
		fmt.Println("这是一个float类型")
	case bool:
		fmt.Println("这是一个bool类型")
	case byte:
		fmt.Println("这是一个byte类型")
	}
}
