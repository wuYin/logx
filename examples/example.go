package main

import "github.com/wuYin/logx"

func main() {
	logger, err := logx.LoadLogger("config.json")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	// 写入所有文件
	logger.Debug("disk usage|%s", "40%")       // 不会写入 ERROR 级 urgent.log
	logger.Fatal("disk usage|%s", "disk full") //

	// 写入指定文件
	logger.DebugLog("handler", "ProfileHandler|uid invalid|%s", "1012912") // 只写入 handler.log
}
