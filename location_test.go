package xtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLocation(t *testing.T) {
	sometime := time.Date(2020,11,11,11,11,11, 0,time.UTC)
	assert.Equal(t, sometime.String(), "2020-11-11 11:11:11 +0000 UTC")
	sometime = sometime.In(LocationChina)
	assert.Equal(t, sometime.String(), "2020-11-11 19:11:11 +0800 CST")
}