package xtime

import "time"


func ParseChinaYear(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutYear, value, LocationChina)
}
func ParseChinaYearAndMonth(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutYearAndMonth, value, LocationChina)
}
func ParseChinaDate(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutDate, value, LocationChina)
}
func ParseChinaTime(value string) (t time.Time, err error) {
	return time.ParseInLocation(LayoutTime, value, LocationChina)
}