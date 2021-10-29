package login

import (
	"client/model"
	"client/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type KeepLoginSut struct {
	Conn net.Conn
}

func (keepLoginSut *KeepLoginSut) KeepLogin() error {

	//循环等待接收对方发送的消息 当对方有消息发送过来及时接收
	readPkgSut := utils.ReadPkgSut{Conn: keepLoginSut.Conn}
	for {
		readNewsErr := readPkgSut.ReadPkg()
		if readNewsErr != nil {
			fmt.Printf("readNewsErr：%v\n", readNewsErr)
			return readNewsErr
		}
		var mes model.Message
		mesUnMarshalErr := json.Unmarshal(readPkgSut.MesByte, &mes)
		if mesUnMarshalErr != nil {
			fmt.Printf("mesUnMarshalErr：%v\n", mesUnMarshalErr)
			return mesUnMarshalErr
		}

		switch mes.Type {
		case model.NotifyUserStateMesTyp:
			updateUserState(&mes)

		case model.SmsMesType:
			smsMgr := &SmsMgr{}
			showServerMesErr := smsMgr.ShowServerMes(&mes)
			if showServerMesErr != nil {
				fmt.Printf("showServerMesErr：%v\n", showServerMesErr)
				return showServerMesErr
			}

		default:
			return errors.New("服务端返回非法消息类型！")
		}

	}

}
