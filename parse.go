package xtime

import "time"


func ParseChinaYear(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutYear, value, LocChina)
}
func ParseChinaYearAndMonth(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutYearAndMonth, value, LocChina)
}
func ParseChinaDate(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutDate, value, LocChina)
}
func ParseChinaTime(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutTime, value, LocChina)
}