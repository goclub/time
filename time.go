package xtime

import (
	"log"
	"time"
)

type InRangeTimeData struct {
	Begin  time.Time
	End    time.Time
	Target time.Time
}
func InRangeTime(d InRangeTimeData) (in bool) {
	begin := d.Begin
	end := d.End
	if d.Begin.After(d.End) {
		begin = d.End
		end = d.Begin
		log.Print("goclub/time: InRangeTime() fixed begin and end")
	}
	// begin <= target <= end -> true
	if ( begin.Before(d.Target) || begin.Equal(d.Target) ) &&
		( d.Target.Before(end) || d.Target.Equal(end) ) {
		in = true
		return
	}
	return
}

// func inRangeDate(begin string, end string, target string) bool
