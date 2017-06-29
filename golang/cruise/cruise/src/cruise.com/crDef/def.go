package crdef

// 定义消息头
type MsgHeader struct {
	cmd  int32 // 协议号
	size int32 // 协议大小
	ver  int32 // 版本号
	seq  int32 // 序号
}
