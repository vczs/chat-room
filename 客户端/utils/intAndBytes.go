package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func IntToBytes(n int) ([]byte , error) {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	intToBytesErr := binary.Write(bytesBuffer,binary.BigEndian, x)
	if intToBytesErr != nil {
		fmt.Printf("intToBytesErr：%v\n",intToBytesErr)
		return nil,intToBytesErr
	}
	return bytesBuffer.Bytes() , nil
}

func BytesToInt(b []byte) (int , error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	bytesToIntErr := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if bytesToIntErr != nil {
		fmt.Printf("bytesToIntErr：%v\n",bytesToIntErr)
		return 0 , bytesToIntErr
	}
	return int(x) , nil
}