package utils

import (
	"errors"
	"fmt"
	"net"
)

type ReadPkgSut struct{
	Conn net.Conn
	MesByte []byte
}

func(readPkgSut *ReadPkgSut) ReadPkg() error {

	readPkgSut.MesByte = make([]byte,1024)//创建一个切片接收读取的消息

	//首先接收对方发来信息的长度
	_ , clientReadPkgLenErr := readPkgSut.Conn.Read(readPkgSut.MesByte)
	if clientReadPkgLenErr != nil {
		fmt.Printf("clientReadPkgLenErr=%v\n",clientReadPkgLenErr)
		return clientReadPkgLenErr
	}

	//读取到的内容存储在readPkgSut.MesByte里 将readPkgSut.MesByte转为int获取信息的长度
	mesReadLen , mesReadLenErr := BytesToInt(readPkgSut.MesByte)
	if mesReadLenErr != nil {
		fmt.Printf("mesReadLenErr=%v\n",mesReadLenErr)
		return mesReadLenErr
	}

	//现在开始读取信息内容
	clientReadMes , clientReadMesErr := readPkgSut.Conn.Read(readPkgSut.MesByte)
	if clientReadMesErr != nil {
		fmt.Printf("clientReadMesErr=%v\n",clientReadMesErr)
		return clientReadMesErr
	}

	//判断接收的信息大小和刚才收到的大小是否相等
	if clientReadMes != mesReadLen {
		infoErr := fmt.Sprintf("信息包错误：{原信息包%v字节  当前信息包%v字节}",mesReadLen,clientReadMes)
		return errors.New(infoErr)
	}

	readPkgSut.MesByte = readPkgSut.MesByte[:mesReadLen]
	return nil
}
