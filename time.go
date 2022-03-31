package xtime

import (
	"database/sql/driver"
	"fmt"
	xerr "github.com/goclub/error"
	"log"
	"strconv"
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
type DateRange struct {
	Begin  Date `note:"当日期是 2022-01-01 时等同于 Range{Begin: 2022-01-01 00:00:00}"`
	End    Date `note:"当日期是 2022-01-03 时等同于 Range{End: 2022-01-03 23:59:59}"`
}
func InRangeFromDate(target time.Time, r DateRange) (in bool) {
	timeRange := Range{
		Begin: FirstSecondOfDate(r.Begin.Time(target.Location())),
		End:   LastSecondOfDate(r.End.Time(target.Location())),
	}
	return InRange(target, timeRange)
}
type Date struct {
	Year int
	Month time.Month
	Day int
}
func NowDate(loc *time.Location) Date {
	year, month, day := time.Now().In(loc).Date()
	return Date{
		Year: year,
		Month: month,
		Day: day,
	}
}
func NewDate(year int, month time.Month, day int) Date {
	return Date{
		year, month,day,
	}
}
func NewDateFromTime(t time.Time) Date {
	return NewDate(t.Date())
}
func NewDateFromString(value string) (d Date, err error) {
	t, err := time.Parse(LayoutDate, value) ; if err != nil {
		err = xerr.WithStack(err)
	    return
	}
	return NewDateFromTime(t), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	value, err := strconv.Unquote(string(b)) ; if err != nil {
		return xerr.WithStack(err)
	}
	date, err := NewDateFromString(value) ; if err != nil {
		return xerr.WithStack(err)
	}
	*d = date
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}
func (d Date) Time(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0,0,0,0, loc)
}
func (d Date) LocalTime() time.Time {
	return d.Time(time.Local)
}
func (d Date) UTCTime() time.Time {
	return d.Time(time.UTC)
}
func (d Date) ChinaTime() ChinaTime {
	return NewChinaTime(d.Time(LocChina))
}
func (d Date) String() string {
	return time.Date(d.Year, d.Month, d.Day, 0,0,0,0, time.UTC).Format(LayoutDate)
}
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}
func (d *Date) Scan(value interface{}) (err error) {
	if value == nil {
		return xerr.New("unsupported NULL xtime.Date value, maybe you should use xtime.NullDate")
	}
	switch v := value.(type) {
	case time.Time:
		*d = NewDateFromTime(v)
	default:
		var date Date
		date, err = NewDateFromString(fmt.Sprintf("%s", value)) ; if err != nil {
		return
	}
		*d = date
	}
	return
}

type NullDate struct {
	date  Date
	valid bool
}

func (v NullDate) Date() Date {
	return v.date
}
func (v NullDate) Valid() bool {
	return v.valid
}
func NewNullDate(year int, month time.Month, day int) NullDate {
	return NullDate{
		date:  NewDate(year, month, day),
		valid: true,
	}
}

func (d *NullDate) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == "" {
		d.valid = false
		return nil
	}
	date := Date{}
	err := date.UnmarshalJSON(b) ; if err != nil {
	    return xerr.WithStack(err)
	}
	*d = NullDate{
		date:  date,
		valid: true,
	}
	return err
}

func (d NullDate) MarshalJSON() ([]byte, error) {
	if d.valid {
		return d.date.MarshalJSON()
	} else {
		return []byte(`null`), nil
	}
}
func (d NullDate) String() string {
	if d.valid {
		return d.date.String()
	}
	return "null"
}

func (d NullDate) Value() (driver.Value, error) {
	if d.valid {
		return d.date.Value()
	}
	return nil, nil
}
func (d *NullDate) Scan(value interface{}) error {
	if value == nil {
		d.valid = false
		return nil
	}
	d.valid = true
	return d.date.Scan(value)
}

func LastSecondOfDate(t time.Time) time.Time {
	y,m,d := t.Date()
	return time.Date(y,m,d,23,59,59,999999999,t.Location())
}
func FirstSecondOfDate(t time.Time) time.Time {
	y,m,d := t.Date()
	return time.Date(y,m,d,0,0,0,0,t.Location())
}
