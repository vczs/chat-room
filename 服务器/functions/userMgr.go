package functions

import (
	"fmt"
)

var (
	UserMgrPointer *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*LoginSut
}

func init() {
	UserMgrPointer = &UserMgr{
		onlineUsers: make(map[int]*LoginSut, 1024),
	}
}

//添加、修改
func (userMgr *UserMgr) AddOnlineUser(admin int, loginSut *LoginSut) {
	userMgr.onlineUsers[admin] = loginSut
}

//删除
func (userMgr *UserMgr) DeleteOnlineUser(admin int) {
	delete(userMgr.onlineUsers, admin)
}

//查询
func (userMgr *UserMgr) GetOnlineUser(admin int) (*LoginSut, error) {
	res, ok := userMgr.onlineUsers[admin]
	if !ok {
		getOnlineUserErr := fmt.Errorf("用户{%v}不在线！", admin)
		return nil, getOnlineUserErr
	}
	return res, nil
}

//返回列表
func (userMgr *UserMgr) GetAllOnlineUser() map[int]*LoginSut {
	return userMgr.onlineUsers
}
