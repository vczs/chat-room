package process

import (
	"client/login"
	"client/utils"
	"fmt"
	"os"
)

func Login() {

	//从终端获取用户输入的登录信息
	admin, password, inputSignInfoErr := inputSignInfo()
	if inputSignInfoErr != nil {
		fmt.Printf("inputSignInfoErr：%v\n", inputSignInfoErr)
		return
	}

	//连接到服务器 拿到conn链接
	connectServerSut := &utils.ConnectServerSut{}
	connServerErr := connectServerSut.ConnectServer()
	if connServerErr != nil {
		fmt.Printf("login connServerErr：%v\n", connServerErr)
		return
	}

	//将用户的登录信息 生成 登录信息json
	jsonMakeSut := &login.JsonMakeSut{
		Admin:    admin,
		Password: password,
	}
	//生成的登录信息json储存在jsonMakeSut.MesJson中
	JsonMakeErr := jsonMakeSut.JsonMake()
	if JsonMakeErr != nil {
		fmt.Printf("JsonMakeErr:%v\n", JsonMakeErr)
		return
	}

	//拿到conn链接后  将登录信息json发送给对方
	writePkgSut := &utils.WritePkgSut{
		Conn:    connectServerSut.Conn,
		JsonCon: jsonMakeSut.MesJson,
	}
	WritePkgErr := writePkgSut.WritePkg()
	if WritePkgErr != nil {
		fmt.Printf("WritePkgErr：%v\n", WritePkgErr)
		return
	}

	//发送登录信息给对方后 接收登录结果信息包
	readPkgSut := &utils.ReadPkgSut{
		Conn: connectServerSut.Conn,
	}
	readPkgErr := readPkgSut.ReadPkg()
	if readPkgErr != nil {
		fmt.Printf("readPkgErr:%v\n", readPkgErr)
		return
	}

	//接收到登录结果信息包 开始解析结果信息包json 判断是否登录成功
	jsonAysSut := &login.JsonAysSut{
		MesByte: readPkgSut.MesByte,
		Conn:    connectServerSut.Conn,
		Admin:   admin,
	}
	jsonAysErr := jsonAysSut.JsonAys()
	if jsonAysErr != nil {
		fmt.Printf("错误%v\n", jsonAysErr)
		return
	}
	/*******如果以上操作都未发生错误则代表登录成功 下面执行登陆成功后的操作*******/

	//开启协程保持与服务端链接 即使接收服务端的消息
	go func() {
		keepLoginSut := &login.KeepLoginSut{Conn: connectServerSut.Conn}
		keepLoginErr := keepLoginSut.KeepLogin()
		if keepLoginErr != nil {
			fmt.Printf("keepLoginErr:%v\n", keepLoginErr)
			func() {
				connCloseErr := connectServerSut.Conn.Close()
				if connCloseErr != nil {
					fmt.Printf("login connCloseErr：%v\n", connCloseErr)
					return
				}
				fmt.Println("login connClose suc...")
			}()
			os.Exit(1)
			return
		}
	}()
	//展示登录成功的菜单
	login.LoginMenu(jsonAysSut.Admin, jsonAysSut.AdminName)
	fmt.Println("程序已退出！！！")
}

func inputSignInfo() (admin int, password string, err error) {

	fmt.Print("请输入帐号：")
	_, err = fmt.Scanln(&admin)
	if err != nil {
		fmt.Printf("inputSignInfo:%v\n", err)
		return
	}

	fmt.Print("请输入密码：")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Printf("inputSignInfo:%v\n", err)
		return
	}

	return
}
