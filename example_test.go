package xtime_test

import (
	xjson "github.com/goclub/json"
	xtime "github.com/goclub/time"
	"log"
	"testing"
	"time"
)

func ExampleParse() {
	log.Print("ExampleParse")
	year, err := xtime.ParseChinaYear("2020") ; if err != nil {panic(err)}
	log.Print(year.String())
	yearAndMonth, err := xtime.ParseChinaYearAndMonth("2020-11") ; if err != nil {panic(err)}
	log.Print(yearAndMonth.String())
	date, err := xtime.ParseChinaDate("2020-11-11") ; if err != nil {panic(err)}
	log.Print(date.String())
	sometime, err := xtime.ParseChinaTime("2020-11-11 21:52:24") ; if err != nil {panic(err)}
	log.Print(sometime.String())
}
func ExampleFormat() {
	log.Print("ExampleFormat")
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	log.Print(xtime.FormatChinaYear(sometime)) // 2021
	log.Print(xtime.FormatChinaYearAndMonth(sometime)) // 2021-01
	log.Print(xtime.FormatChinaDate(sometime)) // 2021-01-01
	log.Print(xtime.FormatChinaTime(sometime)) // 2021-01-01 07:23:23
	log.Print(xtime.FormatChinaHourMinuteSecond(sometime)) // 07:23:23
}

func ExampleLocation() {
	log.Print("now time is :" + time.Now().In(xtime.LocationChina).String())
}
// 直接使用 time.Time 作为json 字段时因为 time.Time{}.UnmarshalJSON() 和 time.Time{}.MarshalJSON() 的原因会以 time.RFC3339 格式作为 layout
func ExampleJSON_RFC3339() {
	log.Print("ExampleJSON_RFC3339")
	reqeust := struct {
		Time time.Time `json:"time"`
	}{}
	err := xjson.Unmarshal([]byte(`{"time": "2020-12-31T23:23:23Z"}`), &reqeust) ; if err != nil {panic(err)}
	log.Printf("reqeust: %+v", reqeust)

	response := struct {
		Time time.Time `json:"time"`
	}{Time: time.Date(2020,12,31,23,23,23, 0,time.UTC)}
	data, err := xjson.Marshal(response) ; if err != nil {panic(err)}
	log.Print("response json : " + string(data)) // {"time":"2020-12-31T23:23:23Z"}
}
func ExampleJSONChinaTime () {
	log.Print("ExampleJSONChinaTime")
	reqeust := struct {
		Time xtime.ChinaTime `json:"time"`
	}{}
	err := xjson.Unmarshal([]byte(`{"time": "2020-12-31 23:23:23"}`), &reqeust) ; if err != nil {panic(err)}
	log.Printf("reqeust: %+v", reqeust)

	response := struct {
		Time xtime.ChinaTime `json:"time"`
	}{Time: xtime.NewChinaTime(time.Date(2020,12,31,23,23,23, 0,time.UTC))}
	data, err := xjson.Marshal(response) ; if err != nil {panic(err)}
	log.Print("response json : " + string(data)) // {"time":"2020-12-31 23:23:23"}
}
func TestExample(t *testing.T) {
	ExampleFormat()
	ExampleLocation()
	ExampleParse()
	ExampleJSON_RFC3339()
	ExampleJSONChinaTime()
}