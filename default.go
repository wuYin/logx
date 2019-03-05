package logx

import (
	"fmt"
	"strings"
)

var stdLogger = NewLogger()

func init() {
	console := NewConsoleLogWriter()
	console.SetFormat("[%D %T] [%L] %M")
	stdLogger = &Logger{
		"stdout": &Filter{DEBUG, console},
	}
}

func Fine(arg0 interface{}, args ...interface{}) {
	level := FINE
	switch v := arg0.(type) {
	case string:
		stdLogger.dispatch(level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		stdLogger.dispatch(level, format, args...)
	}
}

func Info(arg0 interface{}, args ...interface{}) {
	level := INFO
	switch v := arg0.(type) {
	case string:
		stdLogger.dispatch(level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		stdLogger.dispatch(level, format, args...)
	}
}

func Debug(arg0 interface{}, args ...interface{}) {
	level := DEBUG
	switch v := arg0.(type) {
	case string:
		stdLogger.dispatch(level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		stdLogger.dispatch(level, format, args...)
	}
}

func Warn(arg0 interface{}, args ...interface{}) {
	level := WARN
	switch v := arg0.(type) {
	case string:
		stdLogger.dispatch(level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		stdLogger.dispatch(level, format, args...)
	}
}

func Error(arg0 interface{}, args ...interface{}) {
	level := ERROR
	switch v := arg0.(type) {
	case string:
		stdLogger.dispatch(level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		stdLogger.dispatch(level, format, args...)
	}
}

func Fatal(arg0 interface{}, args ...interface{}) {
	level := FATAL
	switch v := arg0.(type) {
	case string:
		stdLogger.dispatch(level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		stdLogger.dispatch(level, format, args...)
	}
}
