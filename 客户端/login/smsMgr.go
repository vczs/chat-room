package login

import (
	"client/model"
	"encoding/json"
	"fmt"
)

type SmsMgr struct{}

func (smsMgr *SmsMgr) ShowServerMes(mes *model.Message) error {
	var smsMes model.SmsMes
	smsMesUmErr := json.Unmarshal([]byte(mes.Data), &smsMes)
	if smsMesUmErr != nil {
		fmt.Printf("smsMesUmErr:%v", smsMesUmErr)
		return smsMesUmErr
	}
	fmt.Printf("%v(%v)：%v\n", smsMes.User.AdminName, smsMes.User.Admin, smsMes.Content)
	return nil
}
