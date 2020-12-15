package xtime_test

import (
	xjson "github.com/goclub/json"
	xtime "github.com/goclub/time"
	"log"
	"testing"
	"time"
)

func ExampleParseChinaYear() {
	log.Print("ExampleParseChinaYear")
	year, err := xtime.ParseChinaYear("2020") ; if err != nil {panic(err)}
	log.Print(year.String()) // 2020-01-01 00:00:00 +0800 CST
}
func ExampleParseChinaYearAndMonth() {
	log.Print("ExampleParseChinaYearAndMonth")
	yearAndMonth, err := xtime.ParseChinaYearAndMonth("2020-11") ; if err != nil {panic(err)}
	log.Print(yearAndMonth.String()) // 2020-11-01 00:00:00 +0800 CST
}
func ExampleParseChinaDate() {
	log.Print("ExampleParseChinaDate")
	date, err := xtime.ParseChinaDate("2020-11-11") ; if err != nil {panic(err)}
	log.Print(date.String()) // 2020-11-11 00:00:00 +0800 CST
}
func ExampleParseChinaTime() {
	log.Print("ExampleParseChinaTime")
	sometime, err := xtime.ParseChinaTime("2020-11-11 21:52:24") ; if err != nil {panic(err)}
	log.Print(sometime.String()) // 2020-11-11 21:52:24 +0800 CST
}
func ExampleFormatChinaYear() {
	log.Print("ExampleFormatChinaYear")
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	log.Print(xtime.FormatChinaYear(sometime)) // 2021
}
func ExampleFormatChinaYearAndMonth() {
	log.Print("ExampleFormatChinaYearAndMonth")
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	log.Print(xtime.FormatChinaYearAndMonth(sometime)) // 2021
}

func ExampleFormatChinaDate() {
	log.Print("ExampleFormatChinaDate")
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	log.Print(xtime.FormatChinaDate(sometime)) // 2021-01-01
}

func ExampleFormatChinaTime() {
	log.Print("ExampleFormatChinaTime")
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	log.Print(xtime.FormatChinaTime(sometime)) // 2021-01-01 07:23:23
}

func ExampleFormatChinaHourMinuteSecond() {
	log.Print("ExampleFormatChinaHourMinuteSecond")
	sometime := time.Date(2020,12,31,23,23,23, 0,time.UTC)
	log.Print(xtime.FormatChinaHourMinuteSecond(sometime)) // 07:23:23
}

func ExampleLocationChina() {
	log.Print("now time is :" + time.Now().In(xtime.LocationChina).String())
}
// 直接使用 time.Time 作为json 字段时因为 time.Time{}.UnmarshalJSON() 和 time.Time{}.MarshalJSON() 的原因会以 time.RFC3339 格式作为 layout
func ExampleJSON_RFC3339() {
	log.Print("ExampleJSON_RFC3339")
	request := struct {
		Time time.Time `json:"time"`
	}{}
	err := xjson.Unmarshal([]byte(`{"time": "2020-12-31T23:23:23Z"}`), &request) ; if err != nil {panic(err)}
	log.Printf("request: %+v", request) // request: {Time:2020-12-31 23:23:23 +0000 UTC}

	response := struct {
		Time time.Time `json:"time"`
	}{Time: time.Date(2020,12,31,23,23,23, 0,time.UTC)}
	data, err := xjson.Marshal(response) ; if err != nil {panic(err)}
	log.Print("response json : " + string(data)) // response json : {"time":"2020-12-31T23:23:23Z"}
}
func ExampleNewChinaTime () {
	log.Print("ExampleNewChinaTime")
	request := struct {
		Time xtime.ChinaTime `json:"time"`
	}{}
	err := xjson.Unmarshal([]byte(`{"time": "2020-12-31 23:23:23"}`), &request) ; if err != nil {panic(err)}
	log.Printf("request: %+v", request) // request: {Time:2020-12-31 23:23:23 +0800 CST}

	response := struct {
		Time xtime.ChinaTime `json:"time"`
	}{Time: xtime.NewChinaTime(time.Date(2020,12,31,23,23,23, 0,time.UTC))}
	data, err := xjson.Marshal(response) ; if err != nil {panic(err)}
	log.Print("response json : " + string(data)) // response json : {"time":"2021-01-01 07:23:23"}
}
func TestExample(t *testing.T) {

}