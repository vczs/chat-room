package login

import (
	"client/model"
	"encoding/json"
	"fmt"
)

type JsonMakeSut struct {
	Admin    int
	Password string
	MesJson  []byte
}

func (jsonMake *JsonMakeSut) JsonMake() error {

	/*******开始登录信息的json生成*******/
	//创建一个model.LoginMes结构体类型的变量 将Admin和Password分别赋值给其字段
	var logMes model.LoginMes
	logMes.Admin = jsonMake.Admin
	logMes.Password = jsonMake.Password
	//赋值后序列化logMes变量
	logMesJson, logMesJsonErr := json.Marshal(logMes)
	if logMesJsonErr != nil {
		fmt.Printf("logMesJsonErr：%v\n", logMesJsonErr)
		return logMesJsonErr
	}

	//创建一个model.Message结构体类型的变量mes 将序列化后的logMes赋值给mes的Data字段
	var mes model.Message
	mes.Type = model.LoginMesType
	mes.Data = string(logMesJson)
	//将mes序列化
	mesJson, mesJsonErr := json.Marshal(mes)
	if mesJsonErr != nil {
		fmt.Printf("mesJsonErr：%v\n", mesJsonErr)
		return mesJsonErr
	}
	jsonMake.MesJson = mesJson
	/*******以上就完成了登录信息的json生成*******/

	return nil
}
