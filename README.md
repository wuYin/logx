 <img src="https://images.yinzige.com/2019-02-28-logx_v1.png" width=40%>

# logx

简单高效的 Golang 日志库

## Feature

- 多级别支持：`FINE, INFO, DEBUG, WARN, ERROR, FATAL`
- 多输出支持：日志可输出到 console / file，其中 file 类型日志支持自动备份
- 多样化配置：支持配置单个日志最大行数，单个日志最大空间，最旧日志自动清理

## Usage

下载依赖：`go get github.com/wuYin/logx`

添加配置： `config.json`

```json
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
        "max_line": "3",
        "max_size": "2k"
    }
]
```

执行代码：

```go
package main

import "github.com/wuYin/logx"

func main() {
	logger, err := logx.LoadLogger("config.json")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Debug("Debug Logs|%s", "disk usage 40%")
	logger.Fatal("Fatal Logs|%s", "disk full") 
}
```

执行写入：配置的 `WARN` 级 file 日志不会写入 `DEBUG` 信息，而 `INFO` 级 console 日志则全部输出：

 <img src="https://images.yinzige.com/2019-03-01-logx2.gif" width=90%>

## Version

- [x] **v0.1**  2019-02-28
  - ~~console 日志写入~~

  - ~~file 日志写入~~