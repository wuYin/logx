package logx

import (
	"testing"
	"time"
)

func TestLogger_Debug(t *testing.T) {
	l := NewLogger()
	l.AddFilter("console", INFO, NewConsoleLogWriter())
	l.Debug("Test|Debug|%v", 10)
	time.Sleep(1 * time.Second) // 等待日志写完毕
}
