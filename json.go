package xtime


import (
"time"
)


type ChinaTime struct {
	time.Time
}
func NewChinaTime(time time.Time) ChinaTime {
	return ChinaTime{Time: time.In(LocationChina)}
}
func (t ChinaTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.In(LocationChina).Format(LayoutTime) + `"`), nil
}
func (t *ChinaTime) UnmarshalJSON(b []byte) error {
	v, err := time.ParseInLocation(`"` + LayoutTime + `"`, string(b), LocationChina)
	t.Time = v
	return err
}
