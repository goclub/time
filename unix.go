package xtime

import "time"

func UnixMilli(t time.Time) int64  {
	return t.UnixNano() / 1e6
}

