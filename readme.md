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

当 json 传递的时间格式不是 RFC3339 而是中国时区年月日时分秒 `2006-01-02 15:04:05`，可以使用 `xtime.ChinaTime` 解析和转换

## Date

在数据库中或者前后端数据传递中都会用到日期, `xtime.Date` 和 `xtime.NullDate` 支持 JSON 和SQL 解析

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