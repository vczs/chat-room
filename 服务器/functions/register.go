package functions

import (
	"client/model"
	"client/user"
	"encoding/json"
	"fmt"
)

type Register struct {
	Mes            *model.Message
	RegisterResPkg []byte
}

var (
	codeRegister int
	errRegister  string
)

func (register *Register) Register() error {

	//取出message.Data反序列化为registerMes
	var registerMes model.RegisterMes
	registerMesUnmarshalErr := json.Unmarshal([]byte(register.Mes.Data), &registerMes)
	if registerMesUnmarshalErr != nil {
		fmt.Printf("registerMesUnmarshalErr:%v", registerMesUnmarshalErr)
		return registerMesUnmarshalErr
	}
	fmt.Println(register.Mes.Data)

	addUserErr := user.MyUserDao.AddUser(&registerMes.User)
	if addUserErr != nil {
		if addUserErr == user.ERROR_USER_EXISTS {
			codeRegister = 501 //状态码501 用户已存在
			errRegister = fmt.Sprintf("用户{%v}已存在！\n", registerMes.User.Admin)
		} else {
			codeRegister = 404 //状态码404 表示服务器发生错误
			errRegister = "服务器发生错误！"
		}
	} else {
		//注册成功  就给registerMes的字段赋值注册成功的内容
		codeRegister = 201 //状态码200 错误信息为空表示注册成功
		errRegister = registerMes.User.AdminName
		fmt.Printf("用户{%v}注册成功！\n", registerMes.User.AdminName)
	}

	//生成登陆结果信息json
	makeLoginResJson := register.makeRegisterResPkg()
	if makeLoginResJson != nil {
		return makeLoginResJson
	}

	return nil
}

//制作登录信息包
func (register *Register) makeRegisterResPkg() error {

	//创建登录结果信息包详细信息
	var registerMes model.RegisterResMes
	registerMes.Code = codeRegister
	registerMes.Error = errRegister
	//序列化logResMes
	registerMesJson, registerMesJsonErr := json.Marshal(registerMes)
	if registerMesJsonErr != nil {
		fmt.Printf("registerMesJsonErr:%v", registerMesJsonErr)
		return registerMesJsonErr
	}

	var mesRes model.Message               //创建登陆结果信息包mesRes
	mesRes.Type = model.RegisterResMesType //mesRes的消息类型是登陆结果信息
	mesRes.Data = string(registerMesJson)  //将序列化后的logResMes转为字符串类型赋值给mesRes.Data
	//系列化mesRes
	mesResJson, mesResJsonErr := json.Marshal(mesRes)
	if mesResJsonErr != nil {
		fmt.Printf("mesResJsonErr:%v", mesResJsonErr)
		return mesResJsonErr
	}

	register.RegisterResPkg = mesResJson //将登录结果信息的json赋值给login.LoginResPkg
	return nil
}
