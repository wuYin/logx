package logx

import (
	"testing"
	"time"
)

func TestLogger_DebugLog(t *testing.T) {
	l := NewLogger()
	defer l.Close()

	l.AddFilter("service", DEBUG, NewFileLogWriter("logs/service.log"))
	l.AddFilter("handler", DEBUG, NewFileLogWriter("logs/handler.log"))
	l.AddFilter("urgent", ERROR, NewFileLogWriter("logs/urgent.log"))
	l.ErrorLog("handler", "UserHandler|Uid|%s", "1023123627834637")
	l.DebugLog("service", "ProfileService|User|%s", "wuYin")
	l.FatalLog("urgent", "System|%s", "disk full")

	time.Sleep(100 * time.Millisecond)
}
