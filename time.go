package xtime

import (
	"database/sql/driver"
	"fmt"
	xerr "github.com/goclub/error"
	"log"
	"math"
	"strconv"
	"time"
)

type Range struct {
	Start time.Time
	End   time.Time
}

func InRange(target time.Time, r Range) (in bool) {
	begin := r.Start
	end := r.End
	if r.Start.After(r.End) {
		begin = r.End
		end = r.Start
		log.Print("goclub/time: InRange(target time.Time, r Range) r.Start can not  be after r.End, InRange() already replacement they")
	}
	// begin <= target <= end -> true
	if (begin.Before(target) || begin.Equal(target)) &&
		(target.Before(end) || target.Equal(end)) {
		in = true
		return
	}
	return
}

type DateRange struct {
	Begin Date `note:"当日期是 2022-01-01 时等同于 Range{Start: 2022-01-01 00:00:00}"`
	End   Date `note:"当日期是 2022-01-03 时等同于 Range{End: 2022-01-03 23:59:59}"`
}

func (r DateRange) Validator(err ...error) error {
	if r.Begin.After(r.End) {
		return xerr.New("goclub/time: DateRange.Begin can not be after DateRange.")
	}
	return nil
}

func InRangeFromDate(target time.Time, r DateRange) (in bool) {
	timeRange := Range{
		Start: FirstSecondOfDate(r.Begin.Time(target.Location())),
		End:   LastSecondOfDate(r.End.Time(target.Location())),
	}
	return InRange(target, timeRange)
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func Today(loc *time.Location) Date {
	year, month, day := time.Now().In(loc).Date()
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}
func NewDate(year int, month time.Month, day int) Date {
	return Date{
		year, month, day,
	}
}
func NewDateFromTime(t time.Time) Date {
	return NewDate(t.Date())
}
func NewDateFromString(value string) (d Date, err error) {
	t, err := time.Parse(LayoutDate, value)
	if err != nil {
		err = xerr.WithStack(err)
		return
	}
	return NewDateFromTime(t), nil
}
func (d Date) IsZero() bool {
	if d.Year == 0 && d.Month == 0 && d.Day == 0 {
		return true
	}
	return false
}

func (d *Date) MarshalRequest(value string) error {
	newDate, err := NewDateFromString(value) // indivisible begin
	if err != nil {                          // indivisible end
		return err
	}
	d = &newDate
	return nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	value, err := strconv.Unquote(string(b))
	if err != nil {
		return xerr.WithStack(err)
	}
	date, err := NewDateFromString(value)
	if err != nil {
		return xerr.WithStack(err)
	}
	*d = date
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}
func (d Date) Time(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
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
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC).Format(LayoutDate)
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
		date, err = NewDateFromString(fmt.Sprintf("%s", value))
		if err != nil {
			return
		}
		*d = date
	}
	return
}

func (d Date) AddDate(years int, months int, days int) (date Date) {
	return NewDateFromTime(d.UTCTime().AddDate(years, months, days))
}

func (d Date) Sub(u Date) (days int) {
	return int(math.Round(d.UTCTime().Sub(u.UTCTime()).Hours())) / 24
}

// FirstDateOfMonth 2022-11-11 => 2022-11-01
func (d Date) FirstDateOfMonth() (first Date) {
	return NewDate(d.Year, d.Month, 1)
}

// LastDateOfMonth 2022-11-11 => 2022-11-30
func (d Date) LastDateOfMonth() (first Date) {
	return d.FirstDateOfMonth().AddDate(0, 1, -1)
}
func (d Date) After(t Date) bool {
	return d.UTCTime().After(t.UTCTime())
}
func (d Date) Before(t Date) bool {
	return d.UTCTime().Before(t.UTCTime())
}
func (d Date) Equal(t Date) bool {
	return d == t
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
	err := date.UnmarshalJSON(b)
	if err != nil {
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

// LastSecondOfDate 2022-11-11 xx:xx:xx => 2022-11-11 23:59:59.999
func LastSecondOfDate(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}

// FirstSecondOfDate 2022-11-11 xx:xx:xx => 2022-11-11 00:00:00
func FirstSecondOfDate(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// TomorrowFirstSecond 2022-11-11 xx:xx:xx => 2022-11-12 00:00:00
func TomorrowFirstSecond(t time.Time) time.Time {
	return FirstSecondOfDate(t.AddDate(0, 0, 1))
}

// TomorrowFirstSecondDuration  2022-11-11 23:59:50   => 10s
func TomorrowFirstSecondDuration(t time.Time) time.Duration {
	return TomorrowFirstSecond(t).Sub(t)
}

func Now(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

func SplitRange(days uint, r DateRange) (splitRanges []DateRange) {
	if days == 0 {
		days = 1
	}
	splitRanges = []DateRange{}
	// 边界: 2022-01-01~2022-01-01
	if r.Begin.Equal(r.End) {
		splitRanges = append(splitRanges, r)
		return
	}
	// 边界: 2022-01-02~2022-01-01
	if err := r.Validator(); err != nil {
		// 格式错误必须 panic 否则即使当前逻辑不出错后续逻辑也会出错
		panic(err)
	}
	slow := r.Begin
	for {
		itemEnd := slow.AddDate(0, 0, int(days))
		if itemEnd.Before(r.End) {
			splitRanges = append(splitRanges, DateRange{slow, itemEnd})
			slow = itemEnd.AddDate(0, 0, 1)
		} else {
			splitRanges = append(splitRanges, DateRange{slow, r.End})
			break
		}
	}
	return splitRanges
}
