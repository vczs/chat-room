package login

import (
	"client/model"
	"client/utils"
	"encoding/json"
	"fmt"
)

func SendGroupMes(content string) error {

	var semMes model.SmsMes
	semMes.User.Admin = curUser.User.Admin
	semMes.User.AdminName = curUser.User.AdminName
	semMes.User.State = curUser.User.State
	semMes.Content = content
	semMesJson, semMesJsonErr := json.Marshal(semMes)
	if semMesJsonErr != nil {
		fmt.Printf("semMesJsonErr%v:\n", semMesJsonErr)
		return semMesJsonErr
	}

	var mes model.Message
	mes.Type = model.SmsMesType
	mes.Data = string(semMesJson)
	mesJson, mesJsonErr := json.Marshal(mes)
	if mesJsonErr != nil {
		fmt.Printf("mesJsonErr%v:\n", mesJsonErr)
		return mesJsonErr
	}

	writePkgSut := &utils.WritePkgSut{
		Conn:    curUser.Conn,
		JsonCon: mesJson,
	}
	mesJsonWriteErr := writePkgSut.WritePkg()
	if mesJsonWriteErr != nil {
		fmt.Printf("mesJsonWriteErr%v:\n", mesJsonWriteErr)
		return mesJsonWriteErr
	}

	return nil
}
