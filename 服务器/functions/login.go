package functions

import (
	"client/model"
	"client/user"
	"client/utils"
	"encoding/json"
	"fmt"
	"net"
)

type LoginSut struct {
	Ip          string
	Conn        net.Conn
	AdminName   string
	Mes         *model.Message
	LoginResPkg []byte
}

var (
	codeLogin int
	errLogin  string
)

func (loginSut *LoginSut) Login() error {

	//取出message.Data反序列化为LoginMes
	var loginMes model.LoginMes
	loginMesUnmarshalErr := json.Unmarshal([]byte(loginSut.Mes.Data), &loginMes)
	if loginMesUnmarshalErr != nil {
		fmt.Printf("loginMesUnmarshalErr:%v", loginMesUnmarshalErr)
		return loginMesUnmarshalErr
	}
	getUser, getUserErr := user.MyUserDao.GetUser(loginMes.Admin, loginMes.Password)
	if getUserErr != nil {
		if getUserErr == user.ERROR_USER_PWD_ERROR || getUserErr == user.ERROR_USER_NOTEXISTS {
			//如果账户不存在或密码不正确  就给logResMes的字段赋值登陆失败的内容
			codeLogin = 500 //状态码500 表示登陆失败
			errLogin = getUserErr.Error()
		} else {
			codeLogin = 404 //状态码404 表示服务器发生错误
			errLogin = "服务器发生错误！"
		}
	} else {
		//如果账户和密码正确  就给logResMes的字段赋值登陆成功的内容
		codeLogin = 200 //状态码200 错误信息为空表示登陆成功
		errLogin = getUser.AdminName
		utils.IpBind[loginSut.Ip] = getUser.Admin
		fmt.Printf("用户{%v}已登录！！！\n", getUser.AdminName)

		loginSut.AdminName = getUser.AdminName
		UserMgrPointer.AddOnlineUser(getUser.Admin, loginSut) //将登录成功的用户的mesProcessSut对象添加到userMgr中的onlineUsers map中

		//通知其他用户我上线了
		loginSucNotifyOtherErr := loginSut.NotifyOtherOnline(getUser.Admin, model.Online)
		if loginSucNotifyOtherErr != nil {
			fmt.Printf("loginSucNotifyOtherErr:%v\n", loginSucNotifyOtherErr)
			return loginSucNotifyOtherErr
		}
	}

	//开始生成登陆结果信息json
	makeLoginResJson := loginSut.makeLoginResPkg()
	if makeLoginResJson != nil {
		return makeLoginResJson
	}

	return nil
}

//制作登录信息包
func (loginSut *LoginSut) makeLoginResPkg() error {

	//创建登录结果信息包详细信息
	var logResMes model.LoginResMes
	logResMes.Code = codeLogin
	logResMes.Error = errLogin
	onlineUser := make(map[int]string, 20)
	//遍历onlineUsers所有的在线用户 将其Admin赋值给LoginResMes的OnlineUser字段中
	for k, v := range UserMgrPointer.onlineUsers {
		onlineUser[k] = v.AdminName
	}
	logResMes.OnlineUser = onlineUser
	//序列化logResMes
	logResMesJson, logResMesJsonErr := json.Marshal(logResMes)
	if logResMesJsonErr != nil {
		fmt.Printf("logResMesJsonErr:%v", logResMesJsonErr)
		return logResMesJsonErr
	}

	var mesRes model.Message            //创建登陆结果信息包mesRes
	mesRes.Type = model.LoginResMesType //mesRes的消息类型是登陆结果信息
	mesRes.Data = string(logResMesJson) //将序列化后的logResMes转为字符串类型赋值给mesRes.Data
	//系列化mesRes
	mesResJson, mesResJsonErr := json.Marshal(mesRes)
	if mesResJsonErr != nil {
		fmt.Printf("mesResJsonErr:%v", mesResJsonErr)
		return mesResJsonErr
	}

	loginSut.LoginResPkg = mesResJson //将登录结果信息的json赋值给login.LoginResPkg
	return nil
}

func (loginSut *LoginSut) NotifyOtherOnline(admin int, flag int) error {
	//遍历所有的onlineUser 除了自身全部发送
	for k, v := range UserMgrPointer.onlineUsers {
		if k == admin {
			continue
		}
		notifyMeOnlineErr := v.NotifyMeOnline(admin, flag)
		if notifyMeOnlineErr != nil {
			fmt.Printf("notifyMeOnlineErr:%v\n", notifyMeOnlineErr)
			return notifyMeOnlineErr
		}
	}
	return nil
}

func (loginSut *LoginSut) NotifyMeOnline(admin int, flag int) error {

	var notifyUserStateMes model.NotifyUserStateMes
	notifyUserStateMes.Admin = admin
	notifyUserStateMes.Flag = flag
	onlineUser := make(map[int]string, 20)
	//遍历onlineUsers所有的在线用户 将其Admin赋值给LoginResMes的OnlineUser字段中
	for k, v := range UserMgrPointer.onlineUsers {
		onlineUser[k] = v.AdminName
	}
	notifyUserStateMes.OnlineUser = onlineUser
	notifyUserStateMesJson, notifyUserStateMesJsonErr := json.Marshal(notifyUserStateMes)
	if notifyUserStateMesJsonErr != nil {
		fmt.Printf("notifyUserStateMesJsonErr:%v\n", notifyUserStateMesJsonErr)
		return notifyUserStateMesJsonErr
	}

	var mes model.Message
	mes.Type = model.NotifyUserStateMesTyp
	mes.Data = string(notifyUserStateMesJson)
	mesJson, mesJsonErr := json.Marshal(mes)
	if mesJsonErr != nil {
		fmt.Printf("mesJsonErr:%v\n", mesJsonErr)
		return mesJsonErr
	}

	serverWritePkgSut := &utils.ServerWritePkgSut{
		Conn: loginSut.Conn,
		Mes:  mesJson,
	}
	writeNotifyMeOnlineErr := serverWritePkgSut.ServerWritePkg()
	if writeNotifyMeOnlineErr != nil {
		fmt.Printf("writeNotifyMeOnlineErr:%v\n", writeNotifyMeOnlineErr)
		return writeNotifyMeOnlineErr
	}

	return nil
}
