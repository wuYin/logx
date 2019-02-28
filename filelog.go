package logx

import (
	"fmt"
	"os"
	"time"
)

// 写出到文件的 logger
type FileLogWriter struct {
	format  string
	fName   string
	file    *os.File
	writeCh chan *LogRecord

	// 日志文件行数
	curLine int
	maxLine int

	// 日志文件大小
	curSize int
	maxSize int

	// 最多日志备份数
	maxBackup int
}

func NewFileLogWriter(fileName string) *FileLogWriter {
	w := &FileLogWriter{
		format:    DefaultFormat,
		fName:     fileName,
		writeCh:   make(chan *LogRecord, LogChanCapacity),
		maxBackup: DefaultMaxBackup,
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file %s fail: %v", fileName, err)
		return nil
	}
	w.file = f

	go w.run()

	return w
}

func (w *FileLogWriter) run() {
	for rec := range w.writeCh {

		// 日志已写满则备份
		if w.curLine >= w.maxLine || w.curSize >= w.maxSize {
			if err := w.backup(); err != nil {
				fmt.Fprintf(os.Stderr, "backup failed: %s", err)
				return
			}
		}

		n, err := fmt.Fprint(w.file, FormatLogRecord(w.format, rec))
		if err != nil {
			fmt.Fprintf(os.Stderr, "write failed: %s", err)
			return
		}

		w.curLine++
		w.curSize += n
	}
}

// 日志已满时将已有文件备份
// 日志命名规则：x.log x.log.1
func (w *FileLogWriter) backup() error {
	if _, err := os.Stat(w.fName); err != nil {
		return err
	}

	var backupPath string
	for n := w.maxBackup - 1; n >= 1; n-- {
		cur := fmt.Sprintf("%s.%d", w.fName, n)
		next := fmt.Sprintf("%s.%d", w.fName, n+1)
		if _, err := os.Stat(cur); err == nil {
			os.Rename(cur, next) // 自动淘汰最旧日志
		}
		backupPath = cur
	}

	// 关闭旧文件资源，准备备份
	w.file.Close()
	err := os.Rename(w.fName, backupPath)
	if err != nil {
		return fmt.Errorf("backup to %s failed: %s", backupPath, err)
	}

	// 重新建日志
	f, err := os.OpenFile(w.fName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	w.file = f

	w.curLine = 0
	w.curSize = 0

	return nil
}

func (w *FileLogWriter) LogWrite(rec *LogRecord) {
	w.writeCh <- rec
}

func (w *FileLogWriter) Close() {
	w.file.Close()
	time.Sleep(10 * time.Millisecond)
}

// 属性设置相关
func (w *FileLogWriter) SetFormat(format string) {
	w.format = format
}

func (w *FileLogWriter) SetMaxLine(line int) {
	w.maxLine = line
}

func (w *FileLogWriter) SetMaxSize(size int) {
	w.maxSize = size
}

func (w *FileLogWriter) SetMaxBackup(backup int) {
	w.maxBackup = backup
}
