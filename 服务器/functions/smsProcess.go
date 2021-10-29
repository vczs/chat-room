package functions

import (
	"client/model"
	"client/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (smsProcess *SmsProcess) SendGroupMes(mes *model.Message) error {
	var smsMes model.SmsMes
	smsMesUmErr := json.Unmarshal([]byte(mes.Data), &smsMes)
	if smsMesUmErr != nil {
		fmt.Printf("smsMesUmErr:%v\n", smsMesUmErr)
		return smsMesUmErr
	}

	mesJson, mesJsonErr := json.Marshal(mes)
	if mesJsonErr != nil {
		fmt.Printf("mesJsonErr:%v\n", mesJsonErr)
		return mesJsonErr
	}

	for k, v := range UserMgrPointer.onlineUsers {
		if k == smsMes.User.Admin {
			continue
		}
		sendToUserErr := smsProcess.sendToUser(mesJson, v.Conn)
		if sendToUserErr != nil {
			fmt.Printf("mesJsonErr:%v\n", sendToUserErr)
			return sendToUserErr
		}
	}

	return nil
}

func (smsProcess *SmsProcess) sendToUser(mes []byte, conn net.Conn) error {
	serverWritePkgSut := &utils.ServerWritePkgSut{
		Conn: conn,
		Mes:  mes,
	}
	sendToUserPkg := serverWritePkgSut.ServerWritePkg()
	if sendToUserPkg != nil {
		fmt.Printf("sendToUserPkg:%v\n", sendToUserPkg)
		return sendToUserPkg
	}
	return nil
}
