package logx

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

// 本地秒级数据格式缓存
type secondCache struct {
	LastUpdateAt         int64
	shortDate, shortTime string
	longDate, longTime   string
}

var (
	secCache  = &secondCache{}
	cacheLock = sync.Mutex{} // 缓存互斥锁
)

func setSecCache(newCache *secondCache) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	secCache = newCache
}

func getSecCache() *secondCache {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	return secCache
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

	cache := getSecCache()

	// 缓存过期则更新
	if cache.LastUpdateAt != recSec {
		y, m, d := recAt.Year(), recAt.Month(), recAt.Day()
		h, i, s := recAt.Hour(), recAt.Minute(), recAt.Second()
		zone, _ := recAt.Zone()
		newCache := &secondCache{
			LastUpdateAt: recSec,
			shortTime:    fmt.Sprintf("%02d:%02d", h, i),
			shortDate:    fmt.Sprintf("%02d/%02d/%02d", y, m, d),
			longTime:     fmt.Sprintf("%02d:%02d:%02d %s", h, i, s, zone),
			longDate:     fmt.Sprintf("%04d/%02d/%02d", y, m, d),
		}
		cache = newCache
		setSecCache(newCache)
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
