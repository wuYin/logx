package logx

import (
	"io/ioutil"
	"os"
	"testing"
)

var conf = `
[
    {
        "enable": true,
        "filter_type": "console",
        "filter_name": "console_1",
        "min_level": "INFO"
    },
    {
        "enable": true,
        "filter_type": "file",
        "filter_name": "file_1",
        "file_name": "warning.log",
        "min_level": "WARN",
        "format": "",
        "max_line": "3",
        "max_size": "2k"
    }
]
`

func TestLoadLogger(t *testing.T) {
	f, err := ioutil.TempFile(".", "conf.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	n, err := f.WriteString(conf)
	if err != nil || n != len(conf) {
		t.Fatal(n, err)
	}

	logger, err := LoadLogger(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Close()

	logger.Debug("Debug Logs|%s", "disk usage 40%")
	logger.Fatal("Fatal Logs|%s", "disk full")
}
