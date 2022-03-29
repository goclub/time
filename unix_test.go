package xtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUnixMilli(t *testing.T) {
	{
		sometime := time.Date(2020,12,31,23,23,23, 0, LocChina)
		assert.Equal(t, UnixMilli(sometime), int64(1609428203000), )
	}
	{
		sometime := time.Date(2020,12,31,23,23,23, 333333333, LocChina)
		assert.Equal(t, UnixMilli(sometime), int64(1609428203333), )
	}
}