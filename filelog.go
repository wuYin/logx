package logx

import (
	"fmt"
	"os"
	"time"
)

type FileLogWriter struct {
	format   string
	fileName string
	file     *os.File
	writeCh  chan *LogRecord
}

func NewFileLogWriter(fileName string) *FileLogWriter {
	w := &FileLogWriter{
		format:   DefaultFormat,
		fileName: fileName,
		writeCh:  make(chan *LogRecord, LogChanCapacity),
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
		s := FormatLogRecord(w.format, rec)
		n, err := w.file.WriteString(s)
		if err != nil || n != len(s) {
			fmt.Println(err, n)
		}
	}
}

func (w *FileLogWriter) LogWrite(rec *LogRecord) {
	w.writeCh <- rec
}

func (w *FileLogWriter) Close() {
	w.file.Close()
	time.Sleep(10 * time.Millisecond)
}
