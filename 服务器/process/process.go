package process

import (
	"client/functions"
	"client/model"
	"client/utils"
	"fmt"
	"io"
	"net"
)

type ProcessSut struct {
	Conn  net.Conn
	Ip    string
	Admin int
}

func (processSut *ProcessSut) Process() {

	defer func() {
		connCloseErr := processSut.Conn.Close() //关闭listen
		if connCloseErr != nil {
			fmt.Printf("connCloseErr=%v\n", connCloseErr)
			return
		}
		fmt.Printf("{%v}connClose suc...\n", processSut.Ip)
	}()

	for {
		//读取对方发来的信息
		serverReadPkgSut := &utils.ServerReadPkgSut{Conn: processSut.Conn}
		serverReadPkgErr := serverReadPkgSut.ServerReadPkg()
		if serverReadPkgErr != nil {
			if serverReadPkgErr == io.EOF {
				fmt.Printf("{%v}断开连接！\n", processSut.Ip)
			} else {
				fmt.Printf("serverReadPkgErr=%v\n", serverReadPkgErr)
			}
			admin, ok := utils.IpBind[processSut.Ip]
			if !ok {
				fmt.Printf("用户{%v}未绑定Admin！", processSut.Ip)
			} else {
				delete(utils.IpBind, processSut.Ip)
				functions.UserMgrPointer.DeleteOnlineUser(admin)
				//通知其他用户我离线了
				loginSut := functions.LoginSut{}
				loginSucNotifyOtherErr := loginSut.NotifyOtherOnline(admin, model.Offline)
				if loginSucNotifyOtherErr != nil {
					fmt.Printf("loginSucNotifyOtherErr:%v\n", loginSucNotifyOtherErr)
				}
			}
			return
		}
		fmt.Printf("json:%v\n", string(serverReadPkgSut.InfoRead))

		//成功读取到信息 解析信息内容 并执行解析到的对应功能
		mesProcessSut := &MesProcessSut{
			Ip:       processSut.Ip,
			Conn:     processSut.Conn,
			InfoRead: serverReadPkgSut.InfoRead,
		}
		mesProcessErr := mesProcessSut.MesProcess()
		if mesProcessErr != nil {
			fmt.Printf("mesProcessErr:%v\n", mesProcessErr)
			return
		}
		if mesProcessSut.Flag == 2 {
			continue
		}

		//将模块 执行结果信息包 返回给对方
		ServerWritePkgSut := &utils.ServerWritePkgSut{
			Conn: mesProcessSut.Conn,
			Mes:  mesProcessSut.ResPkg,
		}
		serverWritePkgErr := ServerWritePkgSut.ServerWritePkg()
		if serverWritePkgErr != nil {
			fmt.Printf("serverWritePkgErr:%v\n", serverWritePkgErr)
			return
		}
	}

}
