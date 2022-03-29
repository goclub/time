package xtime

import "time"
func FormatChinaYear(t time.Time) string {
	return t.In(LocChina).Format(LayoutYear)
}
func FormatChinaYearAndMonth(t time.Time) string {
	return t.In(LocChina).Format(LayoutYearAndMonth)
}
func FormatChinaTime(t time.Time) string {
	return t.In(LocChina).Format(LayoutTime)
}
func FormatChinaDate(t time.Time) string {
	return t.In(LocChina).Format(LayoutDate)
}
func FormatChinaHourMinuteSecond(t time.Time) string {
	return t.In(LocChina).Format(LayoutHourMinuteSecond)
}