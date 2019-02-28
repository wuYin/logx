 <img src="https://images.yinzige.com/2019-02-28-logx_v1.png" width=40%>

# logx

简单高效的 Golang 日志库

## Feature

- 多级别支持：`FINE, INFO, DEBG, WARN, EROR, FATL`
- 多输出支持：日志可输出到 console/file，其中 file 类型日志支持自动备份
- 多样化配置：支持配置单个日志最大行数，单个日志最大空间，最旧日志自动清理

## Usage

下载依赖：`go get github.com/wuYin/logx`

添加配置 `config.json`

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
        "file_name": "demo.log",
        "min_level": "DEBG",
        "max_line": "3",
        "max_size": "2k"
    }
]
```

执行代码：

```go
package main

import (
	"github.com/wuYin/logx"
)

func main() {
	logger, err := logx.LoadLogger("config.json")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Debug("Test|Debug|%v", 10) // line 14 // bingo
}
```

执行效果：

 <img src="https://images.yinzige.com/2019-02-28-logx.gif" width=90%>

## Version

- [x] **v0.1**  2019-02-28
  - ~~console 日志写入~~

  - ~~file 日志写入~~