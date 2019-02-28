package logx

import (
	"fmt"
	"testing"
	"time"
)

func TestFileLogWriter_LogWrite(t *testing.T) {
	w := NewFileLogWriter("demo.log", 3, 10)

	for i := 0; i < 4; i++ {
		rec := &LogRecord{
			Level:   ERROR,
			Created: time.Now(),
			Source:  "filelog_test.go",
			Message: fmt.Sprintf("test backup: %d", i),
		}
		w.LogWrite(rec)
	}

	time.Sleep(1 * time.Second)
}
