package logx

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

// 本地的格式缓存
// 每次格式化日志需要计算当前时间的时分秒信息，当数据量很大时计算代价不能忽略
// 于是在本地缓存 1s 过期时间的格式信息，当下一秒的新日志写入时触发缓存更新
type formatCache struct {
	LastUpdateAt         int64  // 格式最后更新的时间戳
	shortDate, shortTime string // 缓存当前秒的时间格式
	longDate, longTime   string
}

var (
	fmtCache = &formatCache{}
	fmtLock  = sync.Mutex{} // 缓存互斥锁
)

func setFmtCache(newCache *formatCache) {
	fmtLock.Lock()
	defer fmtLock.Unlock()
	fmtCache = newCache
}

func getFmtCache() *formatCache {
	fmtLock.Lock()
	defer fmtLock.Unlock()
	return fmtCache
}

// 格式化日志
// %T - Time 15:04:05 CST
// %t - Time 15:04
// %D - Date 2006/01/02
// %d - Date 01/02/06
// %L - Level
// %S - Source
// %M - Message
func FormatLogRecord(format string, rec *LogRecord) string {
	if format == "" || rec == nil {
		return ""
	}

	buf := bytes.NewBuffer(make([]byte, 0, 64))
	recAt := rec.Created
	recSec := rec.Created.Unix()

	cache := getFmtCache()

	// 缓存过期则更新
	if cache.LastUpdateAt != recSec {
		y, m, d := recAt.Year(), recAt.Month(), recAt.Day()
		h, i, s := recAt.Hour(), recAt.Minute(), recAt.Second()
		zone, _ := recAt.Zone()
		newCache := &formatCache{
			LastUpdateAt: recSec,
			shortTime:    fmt.Sprintf("%02d:%02d", h, i),
			shortDate:    fmt.Sprintf("%02d/%02d/%02d", y, m, d),
			longTime:     fmt.Sprintf("%02d:%02d:%02d %s", h, i, s, zone),
			longDate:     fmt.Sprintf("%04d/%02d/%02d", y, m, d),
		}
		cache = newCache
		setFmtCache(newCache)
	}

	// "[%D %T] [%L] (%S) %M" // 处理预留 [] 等符号
	signs := bytes.Split([]byte(format), []byte{'%'}) // 分割标志
	for i, sign := range signs {
		if i == 0 && len(sign) > 0 {
			buf.Write(sign) // 写入首个 % 前的字符串
			continue
		}
		switch sign[0] {
		case 'T':
			buf.WriteString(cache.longTime)
		case 't':
			buf.WriteString(cache.shortTime)
		case 'D':
			buf.WriteString(cache.longDate)
		case 'd':
			buf.WriteString(cache.shortDate)
		case 'L':
			buf.WriteString(logLevels[rec.Level])
		case 'S':
			buf.WriteString(rec.Source)
		case 's':
			paths := strings.Split(rec.Source, "/")
			buf.WriteString(paths[len(paths)-1])
		case 'M':
			buf.WriteString(rec.Message)
		}
		// 写入剩余字符
		if len(sign) > 1 {
			buf.Write(sign[1:])
		}
	}
	buf.WriteByte('\n')
	return buf.String()
}
