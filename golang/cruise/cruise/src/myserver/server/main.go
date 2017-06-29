package main

import (
	"cruise.com/crLog"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	crlog.Init("./", crlog.LVL_DEBUG)
	crlog.Debug("cpu %v", runtime.NumCPU())

	//	for i := 0; i < 100000; i++ {
	//		crlog.Debug("test %v", i)
	//	}

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

	crlog.Debug("jhx GameServer shutdown!")
}
