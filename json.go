package xtime

import (
	"time"
)

type ChinaTime struct {
	time.Time
}

func NewChinaTime(time time.Time) ChinaTime {
	return ChinaTime{Time: time.In(LocChina)}
}
func (t ChinaTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.In(LocChina).Format(LayoutTime) + `"`), nil
}
func (t *ChinaTime) UnmarshalJSON(b []byte) error {
	v, err := time.ParseInLocation(`"`+LayoutTime+`"`, string(b), LocChina)
	t.Time = v
	return err
}
