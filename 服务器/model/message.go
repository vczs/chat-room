package model

const (
	SmsMesType          =  "SmsMes"
	LoginMesType        =  "LoginMes"
	LoginResMesType     =  "LoginResMes"
	RegisterMesType     =  "RegisterMes"
	RegisterResMesType  =  "RegisterResMes"
	NotifyUserStateMesTyp =  "NotifyUserStateMes"

	Online = 1
	Offline = 2

	UserOffOnline = 0
	UserOnOnline = 1
)

//消息包
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"`//消息的类型
}

//登录
type LoginMes struct {
	Admin int `json:"admin"`
	Password string `json:"password"`
	AdminName string `json:"adminName"`
}
//登录结果
type LoginResMes struct {
	Code int  `json:"code"`//返回状态码 200登录成功 500登录失败 404服务器发生错误
	Error string  `json:"error"`//返回错误信息
	OnlineUser map[int]string `json:"onlineUser"`//返回当前所有在线用户Admin
}


//注册
type RegisterMes struct {
	User User
}
//注册结果
type RegisterResMes struct{
	Code int  `json:"code"`//返回状态码 201注册成功 501注册失败 404服务器发生错误
	Error string  `json:"error"`//返回错误信息
}

//通知用户状态
type NotifyUserStateMes struct{
	Admin int `json:"admin"`
	Flag int `json:"flag"`
	OnlineUser map[int]string `json:"onlineUser"`//返回当前所有在线用户Admin
}

type SmsMes struct {
	User User
	Content string`json:"content"`
}