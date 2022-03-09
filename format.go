package xtime

import "time"
func FormatChinaYear(t time.Time) string {
	return t.In(LocationChina).Format(LayoutYear)
}
func FormatChinaYearAndMonth(t time.Time) string {
	return t.In(LocationChina).Format(LayoutYearAndMonth)
}
func FormatChinaTime(t time.Time) string {
	return t.In(LocationChina).Format(LayoutTime)
}
func FormatChinaDate(t time.Time) string {
	return t.In(LocationChina).Format(LayoutDate)
}
func FormatChinaHourMinuteSecond(t time.Time) string {
	return t.In(LocationChina).Format(LayoutHourMinuteSecond)
}