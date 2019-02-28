package logx

import "time"

const (
	// 在写回前最多可以累计的日志数
	LogBufMsgs = 32
)

type Logger map[string]*Filter

func NewLogger() *Logger {
	l := make(Logger)
	return &l
}

type Filter struct {
	Level Level
	LogWriter
}

type Level int

// 日志级别
const (
	FINE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

var logLevels = [...]string{"FINE", "DEBG", "INFO", "WARN", "EROR", "FATL"}

type LogWriter interface {
	LogWrite(rec *LogRecord) // 写日志
	Close()                  // 日志写完毕后的资源清理工作
}

// 真正的日志行内容
type LogRecord struct {
	Level   Level     // 日志级别
	Created time.Time // 记录时间
	Source  string    // 源文件位置
	Message string    // 日志信息
}

// 将指定日志级别和写入位置的 LogWriter 添加到 logger
func (l Logger) AddFilter(name string, level Level, writer LogWriter) Logger {
	l[name] = &Filter{
		Level:     level,
		LogWriter: writer,
	}
	return l
}
