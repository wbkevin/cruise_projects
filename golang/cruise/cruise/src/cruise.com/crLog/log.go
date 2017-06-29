// 自定义日志系统
// 支持以下功能：
// 1.按小时进行日志记录
// 2.按日志级别进行日志记录

package crlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	//"sync"
	"time"
)

const (
	// content of emnu Level ,level of log
	NULL = 1 << iota
	LVL_DEBUG
	LVL_INFO
	LVL_WARN
	LVL_NOTICE
	LVL_ERROR
	LVL_CRIT
)

type Outputer int

const (
	STD = iota
	FILE
)

type logger struct {
	logFd         *os.File // 文件描述符
	starLev       int      // 日志记录的等级
	buf           []byte   // 缓冲区
	path          string   // 路径
	baseName      string   // 通用名称
	logName       string   // 日志名称
	debugOutputer Outputer // 调试输出
	debugSwitch   bool     // 调试模式切换
	callDepth     int      // 日志文件记录深度
	fullPath      string   // 文件全路径
	lastHour      int      // 上一次记录的小时
	IsShowConsole bool     // 是否控制台显示
	//logList       []*string   // 日志列表
	//mutex         *sync.Mutex // 锁
	logChan chan string
}

var gLogger *logger

// 初始化入口
func Init(path string, lev int) error {
	gLogger = newLogger(path, "log", "Log4Golang", lev)
	gLogger.setCallDepth(3)
	gLogger.start()
	return nil
}

// 创建日志
func newLogger(path, baseName, logName string, level int) *logger {
	var err error
	logger := &logger{path: path, baseName: baseName, logName: logName, starLev: level}

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	logger.logFd = logger.getLoggerFd()

	logger.debugSwitch = true
	logger.debugOutputer = STD
	logger.callDepth = 3

	//logger.mutex = new(sync.Mutex)
	//logger.logList = make([]*string, 0)

	logger.logChan = make(chan string, 8096)

	return logger
}

func (this *logger) getLoggerFd() *os.File {
	var err error
	path := strings.TrimSuffix(this.path, "/")
	flag := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	this.fullPath = path + "/" + this.baseName
	now := time.Now()
	this.fullPath += fmt.Sprintf("%04d_%02d_%02d_%02d", now.Year(), now.Month(), now.Day(), now.Hour())
	this.logFd, err = os.OpenFile(this.fullPath, flag, 0666)
	if err != nil {
		panic(err)
	}
	return this.logFd
}

func (this *logger) start() {
	//go this.onTimer()

	//for i := 0; i < 5; i++ {
	//	go this.autoWrite()
	//}

	go this.autoWrite()
}

//// 1秒写一次
//func (this *logger) onTimer() {
//	timer := time.NewTicker(1 * time.Second)
//	for {
//		select {
//		case <-timer.C:
//			for i := 0; i < 5; i++ {
//				go this.autoWrite()
//			}

//		}
//	}
//}

// 每次最多写入2000条数据
func (this *logger) autoWrite() {
	//	nIndex := 0
	//	this.buf = this.buf[:0]
	//	for {
	//		if nIndex > 2000 {
	//			break
	//		}
	//		strLog := this.popLog()
	//		if nil == strLog {
	//			break
	//		}
	//		this.buf = append(this.buf, bytes.NewBufferString(*strLog).Bytes()...)
	//		nIndex++
	//	}

	//	if len(this.buf) != 0 {
	//		this.writeLog(this.buf)
	//	}

	for {
		//		select {
		//		case str := <-this.logChan:
		//			this.writeLog(bytes.NewBufferString(str).Bytes())
		//		default:
		//		}
		str := <-this.logChan
		this.writeLog(bytes.NewBufferString(str).Bytes())
	}
}

func (this *logger) writeLog(buf []byte) {
	now := time.Now()
	if now.Hour() != this.lastHour {
		//先将当前文件关闭
		err := this.logFd.Close()
		if err != nil {
			str := fmt.Sprintf("关闭日志文件[%v]失败[err:%v]", this.fullPath, err.Error())
			fmt.Println(str)
		}
		//获取下一个索引的文件
		this.logFd = this.getLoggerFd()
	}
	_, err := this.logFd.Write(buf)
	if err != nil {
		fmt.Printf("写入错误")
	}
}

func (this *logger) output(fd io.Writer, level, prefix string, format string, v ...interface{}) (err error) {
	var msg string
	if format == "" {
		msg = fmt.Sprintln(v...)
	} else {
		msg = fmt.Sprintf(format, v...)
	}

	this.buf = this.buf[:0]

	this.buf = append(this.buf, "["+this.logName+"]"...)
	this.buf = append(this.buf, level...)
	this.buf = append(this.buf, prefix...)

	this.buf = append(this.buf, ":"+msg...)
	if len(msg) > 0 && msg[len(msg)-1] != '\n' {
		this.buf = append(this.buf, '\n')
	}

	_, err = fd.Write(this.buf)

	return
}

func (l *logger) setCallDepth(d int) {
	l.callDepth = d
}

func (l *logger) openDebug() {
	l.debugSwitch = true
}

func (l *logger) getFileLine() string {
	_, file, line, ok := runtime.Caller(l.callDepth)
	if !ok {
		file = "???"
		line = 0
	}

	return file + ":" + itoa(line, -1)
}

/**
* Change from Golang's log.go
* Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
* Knows the buffer has capacity.
 */
func itoa(i int, wid int) string {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		return "0"
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	return string(b[bp:])
}

func (l *logger) getTime() string {
	// Time is yyyy-mm-dd hh:mm:ss.microsec
	var buf []byte
	t := time.Now()
	year, month, day := t.Date()
	buf = append(buf, itoa(int(year), 4)+"-"...)
	buf = append(buf, itoa(int(month), 2)+"-"...)
	buf = append(buf, itoa(int(day), 2)+" "...)

	hour, min, sec := t.Clock()
	buf = append(buf, itoa(hour, 2)+":"...)
	buf = append(buf, itoa(min, 2)+":"...)
	buf = append(buf, itoa(sec, 2)...)

	buf = append(buf, '.')
	buf = append(buf, itoa(t.Nanosecond()/1e3, 6)...)

	return string(buf[:])
}

func (l *logger) closeDebug() {
	l.debugSwitch = false
}

func (l *logger) setDebugOutput(o Outputer) {
	l.debugOutputer = o
}

func Debug(format string, v ...interface{}) error {
	return gLogger.addlog(LVL_DEBUG, format, v...)
}

func Info(format string, v ...interface{}) error {
	return gLogger.addlog(LVL_INFO, format, v...)
}

func Warn(format string, v ...interface{}) error {
	return gLogger.addlog(LVL_WARN, format, v...)
}

func Notice(format string, v ...interface{}) error {
	return gLogger.addlog(LVL_NOTICE, format, v...)
}

func Error(format string, v ...interface{}) error {
	return gLogger.addlog(LVL_ERROR, format, v...)
}

func Crit(format string, v ...interface{}) error {
	return gLogger.addlog(LVL_CRIT, format, v...)
}

func (this *logger) toString(logType int) string {
	var str string = ""
	switch logType {
	case LVL_DEBUG:
		str = "[DEBUG]"
	case LVL_INFO:
		str = "[INFO]"
	case LVL_WARN:
		str = "[WARN]"
	case LVL_ERROR:
		str = "[ERROR]"
	case LVL_CRIT:
		str = "[CRIT]"
	default:
		str = "[DEBUG]"
	}
	return str
}

func (this *logger) addlog(logLev int, format string, v ...interface{}) error {
	if logLev < this.starLev {
		return nil
	}

	level := this.toString(logLev)

	// 直接打印
	prefix := "[" + this.getTime() + "][" + this.getFileLine() + "]"

	//! 打印堆栈信息
	// 只打印三层调用
	//	var logContentBuf = new(bytes.Buffer)
	//	skip := 2
	//	for i := 0; i < 3; i++ {
	//		_, file, line, ok := runtime.Caller(skip)
	//		if !ok {
	//			break
	//		}
	//
	//		fmt.Fprintf(logContentBuf, fmt.Sprintf("at [%s:%d]\n", file, line))
	//
	//		skip++
	//	}
	//	prefix := fmt.Sprintf("[%s]\n%s", _logger.getTime(), logContentBuf.Bytes())

	var msg string
	if format == "" {
		msg = fmt.Sprint(v...)
	} else {
		msg = fmt.Sprintf(format, v...)
	}

	// [日志级别][时间][触发文件]日志内容
	strLog := fmt.Sprintf("%s %s %s", level, prefix, msg)
	fmt.Println(strLog)

	//	this.mutex.Lock()
	//	defer this.mutex.Unlock()
	//	strLog += "\n"
	//	this.logList = append(this.logList, &strLog)

	strLog += "\n"
	this.logChan <- strLog

	return nil
}

func (this *logger) popLog() *string {

	//	this.mutex.Lock()
	//	defer this.mutex.Unlock()
	//	if 0 == len(this.logList) {
	//		return nil
	//	}
	//	data := this.logList[0]
	//	this.logList = this.logList[1:]
	//	return data

	str := <-this.logChan
	return &str
}

/*
func (this *logger) getSize() int {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return len(this.logList)
}

*/
