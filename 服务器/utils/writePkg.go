package utils

import (
	"fmt"
	"net"
)

type ServerWritePkgSut struct {
	Conn net.Conn
	Mes []byte
}

func (serverWritePkgSut *ServerWritePkgSut) ServerWritePkg() error {

	//获取serverWritePkgSut.Mes的int类型长度
	mesLen := len(serverWritePkgSut.Mes)
	//将int类型的Mes的长度转为[]byte类型
	mesLenBytes , mesLenBytesErr := IntToBytes(mesLen)
	if mesLenBytesErr != nil {
		fmt.Printf("mesLenBytesErr：%v\n",mesLenBytesErr)
		return mesLenBytesErr
	}

	//先发送mes的长度给对方
	_ , writeMesLenErr := serverWritePkgSut.Conn.Write(mesLenBytes)
	if writeMesLenErr != nil {
		fmt.Printf("writeMesLenErr：%v\n",writeMesLenErr)
		return writeMesLenErr
	}

	//发送mes内容给对方
	_ , writeMesErr := serverWritePkgSut.Conn.Write(serverWritePkgSut.Mes)
	if writeMesErr != nil {
		fmt.Printf("writeMesJsonErr：%v\n",writeMesErr)
		return writeMesErr
	}

	return nil
}