package xtime_test

import (
	xtime "github.com/goclub/time"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInRangeTime(t *testing.T) {
	time1 := time.Now()
	time2 := time1.Add(1*time.Second)
	time3 := time1.Add(2*time.Second)
	time4 := time1.Add(3*time.Second)
	time5 := time1.Add(4*time.Second)

	type Date struct {
		xtime.InRangeTimeData
		In bool
	}
	/*
		1 in 在范围内
		2 F in left 不在范围内, 并小于开始时间
		3 F in right 不在范围内, 并大于结束时间

		4 begin=end
		5 begin<end
		6 begin>end
	*/
	dataList := []Date{
		// 1 in 4 begin=end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time2,
				End:    time2,
				Target: time2,
			},
			In:              true,
		},
		// 1 in 5 begin<end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time3,
			},
			In:              true,
		},
		// 1 in 6 begin>end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time4,
				End:    time2,
				Target: time3,
			},
			In:              true,
		},
		// 2 F in left 4 begin=end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time2,
				End:    time2,
				Target: time1,
			},
			In:              false,
		},
		// 2 F in left 5 begin<end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time1,
			},
			In:              false,
		},
		// 2 F in left 6 begin>end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time4,
				End:    time2,
				Target: time1,
			},
			In:              false,
		},
		// 3 F in right 4 begin=end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time3,
				End:    time3,
				Target: time5,
			},
			In:              false,
		},
		// 3 F in right 5 begin<end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time5,
			},
			In:              false,
		},
		// 3 F in right 6 begin>end
		{
			InRangeTimeData: xtime.InRangeTimeData{
				Begin:  time4,
				End:    time2,
				Target: time5,
			},
			In:              false,
		},
	}

	for k, v := range dataList {
		in := xtime.InRangeTime(v.InRangeTimeData)
		assert.Equal(t, v.In, in, k+1)
	}
}
