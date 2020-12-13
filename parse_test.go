package xtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func mustTime(t time.Time, err error) time.Time {
	if err != nil {panic(err)}
	return t
}
func TestParseChinaYear(t *testing.T) {
	assert.Equal(t, time.Date(2021,1,1,0,0,0,0, LocationChina), mustTime(ParseChinaYear("2021")))
}
func TestParseChinaYearAndMonth(t *testing.T) {
	assert.Equal(t,  time.Date(2021,1,1,0,0,0,0, LocationChina), mustTime(ParseChinaYearAndMonth("2021-01")))
}
func TestParseChinaDate(t *testing.T) {
	assert.Equal(t,  time.Date(2021,01,01,0,0,0,0, LocationChina), mustTime(ParseChinaDate("2021-01-01")))
}
func TestParseChinaTime(t *testing.T) {
	assert.Equal(t,  time.Date(2021,01,01,7,23,23,0, LocationChina), mustTime(ParseChinaTime("2021-01-01 07:23:23")))
}
