package main

import (
	"client/process"
	"client/user"
	"client/utils"
	"fmt"
	"net"
)

func main() {

	utils.InitPool("localhost:6379", 10, 0)      //创建并初始化线程池pool
	user.MyUserDao = user.NewUserDao(utils.Pool) //实例化操作redis数据库的对象MyUserDao 用它直接操作UserDao的所有方法

	//开启监听
	fmt.Println("The 服务器 starts listening on port 9003 ...")
	listen, listenErr := net.Listen("tcp", "0.0.0.0:9003")
	if listenErr != nil {
		fmt.Printf("listenErr：%v\n", listenErr)
		return
	}

	//函数退出后关闭listen
	defer func() {
		listenCloseErr := listen.Close() //关闭listen
		if listenCloseErr != nil {
			fmt.Printf("listenCloseErr=%v\n", listenCloseErr)
			return
		}
		fmt.Println("listenClose suc...")
	}()

	//监听开启成功 等待客户端连接
	for {
		conn, acceptErr := listen.Accept()
		if acceptErr != nil {
			fmt.Printf("acceptErr=%v\n", acceptErr)
		} else {
			ip := conn.RemoteAddr().String()
			fmt.Printf("New: {%v} connect ...\n", ip)
			//有客户端接入 开启协程处理
			processSut := &process.ProcessSut{
				Conn: conn,
				Ip:   ip,
			}
			go processSut.Process()
		}
	}
}
