package logx

import (
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	l := NewLogger()
	defer l.Close()

	l.AddFilter("console", INFO, NewConsoleLogWriter())
	l.Debug("Test|Debug|%v", 10)
}
