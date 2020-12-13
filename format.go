package xtime

import "time"
func FormatChinaYear(t time.Time) string {
	return t.In(LocationChina).Format(layoutYear)
}
func FormatChinaYearAndMonth(t time.Time) string {
	return t.In(LocationChina).Format(layoutYearAndMonth)
}
func FormatChinaTime(t time.Time) string {
	return t.In(LocationChina).Format(layoutTime)
}
func FormatChinaDate(t time.Time) string {
	return t.In(LocationChina).Format(layoutDate)
}
func FormatChinaHourMinuteSecond(t time.Time) string {
	return t.In(LocationChina).Format(layoutHourMinuteSecond)
}