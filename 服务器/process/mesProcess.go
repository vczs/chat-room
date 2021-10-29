package process

import (
	"client/functions"
	"client/model"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type MesProcessSut struct {
	Ip       string
	Conn     net.Conn
	InfoRead []byte
	ResPkg   []byte
	Flag     int //0登录 1注册 2发消息
}

func (mesProcessSut *MesProcessSut) MesProcess() error {
	var mes *model.Message
	//反序列化读取的信息
	mesUnmarshalErr := json.Unmarshal(mesProcessSut.InfoRead, &mes)
	if mesUnmarshalErr != nil {
		fmt.Printf("mesUnmarshalErr:%v", mesUnmarshalErr)
		return mesUnmarshalErr
	}

	switch mes.Type {

	case model.LoginMesType: //登录类型信息包
		//登录请求 执行登录功能
		mesProcessSut.Flag = 0
		loginSut := &functions.LoginSut{Mes: mes, Conn: mesProcessSut.Conn, Ip: mesProcessSut.Ip}
		loginErr := loginSut.Login()
		if loginErr != nil {
			return loginErr
		} //如果执行登录功能发生错误 则返回错误信息
		mesProcessSut.ResPkg = loginSut.LoginResPkg //登录功能执行成功 将登录结果信息包赋值给mesProcessSut.ResPkg

	case model.RegisterMesType:
		//注册请求 执行注册功能
		mesProcessSut.Flag = 1
		register := &functions.Register{Mes: mes}
		registerErr := register.Register()
		if registerErr != nil {
			return registerErr
		} //如果注册功能发生错误 则返回错误信息
		mesProcessSut.ResPkg = register.RegisterResPkg //注册功能执行成功 将注册结果信息包赋值给mesProcessSut.ResPkg

	case model.SmsMesType:
		mesProcessSut.Flag = 2
		smsProcess := &functions.SmsProcess{}
		sendSmsMesErr := smsProcess.SendGroupMes(mes)
		if sendSmsMesErr != nil {
			return sendSmsMesErr
		}

	default:
		mesProcessErr := errors.New(fmt.Sprintln("对方信息包类型无法识别！！！"))
		return mesProcessErr
	}

	return nil
}
