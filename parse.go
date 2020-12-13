package xtime

import "time"


func ParseChinaYear(value string) (t time.Time, err error) {
	return time.ParseInLocation(layoutYear, value, LocationChina)
}
func ParseChinaYearAndMonth(value string) (t time.Time, err error) {
	return time.ParseInLocation(layoutYearAndMonth, value, LocationChina)
}
func ParseChinaDate(value string) (t time.Time, err error) {
	return time.ParseInLocation(layoutDate, value, LocationChina)
}
func ParseChinaTime(value string) (t time.Time, err error) {
	return time.ParseInLocation(layoutTime, value, LocationChina)
}