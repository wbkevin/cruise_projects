package crnet

import "net"

//tcp 会话
type TcpSession struct {
	conn net.Conn
}

func (this *TcpSession) Send() {

}

// 接受数据
func (this *TcpSession) recv() error {
	return nil
}
