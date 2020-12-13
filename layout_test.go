package xtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLayout(t *testing.T) {
	assert.Equal(t, layoutYear, "2006")
	assert.Equal(t, layoutYearAndMonth, "2006-01")
	assert.Equal(t, layoutDate, "2006-01-02")
	assert.Equal(t, layoutTime, "2006-01-02 15:04:05")
	assert.Equal(t, layoutHourMinuteSecond, "15:04:05")
}