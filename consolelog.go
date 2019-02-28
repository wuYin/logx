package logx

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 写出到终端的 logger
type ConsoleLogWriter struct {
	format  string
	writeCh chan *LogRecord // 传输日志的缓冲 channel
}

var (
	consoleOut io.Writer = os.Stdout // 写入目标
)

func NewConsoleLogWriter() *ConsoleLogWriter {
	w := &ConsoleLogWriter{
		format:  DefaultFormat,
		writeCh: make(chan *LogRecord, LogChanCapacity),
	}
	go w.run(consoleOut)
	return w
}

func (cw *ConsoleLogWriter) SetFormat(format string) {
	cw.format = format
}

func (cw *ConsoleLogWriter) run(out io.Writer) {
	for rec := range cw.writeCh {
		fmt.Fprint(out, FormatLogRecord(cw.format, rec)) // 写入日志
	}
}

// 负责 logger 的日志写入
// 注意：若写入过快导致缓冲 channel 装满此方法会阻塞
func (cw *ConsoleLogWriter) LogWrite(rec *LogRecord) {
	cw.writeCh <- rec
}

// 关闭此 logger
// 负责资源的回收
func (cw *ConsoleLogWriter) Close() {
	close(cw.writeCh)
	// TODO: 提高日志写回机制可靠性
	time.Sleep(10 * time.Millisecond)
}
