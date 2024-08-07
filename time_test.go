package xtime_test

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	xtime "github.com/goclub/time"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestInRangeTime2(t *testing.T) {

	// 2000-01-01 00:00:01
	t1 := time.Date(2000, 1, 1, 0, 0, 1, 0, time.Local)
	t2 := t1.Add(1 * time.Second)
	_ = t2
	t3 := t1.Add(2 * time.Second)
	_ = t3
	t4 := t1.Add(3 * time.Second)
	_ = t4
	t5 := t1.Add(4 * time.Second)
	_ = t5
	type Date struct {
		Time []time.Time
		In   bool
	}
	data := []Date{
		// begin target end
		// 2 * 2
		{[]time.Time{t2, t1, t2}, false},
		{[]time.Time{t2, t2, t2}, true},
		{[]time.Time{t2, t3, t2}, false},
		// 2 * 3
		{[]time.Time{t2, t1, t3}, false},
		{[]time.Time{t2, t2, t3}, true},
		{[]time.Time{t2, t3, t3}, true},
		// 2 * 4
		{[]time.Time{t2, t1, t4}, false},
		{[]time.Time{t2, t2, t4}, true},
		{[]time.Time{t2, t3, t4}, true},
		{[]time.Time{t2, t4, t4}, true},
		{[]time.Time{t2, t5, t4}, false},
	}
	for _, item := range data {
		begin := item.Time[0]
		target := item.Time[1]
		end := item.Time[2]
		in := xtime.InRange(target, xtime.Range{
			Start: begin,
			End:   end,
		})
		assert.Equalf(t, in, item.In, "Time: []time.Time{t%d, t%d, t%d}, %v, }", begin.Second(), target.Second(), end.Second(), item.In)
	}
}

func TestInDateRangeTime(t *testing.T) {
	// 2000-01-01
	t1 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	t2 := t1.Add(1 * 24 * time.Hour)
	_ = t2
	t3 := t1.Add(2 * 24 * time.Hour)
	_ = t3
	t4 := t1.Add(3 * 24 * time.Hour)
	_ = t4
	t5 := t1.Add(4 * 24 * time.Hour)
	_ = t5
	type Date struct {
		Time []time.Time
		In   bool
	}
	data := []Date{
		// begin target end
		// 2 * 2
		{[]time.Time{t2, t1, t2}, false},
		{[]time.Time{t2, t2, t2}, true},
		{[]time.Time{t2, t3, t2}, false},
		// 2 * 3
		{[]time.Time{t2, t1, t3}, false},
		{[]time.Time{t2, t2, t3}, true},
		{[]time.Time{t2, t3, t3}, true},
		// 2 * 4
		{[]time.Time{t2, t1, t4}, false},
		{[]time.Time{t2, t2, t4}, true},
		{[]time.Time{t2, t3, t4}, true},
		{[]time.Time{t2, t4, t4}, true},
		{[]time.Time{t2, t5, t4}, false},
	}
	for _, item := range data {
		begin := item.Time[0]
		target := item.Time[1]
		end := item.Time[2]
		in := xtime.InRangeFromDate(target, xtime.DateRange{
			Start: xtime.NewDateFromTime(begin),
			End:   xtime.NewDateFromTime(end),
		})
		assert.Equalf(t, in, item.In, "Date: []time.Time{t%d, t%d, t%d}, %v, }", begin.Second(), target.Second(), end.Second(), item.In)
	}
}
func TestInRangeTime(t *testing.T) {
	time1 := time.Now()
	time2 := time1.Add(1 * time.Second)
	time3 := time1.Add(2 * time.Second)
	time4 := time1.Add(3 * time.Second)
	time5 := time1.Add(4 * time.Second)

	type InRangeTimeData struct {
		Begin  time.Time
		End    time.Time
		Target time.Time
	}
	type Date struct {
		InRangeTimeData
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
		// 1 in 4 begin=end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time2,
				Target: time2,
			},
			In: true,
		},
		// 1 in 5 begin<end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time3,
			},
			In: true,
		},
		// 1 in 6 begin>end
		//{
		//	InRangeTimeData: InRangeTimeData{
		//		Begin:  time4,
		//		End:    time2,
		//		Target: time3,
		//	},
		//	In: true,
		//},
		// 2 F in left 4 begin=end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time2,
				Target: time1,
			},
			In: false,
		},
		// 2 F in left 5 begin<end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time1,
			},
			In: false,
		},
		// 2 F in left 6 begin>end
		//{
		//	InRangeTimeData: InRangeTimeData{
		//		Begin:  time4,
		//		End:    time2,
		//		Target: time1,
		//	},
		//	In: false,
		//},
		// 3 F in right 4 begin=end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time3,
				End:    time3,
				Target: time5,
			},
			In: false,
		},
		// 3 F in right 5 begin<end
		{
			InRangeTimeData: InRangeTimeData{
				Begin:  time2,
				End:    time4,
				Target: time5,
			},
			In: false,
		},
		// 3 F in right 6 begin>end
		//{
		//	InRangeTimeData: InRangeTimeData{
		//		Begin:  time4,
		//		End:    time2,
		//		Target: time5,
		//	},
		//	In: false,
		//},
	}

	for k, v := range dataList {
		in := xtime.InRange(v.Target, xtime.Range{
			Start: v.Begin,
			End:   v.End,
		})
		assert.Equal(t, v.In, in, k+1)
	}
}

func TestDateSQL(t *testing.T) {
	func() struct{} {

		db, err := sql.Open("mysql", "root:somepass@(127.0.0.1:3306)/goclub_time?"+url.Values{
			"charset":   []string{"utf8"},
			"parseTime": []string{"True"},
			"loc":       []string{"Local"},
		}.Encode())
		if err != nil {
			assert.NoError(t, err)
		}
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS date (
		  id int(11) unsigned NOT NULL AUTO_INCREMENT,
		  date date NOT NULL,
		  null_date date DEFAULT NULL,
		  PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4;`)
		assert.NoError(t, err)

		// insert 2022-01-01 null
		{
			sql := "INSERT INTO `date` (`date`, `null_date`) VALUES (?, ?)"
			result, err := db.Exec(sql, xtime.NewDate(2022, 01, 01), xtime.NullDate{})
			if err != nil {
				assert.NoError(t, err)
			}
			id, err := result.LastInsertId()
			if err != nil {
				assert.NoError(t, err)
			}
			sql = "SELECT `date` ,`null_date` FROM `date` WHERE `id` = ?"
			row := db.QueryRow(sql, id)
			value := xtime.Date{}
			nullValue := xtime.NullDate{}
			err = row.Scan(&value, &nullValue)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, value.String(), "2022-01-01")
			assert.Equal(t, nullValue, xtime.NullDate{})
			{
				row := db.QueryRow("SELECT `null_date` FROM `date` WHERE `id` = ?", id)
				value := xtime.Date{}
				err = row.Scan(&value)
				assert.Equal(t, err.Error(), `sql: Scan error on column index 0, name "null_date": unsupported NULL xtime.Date value, maybe you should use xtime.NullDate`)
			}
		}
		// insert 2022-01-01 2022-01-01
		{
			result, err := db.Exec("INSERT INTO `date` (`date`, `null_date`) VALUES (?, ?)", xtime.NewDate(2022, 01, 01), xtime.NewNullDate(2022, 01, 02))
			if err != nil {
				assert.NoError(t, err)
			}
			id, err := result.LastInsertId()
			if err != nil {
				assert.NoError(t, err)
			}
			row := db.QueryRow("SELECT `date` ,`null_date` FROM `date` WHERE `id` = ?", id)
			value := xtime.Date{}
			nullValue := xtime.NullDate{}
			err = row.Scan(&value, &nullValue)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, value.String(), "2022-01-01")
			assert.Equal(t, nullValue, xtime.NewNullDate(2022, 01, 02))
		}
		return struct{}{}
	}()
}
func TestNowDate(t *testing.T) {
	date := xtime.Today(xtime.LocChina)
	y, m, d := time.Now().In(xtime.LocChina).Date()
	assert.Equal(t, date, xtime.NewDate(y, m, d))
}
func TestNewDate(t *testing.T) {
	date := xtime.NewDate(2022, 01, 01)
	assert.Equal(t, date, xtime.Date{2022, 01, 01})
	{
		date := xtime.NewNullDate(2022, 01, 01)
		assert.Equal(t, date, xtime.NewNullDate(2022, 01, 01))
		assert.Equal(t, date.Date(), xtime.Date{2022, 01, 01})
		assert.Equal(t, date.Valid(), true)
		assert.Equal(t, xtime.NullDate{}.Date(), xtime.Date{})
		assert.Equal(t, xtime.NullDate{}.Valid(), false)
	}
}

func TestNewDateFromTime(t *testing.T) {
	date := xtime.NewDateFromTime(time.Date(2022, 01, 01, 0, 0, 0, 0, xtime.LocChina))
	assert.Equal(t, date, xtime.Date{2022, 01, 01})
}
func TestParseDate(t *testing.T) {
	date, err := xtime.NewDateFromString("2022-01-01")
	if err != nil {
		return
	}
	assert.Equal(t, date, xtime.Date{2022, 01, 01})
	_, err = xtime.NewDateFromString("2022-01-0")
	assert.Errorf(t, err, `parsing time "2022-01-0" as "2006-01-02": cannot parse "0" as "02"`)
}

func TestDate_UnmarshalJSON(t *testing.T) {
	func() struct{} {
		{
			v := struct {
				Date xtime.Date `json:"date"`
			}{}
			err := json.Unmarshal([]byte(`{"date":"2022-11-11"}`), &v)
			assert.NoError(t, err)
			assert.Equal(t, v.Date, xtime.Date{2022, 11, 11})
		}
		{
			v := struct {
				Date xtime.Date `json:"date"`
			}{}
			err := json.Unmarshal([]byte(`{"date":"2022-11-1"}`), &v)
			assert.Errorf(t, err, `parsing time "2022-11-1" as "2006-01-02": cannot parse "1" as "02"`)
		}
		return struct{}{}
	}()
}
func TestDate_MarshalJSON(t *testing.T) {
	v := struct {
		Date xtime.Date `json:"date"`
	}{
		Date: xtime.NewDate(2022, 11, 11),
	}
	data, err := json.Marshal(v)
	assert.NoError(t, err)
	assert.Equal(t, string(data), `{"date":"2022-11-11"}`)
}
func TestDate_Time(t *testing.T) {
	date := xtime.NewDate(2022, 11, 11)
	dateTime := date.Time(xtime.LocChina)
	assert.Equal(t, dateTime, time.Date(2022, 11, 11, 0, 0, 0, 0, xtime.LocChina))
}
func TestDate_LocalTime(t *testing.T) {
	date := xtime.NewDate(2022, 11, 11)
	dateTime := date.LocalTime()
	assert.Equal(t, dateTime, time.Date(2022, 11, 11, 0, 0, 0, 0, time.Local))
}
func TestDate_UTCTime(t *testing.T) {
	date := xtime.NewDate(2022, 11, 11)
	dateTime := date.UTCTime()
	assert.Equal(t, dateTime, time.Date(2022, 11, 11, 0, 0, 0, 0, time.UTC))
}
func TestDate_ChinaTime(t *testing.T) {
	date := xtime.NewDate(2022, 11, 11)
	dateTime := date.ChinaTime()
	assert.Equal(t, dateTime, xtime.NewChinaTime(time.Date(2022, 11, 11, 0, 0, 0, 0, xtime.LocChina)))
}
func TestDate_String(t *testing.T) {
	date := xtime.NewDate(2022, 11, 11)
	assert.Equal(t, date.String(), "2022-11-11")
}
func TestDate_Value(t *testing.T) {
	date := xtime.NewDate(2022, 11, 11)
	v, err := date.Value()
	assert.NoError(t, err)
	assert.Equal(t, v, "2022-11-11")
}
func TestNullDate_MarshalJSON(t *testing.T) {
	func() struct{} {
		{
			v := struct {
				Date xtime.NullDate `json:"date"`
			}{
				Date: xtime.NewNullDate(2022, 11, 11),
			}
			data, err := json.Marshal(v)
			assert.NoError(t, err)
			assert.Equal(t, string(data), `{"date":"2022-11-11"}`)
		}
		{
			v := struct {
				Date xtime.NullDate `json:"date"`
			}{
				Date: xtime.NullDate{},
			}
			data, err := json.Marshal(v)
			assert.NoError(t, err)
			assert.Equal(t, string(data), `{"date":null}`)
		}
		return struct{}{}
	}()
}
func TestNullDate_UnmarshalJSON(t *testing.T) {
	func() struct{} {
		{
			v := struct {
				Date xtime.NullDate `json:"date"`
			}{}
			err := json.Unmarshal([]byte(`{"date":null}`), &v)
			assert.NoError(t, err)
			assert.Equal(t, v.Date, xtime.NullDate{})
		}
		{
			v := struct {
				Date xtime.NullDate `json:"date"`
			}{}
			err := json.Unmarshal([]byte(`{"date":"2022-01-01"}`), &v)
			assert.NoError(t, err)
			assert.Equal(t, v.Date, xtime.NewNullDate(2022, 01, 01))
		}
		return struct{}{}
	}()

}
func TestNullDate_String(t *testing.T) {
	assert.Equal(t, xtime.NewNullDate(2022, 01, 01).String(), "2022-01-01")
	assert.Equal(t, xtime.NullDate{}.String(), "null")
}

func TestFirstSecondOfDate(t *testing.T) {
	func() struct{} {
		sometime := time.Date(2022, 01, 01, 12, 12, 12, 88, xtime.LocChina)
		assert.Equal(t, xtime.FirstSecondOfDate(sometime), time.Date(2022, 01, 01, 0, 0, 0, 0, xtime.LocChina))
		assert.Equal(t, xtime.LastSecondOfDate(sometime), time.Date(2022, 01, 01, 23, 59, 59, 999999999, xtime.LocChina))
		return struct{}{}
	}()
}
func TestTomorrowFirstSecond(t *testing.T) {
	func() struct{} {
		// -------------
		sometime := time.Date(2022, 01, 01, 12, 12, 12, 88, xtime.LocChina)
		assert.Equal(t, xtime.UnixMilli(xtime.TomorrowFirstSecond(sometime)), int64(1641052800000))
		assert.Equal(t, xtime.TomorrowFirstSecond(sometime).String(), "2022-01-02 00:00:00 +0800 CST")
		// -------------
		return struct{}{}
	}()
}
func TestTomorrowFirstSecondDuration(t *testing.T) {
	func() struct{} {
		// -------------
		sometime := time.Date(2022, 11, 11, 23, 59, 50, 00, xtime.LocChina)
		assert.Equal(t, xtime.TomorrowFirstSecondDuration(sometime), time.Second*10)
		// -------------
		return struct{}{}
	}()
}

func TestDate_AddDate(t *testing.T) {
	func() struct{} {
		// -------------
		someDate := xtime.NewDate(2022, 10, 01)
		assert.Equal(t, someDate.AddDate(1, 0, 0).String(), "2023-10-01")
		assert.Equal(t, someDate.AddDate(0, 1, 0).String(), "2022-11-01")
		assert.Equal(t, someDate.AddDate(0, 0, 1).String(), "2022-10-02")

		assert.Equal(t, someDate.AddDate(-1, 0, 0).String(), "2021-10-01")
		assert.Equal(t, someDate.AddDate(0, -1, 0).String(), "2022-09-01")
		assert.Equal(t, someDate.AddDate(0, 0, -1).String(), "2022-09-30")

		assert.Equal(t, someDate.AddDate(1, 1, 0).String(), "2023-11-01")
		assert.Equal(t, someDate.AddDate(1, 1, 1).String(), "2023-11-02")

		assert.Equal(t, someDate.AddDate(-1, -1, 0).String(), "2021-09-01")
		assert.Equal(t, someDate.AddDate(-1, -1, -1).String(), "2021-08-31")

		// -------------
		return struct{}{}
	}()
}

func TestDate_Sub(t *testing.T) {
	func() struct{} {
		// -------------
		someDate := xtime.NewDate(2022, 10, 01)
		assert.Equal(t, someDate.Sub(xtime.NewDate(2022, 10, 10)), int(-9))
		assert.Equal(t, someDate.Sub(xtime.NewDate(2022, 9, 30)), int(1))
		assert.Equal(t, someDate.Sub(xtime.NewDate(2023, 10, 01)), int(-365))
		// -------------
		return struct{}{}
	}()
}

func TestDate_FirstDateOfMonth(t *testing.T) {
	func() struct{} {
		// -------------
		var err error
		_ = err
		ctx := context.Background()
		_ = ctx
		{
			someDate := xtime.NewDate(2022, 11, 11)
			assert.Equal(t, someDate.FirstDateOfMonth(), xtime.NewDate(2022, 11, 01))
		}
		{
			someDate := xtime.NewDate(2022, 11, 01)
			assert.Equal(t, someDate.FirstDateOfMonth(), xtime.NewDate(2022, 11, 01))
		}
		{
			someDate := xtime.NewDate(2022, 11, 30)
			assert.Equal(t, someDate.FirstDateOfMonth(), xtime.NewDate(2022, 11, 01))
		}

		// -------------
		return struct{}{}
	}()
}

func TestDate_LastDateOfMonth(t *testing.T) {
	func() struct{} {
		// -------------
		var err error
		_ = err
		ctx := context.Background()
		_ = ctx
		{
			someDate := xtime.NewDate(2022, 11, 11)
			assert.Equal(t, someDate.LastDateOfMonth(), xtime.NewDate(2022, 11, 30))
		}
		{
			someDate := xtime.NewDate(2022, 11, 01)
			assert.Equal(t, someDate.LastDateOfMonth(), xtime.NewDate(2022, 11, 30))
		}
		{
			someDate := xtime.NewDate(2022, 11, 30)
			assert.Equal(t, someDate.FirstDateOfMonth(), xtime.NewDate(2022, 11, 1))
		}

		// -------------
		return struct{}{}
	}()
}

func TestSplitRange(t *testing.T) {
	type Case struct {
		Name   string
		Days   uint
		Begin  xtime.Date
		End    xtime.Date
		Result []xtime.DateRange
		Err    error
	}
	cases := []Case{
		{
			"0 3",
			0,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 3),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 1)},
				{xtime.NewDate(2000, 1, 2), xtime.NewDate(2000, 1, 2)},
				{xtime.NewDate(2000, 1, 3), xtime.NewDate(2000, 1, 3)},
			},
			nil,
		},
		{
			"1 3",
			1,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 3),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 1)},
				{xtime.NewDate(2000, 1, 2), xtime.NewDate(2000, 1, 2)},
				{xtime.NewDate(2000, 1, 3), xtime.NewDate(2000, 1, 3)},
			},
			nil,
		},
		{
			"2 3",
			2,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 3),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 2)},
				{xtime.NewDate(2000, 1, 3), xtime.NewDate(2000, 1, 3)},
			},
			nil,
		},
		{
			"3 3",
			3,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 3),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 3)},
			},
			nil,
		},
		{
			"4 3",
			4,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 3),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 3)},
			},
			nil,
		},
		{
			"2 6",
			2,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 6),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 2)},
				{xtime.NewDate(2000, 1, 3), xtime.NewDate(2000, 1, 4)},
				{xtime.NewDate(2000, 1, 5), xtime.NewDate(2000, 1, 6)},
			},
			nil,
		},
		{
			"2 7",
			2,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 7),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 2)},
				{xtime.NewDate(2000, 1, 3), xtime.NewDate(2000, 1, 4)},
				{xtime.NewDate(2000, 1, 5), xtime.NewDate(2000, 1, 6)},
				{xtime.NewDate(2000, 1, 7), xtime.NewDate(2000, 1, 7)},
			},
			nil,
		},
		{
			"2 8",
			2,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 1, 8),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 2)},
				{xtime.NewDate(2000, 1, 3), xtime.NewDate(2000, 1, 4)},
				{xtime.NewDate(2000, 1, 5), xtime.NewDate(2000, 1, 6)},
				{xtime.NewDate(2000, 1, 7), xtime.NewDate(2000, 1, 8)},
			},
			nil,
		},
		{
			"跨月 10 40",
			10,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 2, 11),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 10)},
				{xtime.NewDate(2000, 1, 11), xtime.NewDate(2000, 1, 20)},
				{xtime.NewDate(2000, 1, 21), xtime.NewDate(2000, 1, 30)},
				{xtime.NewDate(2000, 1, 31), xtime.NewDate(2000, 2, 9)},
				{xtime.NewDate(2000, 2, 10), xtime.NewDate(2000, 2, 11)},
			},
			nil,
		},
		{
			"跨月 33 40",
			33,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2000, 2, 8),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 2, 2)},
				{xtime.NewDate(2000, 2, 3), xtime.NewDate(2000, 2, 8)},
			},
			nil,
		},
		{
			"跨年 350 400",
			350,
			xtime.NewDate(2000, 1, 1),
			xtime.NewDate(2001, 3, 1),
			[]xtime.DateRange{
				{xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 12, 15)},
				{xtime.NewDate(2000, 12, 16), xtime.NewDate(2001, 3, 1)},
			},
			nil,
		},
	}
	DateRangeToString := func(rs []xtime.DateRange) string {
		t := []string{}
		for _, r := range rs {
			t = append(t, r.Start.String()+"~"+r.End.String())
		}
		return strings.Join(t, " ")
	}
	for _, c := range cases {
		r := xtime.DateRange{c.Begin, c.End}
		result := xtime.SplitRange(c.Days, r)
		assert.Equal(t, DateRangeToString(c.Result), DateRangeToString(result), c.Name)
	}
}

func TestDateRange_Validator(t *testing.T) {
	// Start=End
	{
		err := xtime.DateRange{}.Validator()
		assert.NoError(t, err)
	}
	// Start>End error
	{
		err := xtime.DateRange{
			xtime.NewDate(2000, 1, 1), xtime.Date{},
		}.Validator()
		assert.Errorf(t, err, "goclub/time: DateRange.Start can not be after DateRange.End")
	}
	// Start<End
	{
		err := xtime.DateRange{
			xtime.Date{}, xtime.NewDate(2000, 1, 1),
		}.Validator()
		assert.NoError(t, err)
	}
	// Start=End
	{
		err := xtime.DateRange{
			xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 1),
		}.Validator()
		assert.NoError(t, err)
	}
	// Start<End
	{
		err := xtime.DateRange{
			xtime.NewDate(2000, 1, 1), xtime.NewDate(2000, 1, 2),
		}.Validator()
		assert.NoError(t, err)
	}
}

func TestRange_Validator(t *testing.T) {
	// start > end
	{
		err := xtime.Range{
			time.Date(2000, 1, 2, 0, 0, 1, 0, time.Local),
			time.Date(2000, 1, 1, 0, 0, 1, 0, time.Local),
		}.Validator()
		assert.Errorf(t, err, "goclub/time: Range.Start can not be after Range.End")
	}
	// start = end
	{
		err := xtime.Range{
			time.Date(2000, 1, 1, 0, 0, 1, 0, time.Local),
			time.Date(2000, 1, 1, 0, 0, 1, 0, time.Local),
		}.Validator()
		assert.NoError(t, err)
	}
	// start < end
	{
		err := xtime.Range{
			time.Date(2000, 1, 1, 0, 0, 1, 0, time.Local),
			time.Date(2000, 1, 2, 0, 0, 1, 0, time.Local),
		}.Validator()
		assert.NoError(t, err)
	}
}
