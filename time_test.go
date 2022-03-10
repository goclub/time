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
		Time []time.Time // begin, end, target
		In bool
	}
	/*
		a begin=end
		b begin<end
		c begin>end

		1 in 在范围内
		2 in left 在范围内, 并等于开始时间
		3 in right 在范围内, 并等于结束时间
		4 F in left 不在范围内, 并小于开始时间
		5 F in right 不在范围内, 并大于结束时间

		共13种情况
		a 1,4,5
		b 1,2,3,4,5
		c 1,2,3,4,5
	*/
	dataList := []Date{
		// a 1
		{ []time.Time{time2, time2, time2}, true },
		// a 4
		{ []time.Time{time2, time2, time1}, false },
		// a 5
		{ []time.Time{time3, time3, time5}, false },
		// b 1
		{ []time.Time{time2, time4, time3}, true },
		// b 2
		{ []time.Time{time2, time4, time2}, true },
		// b 3
		{ []time.Time{time2, time4, time4}, true },
		// b 4
		{ []time.Time{time2, time4, time1}, false },
		// b 5
		{ []time.Time{time2, time4, time5}, false },
		// c 1
		{ []time.Time{time4, time2, time3}, true },
		// c 2
		{ []time.Time{time4, time2, time2}, true },
		// c 3
		{ []time.Time{time4, time2, time4}, true },
		//  c 4
		{ []time.Time{time4, time2, time1}, false },
		// c 5
		{ []time.Time{time4, time2, time5}, false },
	}

	for k, v := range dataList {
		in := xtime.InRangeTime(xtime.InRangeTimeData{
			Begin:  v.Time[0],
			End:    v.Time[1],
			Target: v.Time[2],
		})
		assert.Equal(t, v.In, in, k+1)
	}
}
