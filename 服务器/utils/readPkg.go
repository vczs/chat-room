package utils

import (
	"errors"
	"fmt"
	"net"
)

type ServerReadPkgSut struct{
	Conn net.Conn
	InfoRead []byte
}

func (serverReadPkgSut *ServerReadPkgSut)ServerReadPkg() error {

	serverReadPkgSut.InfoRead = make([]byte,1024)//创建一个切片接收读取的消息

	//首先接收客户端发来信息的长度
	_ , serverReadMesLenErr := serverReadPkgSut.Conn.Read(serverReadPkgSut.InfoRead)
	if serverReadMesLenErr != nil {
		fmt.Printf("serverReadMesLenErr:%v\n",serverReadMesLenErr)
		return serverReadMesLenErr
	}

	//读取到的内容存储在buf里 将buf转为int获取信息的长度
	mesReadLen , mesReadLenErr := BytesToInt(serverReadPkgSut.InfoRead)
	if mesReadLenErr != nil {
		fmt.Printf("mesReadLenErr=%v\n",mesReadLenErr)
		return mesReadLenErr
	}

	//现在开始读取信息内容
	serverReadMes , serverReadMesErr := serverReadPkgSut.Conn.Read(serverReadPkgSut.InfoRead)
	if serverReadMesErr != nil {
		fmt.Printf("serverReadMesErr:%v\n",serverReadMesErr)
		return serverReadMesErr
	}

	//判断接收的信息大小和刚才收到的大小是否相等
	if serverReadMes != mesReadLen {
		infoErr := fmt.Sprintf("信息包错误：{原信息包%v字节  当前信息包%v字节}",mesReadLen,serverReadMes)
		return errors.New(infoErr)
	}
	//如果信息大小和刚才收到的大小相等 就将将信息内容存储到serverReadPkgSut.InfoRead中
	serverReadPkgSut.InfoRead = serverReadPkgSut.InfoRead[:mesReadLen]

	return nil //error返回空
}
