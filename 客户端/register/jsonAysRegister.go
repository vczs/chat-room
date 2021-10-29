package register

import (
	"client/model"
	"encoding/json"
	"errors"
	"fmt"
)

type JsonAysRegisterSut struct {
	MesByte []byte
}

//处理登陆结果信息包判断是否登录成功
func (jsonAysRegisterSut *JsonAysRegisterSut) JsonAys() error {

	var mes model.Message
	mesUnmarshalErr := json.Unmarshal(jsonAysRegisterSut.MesByte, &mes)
	if mesUnmarshalErr != nil {
		fmt.Printf("mesUnmarshalErr:%v\n", mesUnmarshalErr)
		return mesUnmarshalErr
	}

	if mes.Type != model.RegisterResMesType { //对方返回的mes信息包不是登陆结果类型的信息包
		return errors.New("非法注册结果信息包！！！")
	}

	var registerResMes model.RegisterResMes
	registerResUnmarshalErr := json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResUnmarshalErr != nil {
		fmt.Printf("registerResUnmarshalErr:%v\n", registerResUnmarshalErr)
		return registerResUnmarshalErr
	}

	if registerResMes.Code == 201 {
		fmt.Printf("%v,恭喜您注册成功！\n", registerResMes.Error)
		return nil
	} else {
		loginErr := fmt.Errorf("%v:%s", registerResMes.Code, registerResMes.Error)
		return loginErr
	}
}
