package register

import (
	"client/model"
	"encoding/json"
	"fmt"
)

type JsonMakeRegisterSut struct {
	Admin     int
	Password  string
	AdminName string
	MesJson   []byte
}

func (jsonMakeRegisterSut *JsonMakeRegisterSut) JsonMakeRegister() error {

	/*******开始登录信息的json生成*******/
	//创建一个model.RegisterMes结构体类型的变量 将Admin和Password和AdminName分别赋值给其字段
	var registerMes model.RegisterMes
	registerMes.User.Admin = jsonMakeRegisterSut.Admin
	registerMes.User.Password = jsonMakeRegisterSut.Password
	registerMes.User.AdminName = jsonMakeRegisterSut.AdminName
	//赋值后序列化registerMes变量
	registerMesJson, registerMesJsonErr := json.Marshal(registerMes)
	if registerMesJsonErr != nil {
		fmt.Printf("registerMesJsonErr：%v\n", registerMesJsonErr)
		return registerMesJsonErr
	}

	//创建一个model.Message结构体类型的变量mes 将序列化后的registerMes赋值给mes的Data字段
	var mes model.Message
	mes.Type = model.RegisterMesType
	mes.Data = string(registerMesJson)
	//将mes序列化
	mesJson, mesJsonErr := json.Marshal(mes)
	if mesJsonErr != nil {
		fmt.Printf("mesJsonErr：%v\n", mesJsonErr)
		return mesJsonErr
	}
	jsonMakeRegisterSut.MesJson = mesJson
	/*******以上就完成了登录信息的json生成*******/

	return nil
}
