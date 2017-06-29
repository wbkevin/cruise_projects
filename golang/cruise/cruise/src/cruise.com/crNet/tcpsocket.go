package crnet

import "net"

// socket 通信
type TcpSocket struct {
	conn     *net.Conn
	recvChan chan interface{} // 数据接收通道
	sendChan chan interface{} // 数据发送通道
}

func NewTcpSocket(conn *net.Conn) *TcpSocket {
	sock := new(TcpSocket)
	sock.conn = conn

	sock.recvChan = make([]interface{}, 1024)
	sock.sendChan = make([]interface{}, 1024)

	// 启动接收线程
	go sock.onRecv()

	// 启动发送线程
	go sock.onSend()

	return sock
}

// 数据接收线程
func (this *TcpSocket) onRecv() {
	for {

	}
}

// 数据发送线程
func (this *TcpSocket) onSend() {

}
