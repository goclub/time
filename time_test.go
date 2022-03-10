package xtime_test

import (
	xtime "github.com/goclub/time"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInRangeTime2(t *testing.T) {

	// 2000-01-01 00:00:01
	t1 := time.Date(2000,1,1, 0,0,1,0, time.Local)
	t2 := t1.Add(1 * time.Second)
	t3 := t1.Add(2 * time.Second)
	t4 := t1.Add(3 * time.Second)
	t5 := t1.Add(4 * time.Second)
	type Date struct {
		Time []time.Time
		In bool
	}
	data := []Date{
		// begin target end
		// 2 * 2
		{ []time.Time{t2, t1, t2}, false, },
		{ []time.Time{t2, t2, t2}, true, },
		{ []time.Time{t2, t3, t2}, false, },
		// 2 * 3
		{ []time.Time{t2, t1, t3}, false, },
		{ []time.Time{t2, t2, t3}, true, },
		{ []time.Time{t2, t3, t3}, true, },
		// 2 * 4
		{ []time.Time{t2, t1, t4}, false, },
		{ []time.Time{t2, t2, t4}, true, },
		{ []time.Time{t2, t3, t4}, true, },
		{ []time.Time{t2, t4, t4}, true, },
		{ []time.Time{t2, t5, t4}, false, },

		// replacement begin & end
		{ []time.Time{t3, t2, t1}, true, },
		{ []time.Time{t3, t5, t1}, false, },
	}
	for _, item := range data {
		begin := item.Time[0]
		target := item.Time[1]
		end := item.Time[2]
		in := xtime.InRange(target, xtime.Range{
			Begin:  begin,
			End:    end,
		})
		assert.Equalf(t, in, item.In, "[]time.Time{t%d, t%d, t%d}, %v, }", begin.Second(), target.Second(), end.Second(), item.In)
	}
}
func TestInRangeTime(t *testing.T) {
	time1 := time.Now()
	time2 := time1.Add(1*time.Second)
	time3 := time1.Add(2*time.Second)
	time4 := time1.Add(3*time.Second)
	time5 := time1.Add(4*time.Second)

	type InRangeTimeData struct {
		Begin time.Time
		End time.Time
		Target time.Time
	}
	type Date struct {
		InRangeTimeData
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
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time2,
				Target: time2,
			},
			In:              true,
		},
		// 1 in 5 begin<end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time3,
			},
			In:              true,
		},
		// 1 in 6 begin>end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time4,
				End:    time2,
				Target: time3,
			},
			In:              true,
		},
		// 2 F in left 4 begin=end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time2,
				Target: time1,
			},
			In:              false,
		},
		// 2 F in left 5 begin<end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time1,
			},
			In:              false,
		},
		// 2 F in left 6 begin>end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time4,
				End:    time2,
				Target: time1,
			},
			In:              false,
		},
		// 3 F in right 4 begin=end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time3,
				End:    time3,
				Target: time5,
			},
			In:              false,
		},
		// 3 F in right 5 begin<end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time5,
			},
			In:              false,
		},
		// 3 F in right 6 begin>end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time4,
				End:    time2,
				Target: time5,
			},
			In:              false,
		},
	}

	for k, v := range dataList {
		in := xtime.InRange(v.Target, xtime.Range{
			Begin: v.Begin,
			End: v.End,
		})
		assert.Equal(t, v.In, in, k+1)
	}
}
