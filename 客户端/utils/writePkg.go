package utils

import (
	"fmt"
	"net"
	"time"
)

type WritePkgSut struct {
	Conn    net.Conn
	JsonCon []byte
}

func (writePkgSut *WritePkgSut) WritePkg() error {

	//用len(mesJson)获取mesJson的int类型的长度
	mesJsonIntLen := len(writePkgSut.JsonCon)
	//将int类型的mesJsonIntLen转为[]byte类型
	mesJsonBytesLen, mesJsonBytesLenErr := IntToBytes(mesJsonIntLen)
	if mesJsonBytesLenErr != nil {
		fmt.Printf("mesJsonBytesLenErr：%v\n", mesJsonBytesLenErr)
		return mesJsonBytesLenErr
	}

	//先发送mesJson的长度给服务端
	_, writeMesJsonLenErr := writePkgSut.Conn.Write(mesJsonBytesLen)
	if writeMesJsonLenErr != nil {
		fmt.Printf("writeMesJsonLenErr：%v\n", writeMesJsonLenErr)
		return writeMesJsonLenErr
	}

	time.Sleep(time.Duration(1) * time.Microsecond) //休息一会  防止消息长度和内容连续发送无间隔 导致服务端无法区别

	//发送mesJson给服务器
	_, writeMesJsonErr := writePkgSut.Conn.Write(writePkgSut.JsonCon)
	if writeMesJsonErr != nil {
		fmt.Printf("writeMesJsonErr：%v\n", writeMesJsonErr)
		return writeMesJsonErr
	}

	return nil
}
