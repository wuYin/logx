package logx

import (
	"testing"
	"time"
)

func TestConsoleLogWriter_LogWrite(t *testing.T) {
	l := NewConsoleLogWriter()
	record := &LogRecord{
		Level:   WARNING,
		Created: time.Now(),
		Source:  "logx/console_test.go",
		Message: "test warning msg",
	}
	l.LogWrite(record)
	time.Sleep(1 * time.Second)
}
