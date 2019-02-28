package logx

import (
	"fmt"
	"testing"
	"time"
)

func TestFileLogWriter_LogWrite(t *testing.T) {
	w := NewFileLogWriter("demo.log")

	// // 测试1： ok
	// // 最大为 3 行最多备份 2 个，写入 10 行则第 0~2 行最旧日志将被自动丢弃
	// w.SetMaxLine(3)
	// w.SetMaxBackup(2)
	// for i := 0; i < 3*3+1; i++ {
	// 	rec := &LogRecord{
	// 		Level:   ERROR,
	// 		Created: time.Now(),
	// 		Source:  "filelog_test.go",
	// 		Message: fmt.Sprintf("test backup: %d", i),
	// 	}
	// 	w.LogWrite(rec)
	// }

	// 测试2：ok
	// 每行日志约 66 Byte，设置每个日志最大为 660 Byte，最多备份 2 个，写入 31 行则第 0~9 行最旧日志将被自动丢弃
	w.SetMaxLine(1000)
	w.SetMaxBackup(2)
	w.SetMaxSize(660)
	for i := 0; i < 3*10+1; i++ {
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
