package main

import "fmt"

func main() {
	//
	var zh,mm string
	var zczh,zcmm string="asdf","werf"
	fmt.Scan(&zh)
	if(zh==zczh){
		fmt.Print("账号正确\n")
		fmt.Scan(&mm)
		if(zcmm==mm){
			fmt.Print("密码正确")
		}
	}



}
