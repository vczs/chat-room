package login

import (
	"client/model"
	"encoding/json"
	"fmt"
)

var onlineUser = make(map[int]*model.User, 10)

func updateUserState(mes *model.Message) {

	var notifyUserStateMes model.NotifyUserStateMes
	notifyUserStateMesUnMarshalErr := json.Unmarshal([]byte(mes.Data), &notifyUserStateMes)
	if notifyUserStateMesUnMarshalErr != nil {
		fmt.Printf("notifyUserStateMesUnMarshalErr：%v\n", notifyUserStateMesUnMarshalErr)
		return
	}
	if notifyUserStateMes.Flag == model.Online {
		fmt.Printf("[%v(%v)]上线了！\n", notifyUserStateMes.OnlineUser[notifyUserStateMes.Admin], notifyUserStateMes.Admin)
	}
	if notifyUserStateMes.Flag == model.Offline {
		fmt.Printf("[%v(%v)]已离线！\n", onlineUser[notifyUserStateMes.Admin].AdminName, notifyUserStateMes.Admin)
	}

	onlineUser = make(map[int]*model.User, 10) //清空onlineUser
	for k, v := range notifyUserStateMes.OnlineUser {
		user := &model.User{
			Admin:     k,
			AdminName: v,
			State:     model.UserOnOnline,
		}
		onlineUser[k] = user
	}
}

func showOnOnlineUser() {
	fmt.Printf("当前在线用户：")
	for k, v := range onlineUser {
		fmt.Printf("[%v(%v)]   ", k, v.AdminName)
	}
	fmt.Println()
}
