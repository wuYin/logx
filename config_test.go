package logx

import (
	"testing"
)

func TestLoadLogger(t *testing.T) {
	logger, err := LoadLogger("conf.json")
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Close()

	logger.Debug("Debug Logs|%s", "disk usage 40%")
	logger.Fatal("Fatal Logs|%s", "disk full")
}
