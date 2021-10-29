package login

import (
	"fmt"
)

var (
	loop bool
)

func LoginMenu(admin int , adminName string){
	loop = true
	fmt.Printf("欢迎您，%v(%v)！\n",adminName,admin)
	for loop {
		fmt.Printf("当前帐号:[%v(%v)]\n",adminName,admin)
		fmt.Println("*******登录成功*******")
		fmt.Println("*     1.用户列表     *")
		fmt.Println("*     2.发送消息     *")
		fmt.Println("*     3.信息列表     *")
		fmt.Println("*     4.退出登录     *")
		fmt.Println("*********************")
		loginSelect()
	}
}

func loginSelect() {
	fmt.Println("请选择(1-4):")
	option := 4 //默认退出程序
	_ , loginSelect := fmt.Scanln(&option)
	if loginSelect != nil {
		fmt.Printf("loginSelect:%v\n", loginSelect)
		return
	}
	switch option {
	case 1 :
		showOnOnlineUser()
	case 2 :
		sendMes()
	case 3 :
		fmt.Println("信息列表")
	case 4 :
		fmt.Println("退出登录")
		loop = false
	default:
		fmt.Println("输入有误，请重新输入。")
	}
}

func sendMes(){
	var content string
	fmt.Print("请输入发送内容：")
	_ , contentScanErr := fmt.Scanln(&content)
	if contentScanErr != nil {
		fmt.Printf("contentScanErr:%v\n", contentScanErr)
		return
	}
	SendGroupMesErr := SendGroupMes(content)
	if SendGroupMesErr != nil {
		fmt.Printf("SendGroupMesErr:%v\n", SendGroupMesErr)
		return
	}
}