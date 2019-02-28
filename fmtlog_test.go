package logx

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatLogRecord(t *testing.T) {
	now := time.Now()
	nowStr := now.Format("2006/01/02 15:04:05")
	zone, _ := now.Zone()
	rec := &LogRecord{
		Level:   DEBUG,
		Created: now,
		Source:  "logx/fmtlog_test.go", // runtime 获取处理
		Message: "test debug msg",
	}
	format := "[%D %T] [%L] (%S) %M"
	want := fmt.Sprintf("[%s %s] [DEBG] (logx/fmtlog_test.go) test debug msg\n", nowStr, zone)
	got := FormatLogRecord(format, rec)
	if want != got {
		t.Fatalf("want %s \ngot %s", want, got)
	}
}
