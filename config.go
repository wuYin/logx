package logx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type LogConfigs []FilterConfig

// 配置 filter
type FilterConfig struct {
	Enable     bool   `json:"enable"`
	FilterType string `json:"filter_type"` // 暂只支持 console file 两种输出目标
	FilterName string `json:"filter_name"`
	FileName   string `json:"file_name"`
	MinLevel   string `json:"min_level"`
	Format     string `json:"format"`
	MaxLine    string `json:"max_line"`
	MaxSize    string `json:"max_size"`
}

// 从配置文件加载 logger
func LoadLogger(confPath string) (*Logger, error) {
	if _, err := os.Stat(confPath); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	var confs LogConfigs
	if err = json.Unmarshal(data, &confs); err != nil {
		return nil, err
	}

	logger := NewLogger()
	for _, conf := range confs {
		if !conf.Enable {
			continue
		}

		var w LogWriter
		switch conf.FilterType {
		case "console":
			w = loadConsoleFilter(conf)
		case "file":
			w, err = loadFileFilter(conf)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalid filter type: %s", conf.FilterType)
		}
		level := EnumLevel(conf.MinLevel)
		filter := &Filter{MinLevel: level, LogWriter: w}
		logger.AddFilter(conf.FilterName, level, filter)
	}

	return logger, nil
}

// 加载 console 类型的 filter
func loadConsoleFilter(conf FilterConfig) *ConsoleLogWriter {
	w := NewConsoleLogWriter()
	if conf.Format != "" {
		w.SetFormat(conf.Format)
	}
	return w
}

// 加载 file 类型的 filter
func loadFileFilter(conf FilterConfig) (*FileLogWriter, error) {
	maxLine := parseUnitNum(1000, conf.MaxLine)
	maxSize := parseUnitNum(1024, conf.MaxSize)
	if maxLine < 0 || maxSize < 0 || conf.FileName == "" {
		return nil, fmt.Errorf("invalid args: %v", conf)
	}

	w := NewFileLogWriter(conf.FileName)
	if conf.Format != "" {
		w.SetFormat(conf.Format)
	}
	w.SetMaxLine(maxLine)
	w.SetMaxSize(maxSize)
	return w, nil
}

// 解析单位表达式
func parseUnitNum(base int, unitNum string) int {
	if unitNum == "" {
		return 0
	}

	unit := 1
	n := len(unitNum)
	switch unitNum[n-1] {
	case 'g', 'G':
		unit *= base
	case 'm', 'M':
		unit *= base
		fallthrough
	case 'k', 'K': // 1024 bytes / 1000
		unit *= base
		unitNum = unitNum[:n-1] // 去掉单位
	}

	num, _ := strconv.Atoi(unitNum)
	return num * unit
}
