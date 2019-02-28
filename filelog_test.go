package logx

import (
	"testing"
	"time"
)

func TestFileLogWriter_LogWrite(t *testing.T) {
	w := NewFileLogWriter("demo.log")
	rec := &LogRecord{
		Level:   ERROR,
		Created: time.Now(),
		Source:  "filelog_test.go",
		Message: "test filelog write",
	}
	w.LogWrite(rec)
	time.Sleep(1 * time.Second)
}
