package logx

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	LogBufMsgs = 32 // 在写回前最多可以累计的日志数
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

type LogWriter interface {
	LogWrite(rec *LogRecord) // 写日志
	Close()                  // 日志写完毕后的资源清理工作
}

type Level int

// 日志级别
const (
	FINE Level = iota
	INFO
	DEBUG
	WARN
	ERROR
	FATAL
)

var logLevels = [...]string{"FINE", "INFO", "DEBG", "WARN", "EROR", "FATL"}

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

// 记录 debug 日志
func (l Logger) Debug(arg0 interface{}, args ...interface{}) {
	level := DEBUG
	switch v := arg0.(type) {
	case string:
		l.dispatch(level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args)) // 原样输出
		l.dispatch(level, format, args...)
	}
}

// 日志分发
func (l Logger) dispatch(level Level, format string, args ...interface{}) {
	// 是否有能写 level 级的 filter
	valid := false
	for _, f := range l {
		if f.Level <= level {
			valid = true
			break
		}
	}
	if !valid {
		return
	}

	// 获取日志调用处
	var src string
	pc, _, fileLine, ok := runtime.Caller(2)
	if ok {
		src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), fileLine)
	}

	// 格式化日志
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	// 分发日志
	rec := &LogRecord{
		Level:   level,
		Created: time.Now(),
		Source:  src,
		Message: msg,
	}

	for _, f := range l {
		if f.Level <= level {
			f.LogWrite(rec)
		}
	}
}
