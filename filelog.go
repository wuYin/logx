package logx

import (
	"fmt"
	"os"
	"time"
)

// 写文件时
type FileLogWriter struct {
	format  string
	fName   string
	file    *os.File
	writeCh chan *LogRecord

	// 日志文件行数
	curLine int
	maxLine int

	// 最多日志备份数
	maxBackup int
}

func NewFileLogWriter(fileName string, maxLine, maxBackup int) *FileLogWriter {
	if fileName == "" || maxLine == 0 {
		panic("invalid param")
	}

	w := &FileLogWriter{
		format:    DefaultFormat,
		fName:     fileName,
		maxLine:   maxLine,
		maxBackup: maxBackup,
		writeCh:   make(chan *LogRecord, LogChanCapacity),
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		return nil
	}
	w.file = f

	go w.run()

	return w
}

func (w *FileLogWriter) run() {
	for rec := range w.writeCh {
		// 日志已写满则备份
		if w.curLine >= w.maxLine {
			if err := w.backup(); err != nil {
				fmt.Fprintf(os.Stderr, "backup failed: %s", err)
				return
			}
		}

		_, err := fmt.Fprint(w.file, FormatLogRecord(w.format, rec))
		if err != nil {
			fmt.Fprintf(os.Stderr, "write failed: %s", err)
			return
		}

		w.curLine++
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

	return nil
}

func (w *FileLogWriter) LogWrite(rec *LogRecord) {
	w.writeCh <- rec
}

func (w *FileLogWriter) Close() {
	w.file.Close()
	time.Sleep(10 * time.Millisecond)
}
