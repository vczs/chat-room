package login

import (
	"client/model"
	"net"
)

var curUser CurUser

type CurUser struct {
	Conn net.Conn
	User model.User
}
