package logx

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// 将各级别日志记录到指定的 filter 中
func (l Logger) FineLog(filter string, arg0 interface{}, args ...interface{}) {
	level := FINE
	switch v := arg0.(type) {
	case string:
		l.dispatch2Filter(filter, level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args)) // 原样输出
		l.dispatch2Filter(filter, level, format, args...)
	}
}

func (l Logger) InfoLog(filter string, arg0 interface{}, args ...interface{}) {
	level := INFO
	switch v := arg0.(type) {
	case string:
		l.dispatch2Filter(filter, level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		l.dispatch2Filter(filter, level, format, args...)
	}
}

func (l Logger) DebugLog(filter string, arg0 interface{}, args ...interface{}) {
	level := DEBUG
	switch v := arg0.(type) {
	case string:
		l.dispatch2Filter(filter, level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		l.dispatch2Filter(filter, level, format, args...)
	}
}

func (l Logger) WarnLog(filter string, arg0 interface{}, args ...interface{}) {
	level := WARN
	switch v := arg0.(type) {
	case string:
		l.dispatch2Filter(filter, level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		l.dispatch2Filter(filter, level, format, args...)
	}
}

func (l Logger) ErrorLog(filter string, arg0 interface{}, args ...interface{}) {
	level := ERROR
	switch v := arg0.(type) {
	case string:
		l.dispatch2Filter(filter, level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		l.dispatch2Filter(filter, level, format, args...)
	}
}

func (l Logger) FatalLog(filter string, arg0 interface{}, args ...interface{}) {
	level := FATAL
	switch v := arg0.(type) {
	case string:
		l.dispatch2Filter(filter, level, v, args...)
	default:
		format := fmt.Sprint(arg0) + strings.Repeat(" %v", len(args))
		l.dispatch2Filter(filter, level, format, args...)
	}
}

// 将日志写到指定的 filter 中
func (l Logger) dispatch2Filter(name string, level Level, format string, args ...interface{}) {
	filter, ok := l.getFilter(name)
	if !ok || filter.MinLevel > level {
		return
	}

	var src string
	pc, _, fileLine, ok := runtime.Caller(2)
	if ok {
		src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), fileLine)
	}

	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	rec := &LogRecord{
		Level:   level,
		Created: time.Now(),
		Source:  src,
		Message: msg,
	}
	filter.LogWrite(rec)
}

// 获取指定的 filter
func (l Logger) getFilter(filterName string) (*Filter, bool) {
	f, ok := l[filterName]
	if !ok {
		f, ok = l["stdout"]
	}
	return f, ok
}
