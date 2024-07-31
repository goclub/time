---
permalink: /
sidebarBasedOnContent: true
---

# xtime

```go
import "github.com/goclub/time"
```

`time` 是 go 中非常重要的一个包,使用 time.Time 时一定要注意时区。

## Parse & Format 

在将字符串解析为 `time.Time` 和将 `time.Time` 转换为字符串时如果不指定时区则会按照当前机器的本地时区解析和转换。

```go
layoutTime := "2006-01-02 15:04:05"
someTime, err := time.Parse(layoutTime, "2020-12-31 22:00:00") ; if err != nil {
    panic(err)
}
log.Print(someTime.String()) // 不同时区的机器运行结果不同
```

当 layout 不包含时区时，以服务端时区解析字符串是不严谨的。
Format 也有相同的问题。

故此 goclub/time 提供了一些指定时区的 parse 和 format，目前主要是 China 时区 `time.FixedZone("CST", 8*3600)`

当 json 传递的时间格式不是 RFC3339 而是中国时区年月日时分秒 `2006-01-02 15:04:05`，可以使用 `xtime.ChinaTime` 或者 `xtime.ChinaRange` 解析和转换

## Date

在数据库中或者前后端数据传递中都会用到日期, `xtime.Date` 和 `xtime.NullDate` 支持 JSON 和SQL 解析 

> 无论 sql 是否打开 `parseTime=true`, `xtime.Date` 都能正确解析日期

并提供 `xtime.NewDate` `xtime.NewDateFromTime` `xtime.NewDateFromString` 等方法创建数据

> 注意 xtime.Date 表达的是日期所以无需时区,但是 xtime.Date 转换为 time.Time 时候需要指定时区


## FirstSecondOfDate & LastSecondOfDate

拿到指定日期的第一秒和最后一秒

**看源码辅助理解**
```go
func LastSecondOfDate(t time.Time) time.Time {
y,m,d := t.Date()
return time.Date(y,m,d,23,59,59,999999999,t.Location())
}
func FirstSecondOfDate(t time.Time) time.Time {
y,m,d := t.Date()
return time.Date(y,m,d,0,0,0,0,t.Location())
}
```

## TomorrowFirstSecond

明天第0秒,多用于 redis [PEXPIREAT](pexpireat)  

```go
sometime := time.Date(2022, 01, 01, 12, 12, 12, 88, xtime.LocChina)

tomorrowFirstSecond := xtime.TomorrowFirstSecond(sometime)

tomorrowFirstSecond.String() // 2022-01-02 00:00:00 +0800 CST
tomorrowFirstSecond.UnixMilli() // int64(1641052800000)  or xtime.UnixMilli(tomorrowFirstSecond)

// 计算现在距离明天00:00:00的时间有多少
now := time.Now().In(xtime.LocChina)
xtime.TomorrowFirstSecond(now).Sub(now)

```

## InRange & InRangeFromDate

`InRange` 和 `InRangeFromDate` 可以判断一个时间或者日期是否在指定范围内.

```go

xtime.InRangeFromDate(time.Date(2022,01,02,0,0,0,0, xtime.LocChina), xtime.DateRange{
    Begin:  xtime.NewDate(2022,01,01),
    End:    xtime.NewDate(2022,01,03),
}) // true
xtime.InRangeFromDate(time.Date(2022,01,05,0,0,0,0, xtime.LocChina), xtime.DateRange{
Begin:  xtime.NewDate(2022,01,01),
End:    xtime.NewDate(2022,01,03),
}) // false


var begin time.Time
var end time.Time
in := xtime.InRange(target, xtime.Range{
    Begin:  begin,
    End:    end,
})
```

## SplitDateRange

`SplitDateRange` 将一个时间范围分隔成多个指定间隔的日期范围

用来满足一些业务场景中希望将长跨越日期改成短跨度日期后分段查询的需求.

```go
var days uint = 10
splitDateRanges := xtime.SplitDateRange(days, xtime.DateRange{
    Begin:  xtime.NewDate(2000,1,1),
    End:    xtime.NewDate(2000,2,8),
})
//splitDateRanges:
//[
//    2000-01-01~2000-01-10
//    2000-01-11~2000-01-20
//    2000-01-21~2000-01-30
//    2000-01-31~2000-02-09
//    2000-02-10~2000-02-11
//]
```