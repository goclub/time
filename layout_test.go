package xtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLayout(t *testing.T) {
	assert.Equal(t, LayoutYear, "2006")
	assert.Equal(t, LayoutYearAndMonth, "2006-01")
	assert.Equal(t, LayoutDate, "2006-01-02")
	assert.Equal(t, LayoutTime, "2006-01-02 15:04:05")
	assert.Equal(t, LayoutHourMinuteSecond, "15:04:05")
}