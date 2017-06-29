package crnet

import (
	"cruise.com/crLog"
	"fmt"
	"net"
)

// 服务器端监听者
type ServerListener struct {
	listener  net.Listener   // 监听者
	isRunning bool           // 运行标记
	sockMgr   *socketmanager // socket管理器
}

// 创建监听者
func NewServerListener() *ServerListener {
	listener := new(ServerListener)
	listener.sockMgr = newSocketManager()
	return listener
}

// 启动对指定端口的监听
func (this *ServerListener) Start(port uint32) error {

	crlog.Debug("start listernning at port[%v]...", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	this.listener = ln
	this.isRunning = true

	// 开始网络检测
	go this.sockMgr.checkstatus()

	// 接受连接
	go this.accept()

	return nil
}

func (this *ServerListener) accept() error {
	for this.isRunning {
		conn, err := this.listener.Accept()
		if err != nil {
			return err
		}
		crlog.Debug("[%v] connected...", conn.RemoteAddr())

		sock := NewTcpSocket(&conn)
		if nil == sock {
			crlog.Debug("new tcpsocket err, conn=%v", conn)
		}

		this.sockMgr.add(sock)
	}

	crlog.Crit("server listener exited.")
	return nil
}
