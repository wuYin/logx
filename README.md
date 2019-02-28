 <img src="https://images.yinzige.com/2019-02-28-logx_v1.png" width=40%>

# logx

简单高效的 Golang 日志库

## Feature

- 多级别日志类型支持：`FINE, INFO, DEBG, WARN, EROR, FATL`
- 多输出：v0.1 支持 console stdout 输出

## Usage

下载依赖：`go get github.com/wuYin/logx`

输出日志：

```go
package main

import "github.com/wuYin/logx"

func main() {
	l := logx.NewLogger()
	defer l.Close()

	l.AddFilter("console", logx.INFO, logx.NewConsoleLogWriter())
	l.Debug("Test|Debug|%v", 10)
}
```

执行效果：

 <img src="https://images.yinzige.com/2019-02-28-103201.png" width=75%>

## Version

- [ ] **v0.1**  2019-02-28

  - ~~console 日志级别写入~~

  - file 日志级别写入