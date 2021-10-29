package process

import (
	"client/register"
	"client/utils"
	"fmt"
)

func Register() {

	//从终端获取用户输入的注册信息
	admin, password, adminName, inputRegisterInfoErr := inputRegisterInfo()
	if inputRegisterInfoErr != nil {
		fmt.Printf("inputRegisterInfoErr：%v\n", inputRegisterInfoErr)
		return
	}

	//连接到服务器 拿到conn链接
	connectServerSut := &utils.ConnectServerSut{}
	connServerErr := connectServerSut.ConnectServer()
	if connServerErr != nil {
		fmt.Printf("register connServerErr：%v\n", connServerErr)
		return
	}
	//函数退出conn链接
	defer func() {
		connCloseErr := connectServerSut.Conn.Close()
		if connCloseErr != nil {
			fmt.Printf("register connCloseErr：%v\n", connCloseErr)
			return
		}
		fmt.Println("register connClose sec...")
	}()

	//将用户的注册信息 生成 注册信息json
	jsonMakeRegisterSut := &register.JsonMakeRegisterSut{
		Admin:     admin,
		Password:  password,
		AdminName: adminName,
	}
	//生成的注册信息json储存在jsonMakeRegisterSut.MesJson中
	JsonMakeRegisterErr := jsonMakeRegisterSut.JsonMakeRegister()
	if JsonMakeRegisterErr != nil {
		fmt.Printf("JsonMakeRegisterErr:%v\n", JsonMakeRegisterErr)
		return
	}

	//拿到conn链接后  将注册信息json 发送给对方
	writePkgSut := &utils.WritePkgSut{
		Conn:    connectServerSut.Conn,
		JsonCon: jsonMakeRegisterSut.MesJson,
	}
	WritePkgErr := writePkgSut.WritePkg()
	if WritePkgErr != nil {
		fmt.Printf("WritePkgErr：%v\n", WritePkgErr)
		return
	}

	//发送注册信息给对方后 接收注册结果信息包
	readPkgSut := &utils.ReadPkgSut{
		Conn: connectServerSut.Conn,
	}
	readPkgErr := readPkgSut.ReadPkg()
	if readPkgErr != nil {
		fmt.Printf("readPkgErr:%v\n", readPkgErr)
		return
	}

	//接收到登录结果信息包 开始解析结果信息包json 判断是否登录成功
	jsonAysRegisterSut := &register.JsonAysRegisterSut{
		MesByte: readPkgSut.MesByte,
	}
	jsonAysErr := jsonAysRegisterSut.JsonAys()
	if jsonAysErr != nil {
		fmt.Printf("错误%v\n", jsonAysErr)
		return
	}
	/*******如果以上操作都未发生错误则代表注册成功*******/
}

func inputRegisterInfo() (admin int, password string, adminName string, err error) {

	fmt.Print("请输入帐号：")
	_, err = fmt.Scanln(&admin)
	if err != nil {
		fmt.Printf("inputRegisterInfo:%v\n", err)
		return
	}

	fmt.Print("请输入密码：")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Printf("inputRegisterInfo:%v\n", err)
		return
	}

	fmt.Print("请输入昵称：")
	_, err = fmt.Scanln(&adminName)
	if err != nil {
		fmt.Printf("inputRegisterInfo:%v\n", err)
		return
	}

	return
}
