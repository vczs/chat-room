package login

import (
	"client/model"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type JsonAysSut struct {
	MesByte   []byte
	Conn      net.Conn
	Admin     int
	AdminName string
}

//处理登陆结果信息包判断是否登录成功
func (jsonAysSut *JsonAysSut) JsonAys() error {

	var mes model.Message
	mesUnmarshalErr := json.Unmarshal(jsonAysSut.MesByte, &mes)
	if mesUnmarshalErr != nil {
		fmt.Printf("mesUnmarshalErr:%v\n", mesUnmarshalErr)
		return mesUnmarshalErr
	}

	if mes.Type != model.LoginResMesType { //对方返回的mes信息包不是登陆结果类型的信息包
		return errors.New("非法登录结果信息包！！！")
	}

	var loginResMes model.LoginResMes
	loginResUnmarshalErr := json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResUnmarshalErr != nil {
		fmt.Printf("loginResUnmarshalErr:%v\n", loginResUnmarshalErr)
		return loginResUnmarshalErr
	}

	if loginResMes.Code == 200 {
		jsonAysSut.AdminName = loginResMes.Error
		fmt.Println("登陆成功！！！")

		curUser = CurUser{
			Conn: jsonAysSut.Conn,
			User: model.User{
				Admin:     jsonAysSut.Admin,
				AdminName: jsonAysSut.AdminName,
				State:     model.UserOnOnline,
			},
		}

		for k, v := range loginResMes.OnlineUser {
			user := &model.User{
				Admin:     k,
				AdminName: v,
				State:     model.UserOnOnline,
			}
			onlineUser[k] = user
		}
		return nil
	} else {
		loginErr := fmt.Errorf("%v:%v", loginResMes.Code, loginResMes.Error)
		return loginErr
	}
}
