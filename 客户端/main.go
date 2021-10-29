package main

import (
	"client/process"
	"fmt"
)

var loop bool

func main() {
	homeMenu()
}

func homeMenu() {
	loop = true
	for loop {
		fmt.Println("***用户通讯系统***")
		fmt.Println("*     1.登录     *")
		fmt.Println("*     2.注册     *")
		fmt.Println("*     3.退出     *")
		fmt.Println("*****************")
		sOption()
	}
}

func sOption() {
	fmt.Println("请选择(1-3):")
	option := 3 //默认退出程序
	_, sOptionErr := fmt.Scanln(&option)
	if sOptionErr != nil {
		fmt.Printf("sOptionErr:%v\n", sOptionErr)
		return
	}
	switch option {
	case 1:
		process.Login()
		loop = false
	case 2:
		process.Register()
		loop = false
	case 3:
		fmt.Println("程序已退出。。。")
		loop = false
	default:
		fmt.Println("输入有误，请重新输入。")
	}
}
