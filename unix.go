package xtime

import "time"

// https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
func UnixMilli(t time.Time) int64  {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

