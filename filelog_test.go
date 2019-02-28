package logx

import (
	"fmt"
	"testing"
	"time"
)

func TestFileLogWriter_LogWrite(t *testing.T) {
	w := NewFileLogWriter("demo.log", 3, 2)

	// 最大为 3 行最多备份 2 个，写入 10 行第 0~2 行最旧日志将被自动丢弃
	for i := 0; i < 3*3+1; i++ {
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
