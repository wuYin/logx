 <img src="https://images.yinzige.com/2019-02-28-logx_v1.png" width=40%>

# logx

简单高效的 Golang 日志库

## Feature

- 多级别支持：`FINE, INFO, DEBUG, WARN, ERROR, FATAL`
- 多输出支持：日志可输出到 console / file，其中 file 类型日志支持自动备份
- 多样化配置：支持配置单个日志最大行数 / 最大空间，支持最旧日志自动清理

## Usage

### 添加配置： `config.json`

```json
[
    {
        "enable": true,
        "filter_type": "console",
        "filter_name": "iterm2",
        "min_level": "INFO"
    },
    {
        "enable": true,
        "filter_type": "file",
        "filter_name": "handler",
        "file_name": "logs/handler.log",
        "min_level": "DEBG",
        "max_line": "2"
    },
    {
        "enable": true,
        "filter_type": "file",
        "filter_name": "urgent",
        "file_name": "urgent.log",
        "min_level": "EROR",
        "max_size": "20M"
    }
]
```



### 执行代码：

```go
package main

import "github.com/wuYin/logx"

func main() {
	logger, err := logx.LoadLogger("config.json")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	// 写入所有文件
	logger.Debug("disk usage|%s", "40%")
	logger.Fatal("disk usage|%s", "disk full")

	// 写入指定文件
	logger.DebugLog("handler", "ProfileHandler|uid invalid|%s", "1012912")
}
```



### 执行输出

- 日志分级：Debug 日志 `disk usage|%40` 未写入 EROR 级的 urgent.log
- 日志分割：handler 配置了单个日志最多 2 行，如上的 3 条日志被切分到两个 handler.*log
- 单文件写入：DebugLog 只写入到 handler.log

![](https://images.yinzige.com/2019-03-01-093031.png)



## Version

- [x] **v0.1**  2019-02-28
  - ~~console 日志写入~~
  - ~~file 日志写入~~
  - ~~日志写入到指定 filter~~