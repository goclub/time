package xtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatChinaYear(t *testing.T) {
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	assert.Equal(t, FormatChinaYear(sometime), "2021")
}
func TestFormatChinaYearAndMonth(t *testing.T) {
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	assert.Equal(t, FormatChinaYearAndMonth(sometime), "2021-01")
}
func TestFormatChinaDate(t *testing.T) {
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	assert.Equal(t, FormatChinaDate(sometime), "2021-01-01")
}
func TestFormatChinaTime(t *testing.T) {
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	assert.Equal(t, FormatChinaTime(sometime), "2021-01-01 07:23:23")
}
func TestFormatChinaHourMinuteSecond(t *testing.T) {
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	assert.Equal(t, FormatChinaHourMinuteSecond(sometime), "07:23:23")
}
