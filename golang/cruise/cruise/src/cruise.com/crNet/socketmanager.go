package crnet

import (
	"sync"
)

type socketmanager struct {
	wrLock     *sync.RWMutex
	socketlist []*TcpSocket
}

func newSocketManager() *socketmanager {
	mgr := new(socketmanager)
	mgr.socketlist = make([]*TcpSocket, 0)
	mgr.wrLock = new(sync.RWMutex)
	return mgr
}

func (this *socketmanager) add(sock *TcpSocket) {
	this.wrLock.Lock()
	defer this.wrLock.Unlock()
	this.socketlist = append(this.socketlist, sock)
}

// 检测网络连接状态
func (this *socketmanager) checkstatus() {

}
