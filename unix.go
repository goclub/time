package xtime

import "time"

// UnixMilli 老版本 go 没有 time.Time{}.UnixMilli() 方法,故此提供了 xtime.UnixMilli(t)
func UnixMilli(t time.Time) int64  {
	return t.UnixNano() / 1e6
}

