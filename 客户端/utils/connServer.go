package utils

import (
	"fmt"
	"net"
)

type ConnectServerSut struct {
	Conn net.Conn
}

func (connectServerSut *ConnectServerSut) ConnectServer() error {
	connectServer, dialServerErr := net.Dial("tcp", "127.0.0.1:9003")
	if dialServerErr != nil {
		fmt.Printf("dialServerErr：%v\n", dialServerErr)
		return dialServerErr
	}
	connectServerSut.Conn = connectServer
	return nil
}
