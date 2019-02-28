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

	logger.Debug("Test|Debug|%v", 10)
}
