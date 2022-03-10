package xtime

import (
	"log"
	"time"
)

type Range struct {
	Begin  time.Time
	End    time.Time
}
func InRange(target time.Time, r Range) (in bool) {
	begin := r.Begin
	end := r.End
	if r.Begin.After(r.End) {
		begin = r.End
		end = r.Begin
		log.Print("goclub/time: InRange(target time.Time, r Range) r.Begin can not  be after r.End, InRange() already replacement they")
	}
	// begin <= target <= end -> true
	if ( begin.Before(target) || begin.Equal(target) ) &&
		( target.Before(end) || target.Equal(end) ) {
		in = true
		return
	}
	return
}

// func inRangeDate(begin string, end string, target string) bool
