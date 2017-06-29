package main

import (
	"cruise.com/crlog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	crlog.Init("./", crlog.LVL_DEBUG)

	chTimeOut := time.Tick(10 * time.Second)
	osSignalCh := make(chan os.Signal, 1)
	signal.Notify(osSignalCh, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	for {
		select {
		case s := <-osSignalCh:
			crlog.Debug("Receive os.signal:%v", s)

			// 关闭网络
			// 发送消息通知所有链接用户
			// 保存所有用户数据
			// 记录退出日志
			//gGameService.Stop()
			//comm.StopPerformanceMonitor("game")
			return
		case <-chTimeOut:
		}
	}
}
